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

func configPageTypesRouter(router *httprouter.Router) {
	router.GET("/pagetypes", GetAllPageTypes)
	router.POST("/pagetypes", AddPageTypes)
	router.GET("/pagetypes/:argId", GetPageTypes)
	router.PUT("/pagetypes/:argId", UpdatePageTypes)
	router.DELETE("/pagetypes/:argId", DeletePageTypes)
}

func configGinPageTypesRouter(router gin.IRoutes) {
	router.GET("/pagetypes", ConverHttprouterToGin(GetAllPageTypes))
	router.POST("/pagetypes", ConverHttprouterToGin(AddPageTypes))
	router.GET("/pagetypes/:argId", ConverHttprouterToGin(GetPageTypes))
	router.PUT("/pagetypes/:argId", ConverHttprouterToGin(UpdatePageTypes))
	router.DELETE("/pagetypes/:argId", ConverHttprouterToGin(DeletePageTypes))
}

// GetAllPageTypes is a function to get a slice of record(s) from page_types table in the laravel database
// @Summary Get list of PageTypes
// @Tags PageTypes
// @Description GetAllPageTypes is a handler to get a slice of record(s) from page_types table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PageTypes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pagetypes [get]
// http "http://localhost:8080/pagetypes?page=0&pagesize=20" X-Api-User:user123
func GetAllPageTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "page_types", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPageTypes(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPageTypes is a function to get a single record from the page_types table in the laravel database
// @Summary Get record from table PageTypes by  argId
// @Tags PageTypes
// @ID argId
// @Description GetPageTypes is a function to get a single record from the page_types table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.PageTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /pagetypes/{argId} [get]
// http "http://localhost:8080/pagetypes/1" X-Api-User:user123
func GetPageTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "page_types", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPageTypes(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPageTypes add to add a single record to page_types table in the laravel database
// @Summary Add an record to page_types table
// @Description add to add a single record to page_types table in the laravel database
// @Tags PageTypes
// @Accept  json
// @Produce  json
// @Param PageTypes body model.PageTypes true "Add PageTypes"
// @Success 200 {object} model.PageTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pagetypes [post]
// echo '{"is_email": 37,"name": "KaYxUKiDBIyOJZjIjFjRQXEHF","slug": "NwhJvZQBUFOCeUHUsmgOsZOEn","created_at": "2082-11-05T03:36:08.833201116-03:00","updated_at": "2267-03-12T06:23:43.851351257-03:00","id": 33,"app_id": 89}' | http POST "http://localhost:8080/pagetypes" X-Api-User:user123
func AddPageTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	pagetypes := &model.PageTypes{}

	if err := readJSON(r, pagetypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := pagetypes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	pagetypes.Prepare()

	if err := pagetypes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "page_types", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	pagetypes, _, err = dao.AddPageTypes(ctx, pagetypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, pagetypes)
}

// UpdatePageTypes Update a single record from page_types table in the laravel database
// @Summary Update an record in table page_types
// @Description Update a single record from page_types table in the laravel database
// @Tags PageTypes
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  PageTypes body model.PageTypes true "Update PageTypes record"
// @Success 200 {object} model.PageTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pagetypes/{argId} [put]
// echo '{"is_email": 37,"name": "KaYxUKiDBIyOJZjIjFjRQXEHF","slug": "NwhJvZQBUFOCeUHUsmgOsZOEn","created_at": "2082-11-05T03:36:08.833201116-03:00","updated_at": "2267-03-12T06:23:43.851351257-03:00","id": 33,"app_id": 89}' | http PUT "http://localhost:8080/pagetypes/1"  X-Api-User:user123
func UpdatePageTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	pagetypes := &model.PageTypes{}
	if err := readJSON(r, pagetypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := pagetypes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	pagetypes.Prepare()

	if err := pagetypes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "page_types", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	pagetypes, _, err = dao.UpdatePageTypes(ctx,
		argId,
		pagetypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, pagetypes)
}

// DeletePageTypes Delete a single record from page_types table in the laravel database
// @Summary Delete a record from page_types
// @Description Delete a single record from page_types table in the laravel database
// @Tags PageTypes
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.PageTypes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /pagetypes/{argId} [delete]
// http DELETE "http://localhost:8080/pagetypes/1" X-Api-User:user123
func DeletePageTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "page_types", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePageTypes(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
