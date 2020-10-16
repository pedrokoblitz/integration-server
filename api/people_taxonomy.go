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

func configPeopleTaxonomyRouter(router *httprouter.Router) {
	router.GET("/peopletaxonomy", GetAllPeopleTaxonomy)
	router.POST("/peopletaxonomy", AddPeopleTaxonomy)
	router.GET("/peopletaxonomy/:argId", GetPeopleTaxonomy)
	router.PUT("/peopletaxonomy/:argId", UpdatePeopleTaxonomy)
	router.DELETE("/peopletaxonomy/:argId", DeletePeopleTaxonomy)
}

func configGinPeopleTaxonomyRouter(router gin.IRoutes) {
	router.GET("/peopletaxonomy", ConverHttprouterToGin(GetAllPeopleTaxonomy))
	router.POST("/peopletaxonomy", ConverHttprouterToGin(AddPeopleTaxonomy))
	router.GET("/peopletaxonomy/:argId", ConverHttprouterToGin(GetPeopleTaxonomy))
	router.PUT("/peopletaxonomy/:argId", ConverHttprouterToGin(UpdatePeopleTaxonomy))
	router.DELETE("/peopletaxonomy/:argId", ConverHttprouterToGin(DeletePeopleTaxonomy))
}

// GetAllPeopleTaxonomy is a function to get a slice of record(s) from people_taxonomy table in the laravel database
// @Summary Get list of PeopleTaxonomy
// @Tags PeopleTaxonomy
// @Description GetAllPeopleTaxonomy is a handler to get a slice of record(s) from people_taxonomy table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PeopleTaxonomy}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /peopletaxonomy [get]
// http "http://localhost:8080/peopletaxonomy?page=0&pagesize=20" X-Api-User:user123
func GetAllPeopleTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "people_taxonomy", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPeopleTaxonomy(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPeopleTaxonomy is a function to get a single record from the people_taxonomy table in the laravel database
// @Summary Get record from table PeopleTaxonomy by  argId
// @Tags PeopleTaxonomy
// @ID argId
// @Description GetPeopleTaxonomy is a function to get a single record from the people_taxonomy table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.PeopleTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /peopletaxonomy/{argId} [get]
// http "http://localhost:8080/peopletaxonomy/1" X-Api-User:user123
func GetPeopleTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "people_taxonomy", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPeopleTaxonomy(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPeopleTaxonomy add to add a single record to people_taxonomy table in the laravel database
// @Summary Add an record to people_taxonomy table
// @Description add to add a single record to people_taxonomy table in the laravel database
// @Tags PeopleTaxonomy
// @Accept  json
// @Produce  json
// @Param PeopleTaxonomy body model.PeopleTaxonomy true "Add PeopleTaxonomy"
// @Success 200 {object} model.PeopleTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /peopletaxonomy [post]
// echo '{"updated_at": "2054-06-15T05:00:38.66784353-03:00","id": 94,"person_id": 77,"product_id": 72,"created_at": "2176-03-14T21:07:30.173730685-03:00"}' | http POST "http://localhost:8080/peopletaxonomy" X-Api-User:user123
func AddPeopleTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	peopletaxonomy := &model.PeopleTaxonomy{}

	if err := readJSON(r, peopletaxonomy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := peopletaxonomy.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	peopletaxonomy.Prepare()

	if err := peopletaxonomy.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "people_taxonomy", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	peopletaxonomy, _, err = dao.AddPeopleTaxonomy(ctx, peopletaxonomy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, peopletaxonomy)
}

// UpdatePeopleTaxonomy Update a single record from people_taxonomy table in the laravel database
// @Summary Update an record in table people_taxonomy
// @Description Update a single record from people_taxonomy table in the laravel database
// @Tags PeopleTaxonomy
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  PeopleTaxonomy body model.PeopleTaxonomy true "Update PeopleTaxonomy record"
// @Success 200 {object} model.PeopleTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /peopletaxonomy/{argId} [put]
// echo '{"updated_at": "2054-06-15T05:00:38.66784353-03:00","id": 94,"person_id": 77,"product_id": 72,"created_at": "2176-03-14T21:07:30.173730685-03:00"}' | http PUT "http://localhost:8080/peopletaxonomy/1"  X-Api-User:user123
func UpdatePeopleTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	peopletaxonomy := &model.PeopleTaxonomy{}
	if err := readJSON(r, peopletaxonomy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := peopletaxonomy.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	peopletaxonomy.Prepare()

	if err := peopletaxonomy.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "people_taxonomy", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	peopletaxonomy, _, err = dao.UpdatePeopleTaxonomy(ctx,
		argId,
		peopletaxonomy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, peopletaxonomy)
}

// DeletePeopleTaxonomy Delete a single record from people_taxonomy table in the laravel database
// @Summary Delete a record from people_taxonomy
// @Description Delete a single record from people_taxonomy table in the laravel database
// @Tags PeopleTaxonomy
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.PeopleTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /peopletaxonomy/{argId} [delete]
// http DELETE "http://localhost:8080/peopletaxonomy/1" X-Api-User:user123
func DeletePeopleTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "people_taxonomy", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePeopleTaxonomy(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
