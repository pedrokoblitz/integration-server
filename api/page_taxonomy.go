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

func configPageTaxonomyRouter(router *httprouter.Router) {
	router.GET("/pagetaxonomy", GetAllPageTaxonomy)
	router.POST("/pagetaxonomy", AddPageTaxonomy)
	router.GET("/pagetaxonomy/:argId", GetPageTaxonomy)
	router.PUT("/pagetaxonomy/:argId", UpdatePageTaxonomy)
	router.DELETE("/pagetaxonomy/:argId", DeletePageTaxonomy)
}

func configGinPageTaxonomyRouter(router gin.IRoutes) {
	router.GET("/pagetaxonomy", ConverHttprouterToGin(GetAllPageTaxonomy))
	router.POST("/pagetaxonomy", ConverHttprouterToGin(AddPageTaxonomy))
	router.GET("/pagetaxonomy/:argId", ConverHttprouterToGin(GetPageTaxonomy))
	router.PUT("/pagetaxonomy/:argId", ConverHttprouterToGin(UpdatePageTaxonomy))
	router.DELETE("/pagetaxonomy/:argId", ConverHttprouterToGin(DeletePageTaxonomy))
}

// GetAllPageTaxonomy is a function to get a slice of record(s) from page_taxonomy table in the laravel database
// @Summary Get list of PageTaxonomy
// @Tags PageTaxonomy
// @Description GetAllPageTaxonomy is a handler to get a slice of record(s) from page_taxonomy table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PageTaxonomy}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pagetaxonomy [get]
// http "http://localhost:8080/pagetaxonomy?page=0&pagesize=20" X-Api-User:user123
func GetAllPageTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "page_taxonomy", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPageTaxonomy(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPageTaxonomy is a function to get a single record from the page_taxonomy table in the laravel database
// @Summary Get record from table PageTaxonomy by  argId
// @Tags PageTaxonomy
// @ID argId
// @Description GetPageTaxonomy is a function to get a single record from the page_taxonomy table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.PageTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /pagetaxonomy/{argId} [get]
// http "http://localhost:8080/pagetaxonomy/1" X-Api-User:user123
func GetPageTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "page_taxonomy", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPageTaxonomy(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPageTaxonomy add to add a single record to page_taxonomy table in the laravel database
// @Summary Add an record to page_taxonomy table
// @Description add to add a single record to page_taxonomy table in the laravel database
// @Tags PageTaxonomy
// @Accept  json
// @Produce  json
// @Param PageTaxonomy body model.PageTaxonomy true "Add PageTaxonomy"
// @Success 200 {object} model.PageTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pagetaxonomy [post]
// echo '{"id": 32,"term_id": 94,"page_id": 38,"created_at": "2214-11-12T19:51:47.877572037-03:00","updated_at": "2248-03-09T11:46:29.808895175-03:00"}' | http POST "http://localhost:8080/pagetaxonomy" X-Api-User:user123
func AddPageTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	pagetaxonomy := &model.PageTaxonomy{}

	if err := readJSON(r, pagetaxonomy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := pagetaxonomy.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	pagetaxonomy.Prepare()

	if err := pagetaxonomy.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "page_taxonomy", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	pagetaxonomy, _, err = dao.AddPageTaxonomy(ctx, pagetaxonomy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, pagetaxonomy)
}

// UpdatePageTaxonomy Update a single record from page_taxonomy table in the laravel database
// @Summary Update an record in table page_taxonomy
// @Description Update a single record from page_taxonomy table in the laravel database
// @Tags PageTaxonomy
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  PageTaxonomy body model.PageTaxonomy true "Update PageTaxonomy record"
// @Success 200 {object} model.PageTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pagetaxonomy/{argId} [put]
// echo '{"id": 32,"term_id": 94,"page_id": 38,"created_at": "2214-11-12T19:51:47.877572037-03:00","updated_at": "2248-03-09T11:46:29.808895175-03:00"}' | http PUT "http://localhost:8080/pagetaxonomy/1"  X-Api-User:user123
func UpdatePageTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	pagetaxonomy := &model.PageTaxonomy{}
	if err := readJSON(r, pagetaxonomy); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := pagetaxonomy.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	pagetaxonomy.Prepare()

	if err := pagetaxonomy.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "page_taxonomy", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	pagetaxonomy, _, err = dao.UpdatePageTaxonomy(ctx,
		argId,
		pagetaxonomy)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, pagetaxonomy)
}

// DeletePageTaxonomy Delete a single record from page_taxonomy table in the laravel database
// @Summary Delete a record from page_taxonomy
// @Description Delete a single record from page_taxonomy table in the laravel database
// @Tags PageTaxonomy
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.PageTaxonomy
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /pagetaxonomy/{argId} [delete]
// http DELETE "http://localhost:8080/pagetaxonomy/1" X-Api-User:user123
func DeletePageTaxonomy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "page_taxonomy", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePageTaxonomy(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
