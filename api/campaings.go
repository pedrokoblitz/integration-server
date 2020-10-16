package api

import (
	"net/http"

	"github.com/pedrokoblitz/comm/dao"
	"github.com/pedrokoblitz/comm/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configCampaingsRouter(router *httprouter.Router) {
	router.GET("/campaings", GetAllCampaings)
	router.POST("/campaings", AddCampaings)
	router.GET("/campaings/:argId", GetCampaings)
	router.PUT("/campaings/:argId", UpdateCampaings)
	router.DELETE("/campaings/:argId", DeleteCampaings)
}

func configGinCampaingsRouter(router gin.IRoutes) {
	router.GET("/campaings", ConverHttprouterToGin(GetAllCampaings))
	router.POST("/campaings", ConverHttprouterToGin(AddCampaings))
	router.GET("/campaings/:argId", ConverHttprouterToGin(GetCampaings))
	router.PUT("/campaings/:argId", ConverHttprouterToGin(UpdateCampaings))
	router.DELETE("/campaings/:argId", ConverHttprouterToGin(DeleteCampaings))
}

// GetAllCampaings is a function to get a slice of record(s) from campaings table in the laravel database
// @Summary Get list of Campaings
// @Tags Campaings
// @Description GetAllCampaings is a handler to get a slice of record(s) from campaings table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Campaings}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /campaings [get]
// http "http://localhost:8080/campaings?page=0&pagesize=20" X-Api-User:user123
func GetAllCampaings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "campaings", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCampaings(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetCampaings is a function to get a single record from the campaings table in the laravel database
// @Summary Get record from table Campaings by  argId
// @Tags Campaings
// @ID argId
// @Description GetCampaings is a function to get a single record from the campaings table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.Campaings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /campaings/{argId} [get]
// http "http://localhost:8080/campaings/1" X-Api-User:user123
func GetCampaings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "campaings", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCampaings(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddCampaings add to add a single record to campaings table in the laravel database
// @Summary Add an record to campaings table
// @Description add to add a single record to campaings table in the laravel database
// @Tags Campaings
// @Accept  json
// @Produce  json
// @Param Campaings body model.Campaings true "Add Campaings"
// @Success 200 {object} model.Campaings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /campaings [post]
// echo '{"app_id": 97,"page_id": 28,"name": "JvVDYaMbpQJqxnEDaRJcAOiFo","slug": "NVrwkCPMqYXKHxegChpeSiequ","last_sent": "2309-11-25T03:30:31.958107909-03:00","created_at": "2126-04-10T13:06:44.06588567-03:00","updated_at": "2153-07-13T23:51:54.485129851-03:00","id": 63}' | http POST "http://localhost:8080/campaings" X-Api-User:user123
func AddCampaings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	campaings := &model.Campaings{}

	if err := readJSON(r, campaings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := campaings.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	campaings.Prepare()

	if err := campaings.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "campaings", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	campaings, _, err = dao.AddCampaings(ctx, campaings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, campaings)
}

// UpdateCampaings Update a single record from campaings table in the laravel database
// @Summary Update an record in table campaings
// @Description Update a single record from campaings table in the laravel database
// @Tags Campaings
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  Campaings body model.Campaings true "Update Campaings record"
// @Success 200 {object} model.Campaings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /campaings/{argId} [put]
// echo '{"app_id": 97,"page_id": 28,"name": "JvVDYaMbpQJqxnEDaRJcAOiFo","slug": "NVrwkCPMqYXKHxegChpeSiequ","last_sent": "2309-11-25T03:30:31.958107909-03:00","created_at": "2126-04-10T13:06:44.06588567-03:00","updated_at": "2153-07-13T23:51:54.485129851-03:00","id": 63}' | http PUT "http://localhost:8080/campaings/1"  X-Api-User:user123
func UpdateCampaings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	campaings := &model.Campaings{}
	if err := readJSON(r, campaings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := campaings.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	campaings.Prepare()

	if err := campaings.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "campaings", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	campaings, _, err = dao.UpdateCampaings(ctx,
		argId,
		campaings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, campaings)
}

// DeleteCampaings Delete a single record from campaings table in the laravel database
// @Summary Delete a record from campaings
// @Description Delete a single record from campaings table in the laravel database
// @Tags Campaings
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.Campaings
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /campaings/{argId} [delete]
// http DELETE "http://localhost:8080/campaings/1" X-Api-User:user123
func DeleteCampaings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "campaings", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCampaings(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
