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

func configTermTypesRouter(router *httprouter.Router) {
	router.GET("/termtypes", GetAllTermTypes)
	router.POST("/termtypes", AddTermTypes)
	router.GET("/termtypes/:argId", GetTermTypes)
	router.PUT("/termtypes/:argId", UpdateTermTypes)
	router.DELETE("/termtypes/:argId", DeleteTermTypes)
}

func configGinTermTypesRouter(router gin.IRoutes) {
	router.GET("/termtypes", ConverHttprouterToGin(GetAllTermTypes))
	router.POST("/termtypes", ConverHttprouterToGin(AddTermTypes))
	router.GET("/termtypes/:argId", ConverHttprouterToGin(GetTermTypes))
	router.PUT("/termtypes/:argId", ConverHttprouterToGin(UpdateTermTypes))
	router.DELETE("/termtypes/:argId", ConverHttprouterToGin(DeleteTermTypes))
}

// GetAllTermTypes is a function to get a slice of record(s) from term_types table in the laravel database
// @Summary Get list of TermTypes
// @Tags TermTypes
// @Description GetAllTermTypes is a handler to get a slice of record(s) from term_types table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.TermTypes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /termtypes [get]
// http "http://localhost:8080/termtypes?page=0&pagesize=20" X-Api-User:user123
func GetAllTermTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "term_types", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTermTypes(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTermTypes is a function to get a single record from the term_types table in the laravel database
// @Summary Get record from table TermTypes by  argId
// @Tags TermTypes
// @ID argId
// @Description GetTermTypes is a function to get a single record from the term_types table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.TermTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /termtypes/{argId} [get]
// http "http://localhost:8080/termtypes/1" X-Api-User:user123
func GetTermTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "term_types", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTermTypes(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTermTypes add to add a single record to term_types table in the laravel database
// @Summary Add an record to term_types table
// @Description add to add a single record to term_types table in the laravel database
// @Tags TermTypes
// @Accept  json
// @Produce  json
// @Param TermTypes body model.TermTypes true "Add TermTypes"
// @Success 200 {object} model.TermTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /termtypes [post]
// echo '{"id": 30,"slug": "nlGcWMqHaFHypCjgTyDSwcjlW","name": "LsmHEBZaKLvZNweVkdtVtSDoE","created_at": "2170-01-27T05:23:11.311774884-03:00","updated_at": "2195-07-15T00:33:32.547701313-03:00"}' | http POST "http://localhost:8080/termtypes" X-Api-User:user123
func AddTermTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	termtypes := &model.TermTypes{}

	if err := readJSON(r, termtypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := termtypes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	termtypes.Prepare()

	if err := termtypes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "term_types", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	termtypes, _, err = dao.AddTermTypes(ctx, termtypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, termtypes)
}

// UpdateTermTypes Update a single record from term_types table in the laravel database
// @Summary Update an record in table term_types
// @Description Update a single record from term_types table in the laravel database
// @Tags TermTypes
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  TermTypes body model.TermTypes true "Update TermTypes record"
// @Success 200 {object} model.TermTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /termtypes/{argId} [put]
// echo '{"id": 30,"slug": "nlGcWMqHaFHypCjgTyDSwcjlW","name": "LsmHEBZaKLvZNweVkdtVtSDoE","created_at": "2170-01-27T05:23:11.311774884-03:00","updated_at": "2195-07-15T00:33:32.547701313-03:00"}' | http PUT "http://localhost:8080/termtypes/1"  X-Api-User:user123
func UpdateTermTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	termtypes := &model.TermTypes{}
	if err := readJSON(r, termtypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := termtypes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	termtypes.Prepare()

	if err := termtypes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "term_types", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	termtypes, _, err = dao.UpdateTermTypes(ctx,
		argId,
		termtypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, termtypes)
}

// DeleteTermTypes Delete a single record from term_types table in the laravel database
// @Summary Delete a record from term_types
// @Description Delete a single record from term_types table in the laravel database
// @Tags TermTypes
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.TermTypes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /termtypes/{argId} [delete]
// http DELETE "http://localhost:8080/termtypes/1" X-Api-User:user123
func DeleteTermTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "term_types", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTermTypes(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
