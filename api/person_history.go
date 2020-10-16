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

func configPersonHistoryRouter(router *httprouter.Router) {
	router.GET("/personhistory", GetAllPersonHistory)
	router.POST("/personhistory", AddPersonHistory)
	router.GET("/personhistory/:argId", GetPersonHistory)
	router.PUT("/personhistory/:argId", UpdatePersonHistory)
	router.DELETE("/personhistory/:argId", DeletePersonHistory)
}

func configGinPersonHistoryRouter(router gin.IRoutes) {
	router.GET("/personhistory", ConverHttprouterToGin(GetAllPersonHistory))
	router.POST("/personhistory", ConverHttprouterToGin(AddPersonHistory))
	router.GET("/personhistory/:argId", ConverHttprouterToGin(GetPersonHistory))
	router.PUT("/personhistory/:argId", ConverHttprouterToGin(UpdatePersonHistory))
	router.DELETE("/personhistory/:argId", ConverHttprouterToGin(DeletePersonHistory))
}

// GetAllPersonHistory is a function to get a slice of record(s) from person_history table in the laravel database
// @Summary Get list of PersonHistory
// @Tags PersonHistory
// @Description GetAllPersonHistory is a handler to get a slice of record(s) from person_history table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PersonHistory}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personhistory [get]
// http "http://localhost:8080/personhistory?page=0&pagesize=20" X-Api-User:user123
func GetAllPersonHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "person_history", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPersonHistory(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPersonHistory is a function to get a single record from the person_history table in the laravel database
// @Summary Get record from table PersonHistory by  argId
// @Tags PersonHistory
// @ID argId
// @Description GetPersonHistory is a function to get a single record from the person_history table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.PersonHistory
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /personhistory/{argId} [get]
// http "http://localhost:8080/personhistory/1" X-Api-User:user123
func GetPersonHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "person_history", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPersonHistory(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPersonHistory add to add a single record to person_history table in the laravel database
// @Summary Add an record to person_history table
// @Description add to add a single record to person_history table in the laravel database
// @Tags PersonHistory
// @Accept  json
// @Produce  json
// @Param PersonHistory body model.PersonHistory true "Add PersonHistory"
// @Success 200 {object} model.PersonHistory
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personhistory [post]
// echo '{"app_id": 81,"updated_at": "2192-11-08T10:06:09.192112558-03:00","user_identifier": "latgmrgMCEuODfAMPJngIGbKC","deleted_at": "2171-03-14T15:17:56.325320534-03:00","created_at": "2118-04-29T09:08:03.664289448-03:00","id": 1,"action_id": 81,"person_id": 29,"campaign_id": 34}' | http POST "http://localhost:8080/personhistory" X-Api-User:user123
func AddPersonHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	personhistory := &model.PersonHistory{}

	if err := readJSON(r, personhistory); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := personhistory.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	personhistory.Prepare()

	if err := personhistory.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "person_history", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	personhistory, _, err = dao.AddPersonHistory(ctx, personhistory)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, personhistory)
}

// UpdatePersonHistory Update a single record from person_history table in the laravel database
// @Summary Update an record in table person_history
// @Description Update a single record from person_history table in the laravel database
// @Tags PersonHistory
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  PersonHistory body model.PersonHistory true "Update PersonHistory record"
// @Success 200 {object} model.PersonHistory
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personhistory/{argId} [put]
// echo '{"app_id": 81,"updated_at": "2192-11-08T10:06:09.192112558-03:00","user_identifier": "latgmrgMCEuODfAMPJngIGbKC","deleted_at": "2171-03-14T15:17:56.325320534-03:00","created_at": "2118-04-29T09:08:03.664289448-03:00","id": 1,"action_id": 81,"person_id": 29,"campaign_id": 34}' | http PUT "http://localhost:8080/personhistory/1"  X-Api-User:user123
func UpdatePersonHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	personhistory := &model.PersonHistory{}
	if err := readJSON(r, personhistory); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := personhistory.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	personhistory.Prepare()

	if err := personhistory.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "person_history", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	personhistory, _, err = dao.UpdatePersonHistory(ctx,
		argId,
		personhistory)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, personhistory)
}

// DeletePersonHistory Delete a single record from person_history table in the laravel database
// @Summary Delete a record from person_history
// @Description Delete a single record from person_history table in the laravel database
// @Tags PersonHistory
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.PersonHistory
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /personhistory/{argId} [delete]
// http DELETE "http://localhost:8080/personhistory/1" X-Api-User:user123
func DeletePersonHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "person_history", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePersonHistory(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
