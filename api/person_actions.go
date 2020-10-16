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

func configPersonActionsRouter(router *httprouter.Router) {
	router.GET("/personactions", GetAllPersonActions)
	router.POST("/personactions", AddPersonActions)
	router.GET("/personactions/:argId", GetPersonActions)
	router.PUT("/personactions/:argId", UpdatePersonActions)
	router.DELETE("/personactions/:argId", DeletePersonActions)
}

func configGinPersonActionsRouter(router gin.IRoutes) {
	router.GET("/personactions", ConverHttprouterToGin(GetAllPersonActions))
	router.POST("/personactions", ConverHttprouterToGin(AddPersonActions))
	router.GET("/personactions/:argId", ConverHttprouterToGin(GetPersonActions))
	router.PUT("/personactions/:argId", ConverHttprouterToGin(UpdatePersonActions))
	router.DELETE("/personactions/:argId", ConverHttprouterToGin(DeletePersonActions))
}

// GetAllPersonActions is a function to get a slice of record(s) from person_actions table in the laravel database
// @Summary Get list of PersonActions
// @Tags PersonActions
// @Description GetAllPersonActions is a handler to get a slice of record(s) from person_actions table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PersonActions}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personactions [get]
// http "http://localhost:8080/personactions?page=0&pagesize=20" X-Api-User:user123
func GetAllPersonActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "person_actions", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPersonActions(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPersonActions is a function to get a single record from the person_actions table in the laravel database
// @Summary Get record from table PersonActions by  argId
// @Tags PersonActions
// @ID argId
// @Description GetPersonActions is a function to get a single record from the person_actions table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.PersonActions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /personactions/{argId} [get]
// http "http://localhost:8080/personactions/1" X-Api-User:user123
func GetPersonActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "person_actions", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPersonActions(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPersonActions add to add a single record to person_actions table in the laravel database
// @Summary Add an record to person_actions table
// @Description add to add a single record to person_actions table in the laravel database
// @Tags PersonActions
// @Accept  json
// @Produce  json
// @Param PersonActions body model.PersonActions true "Add PersonActions"
// @Success 200 {object} model.PersonActions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personactions [post]
// echo '{"slug": "sWfbZaaPyFkvMJCIoLXxaadTd","deleted_at": "2290-01-27T12:20:13.379211802-03:00","created_at": "2090-12-31T12:14:10.761563173-03:00","updated_at": "2209-09-09T12:51:28.410053047-03:00","id": 88,"name": "eRhJLaIltrbqaUoLAkaNLvZGn"}' | http POST "http://localhost:8080/personactions" X-Api-User:user123
func AddPersonActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	personactions := &model.PersonActions{}

	if err := readJSON(r, personactions); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := personactions.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	personactions.Prepare()

	if err := personactions.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "person_actions", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	personactions, _, err = dao.AddPersonActions(ctx, personactions)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, personactions)
}

// UpdatePersonActions Update a single record from person_actions table in the laravel database
// @Summary Update an record in table person_actions
// @Description Update a single record from person_actions table in the laravel database
// @Tags PersonActions
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  PersonActions body model.PersonActions true "Update PersonActions record"
// @Success 200 {object} model.PersonActions
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /personactions/{argId} [put]
// echo '{"slug": "sWfbZaaPyFkvMJCIoLXxaadTd","deleted_at": "2290-01-27T12:20:13.379211802-03:00","created_at": "2090-12-31T12:14:10.761563173-03:00","updated_at": "2209-09-09T12:51:28.410053047-03:00","id": 88,"name": "eRhJLaIltrbqaUoLAkaNLvZGn"}' | http PUT "http://localhost:8080/personactions/1"  X-Api-User:user123
func UpdatePersonActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	personactions := &model.PersonActions{}
	if err := readJSON(r, personactions); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := personactions.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	personactions.Prepare()

	if err := personactions.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "person_actions", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	personactions, _, err = dao.UpdatePersonActions(ctx,
		argId,
		personactions)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, personactions)
}

// DeletePersonActions Delete a single record from person_actions table in the laravel database
// @Summary Delete a record from person_actions
// @Description Delete a single record from person_actions table in the laravel database
// @Tags PersonActions
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.PersonActions
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /personactions/{argId} [delete]
// http DELETE "http://localhost:8080/personactions/1" X-Api-User:user123
func DeletePersonActions(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "person_actions", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePersonActions(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
