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

func configPagesRouter(router *httprouter.Router) {
	router.GET("/pages", GetAllPages)
	router.POST("/pages", AddPages)
	router.GET("/pages/:argId", GetPages)
	router.PUT("/pages/:argId", UpdatePages)
	router.DELETE("/pages/:argId", DeletePages)
}

func configGinPagesRouter(router gin.IRoutes) {
	router.GET("/pages", ConverHttprouterToGin(GetAllPages))
	router.POST("/pages", ConverHttprouterToGin(AddPages))
	router.GET("/pages/:argId", ConverHttprouterToGin(GetPages))
	router.PUT("/pages/:argId", ConverHttprouterToGin(UpdatePages))
	router.DELETE("/pages/:argId", ConverHttprouterToGin(DeletePages))
}

// GetAllPages is a function to get a slice of record(s) from pages table in the laravel database
// @Summary Get list of Pages
// @Tags Pages
// @Description GetAllPages is a handler to get a slice of record(s) from pages table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Pages}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pages [get]
// http "http://localhost:8080/pages?page=0&pagesize=20" X-Api-User:user123
func GetAllPages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "pages", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPages(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPages is a function to get a single record from the pages table in the laravel database
// @Summary Get record from table Pages by  argId
// @Tags Pages
// @ID argId
// @Description GetPages is a function to get a single record from the pages table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.Pages
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /pages/{argId} [get]
// http "http://localhost:8080/pages/1" X-Api-User:user123
func GetPages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "pages", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPages(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPages add to add a single record to pages table in the laravel database
// @Summary Add an record to pages table
// @Description add to add a single record to pages table in the laravel database
// @Tags Pages
// @Accept  json
// @Produce  json
// @Param Pages body model.Pages true "Add Pages"
// @Success 200 {object} model.Pages
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pages [post]
// echo '{"parent_id": 36,"css": "RYKQVIdxLnghdGStVdBmilsoJ","created_at": "2261-02-23T15:14:49.85064681-03:00","updated_at": "2120-07-07T05:00:44.103625424-03:00","id": 58,"type_id": 59,"meta_title": "mEjpZTDGEAFMZdGMYHNuPCMsn","meta_keywords": "YBpGaWLDxYdOJjBRLRodLlaoj","app_id": 65,"slug": "BInbuVLVwenhMlUabpqNsUviB","status": "rWkblgVQMUSdNRavCsdnRikeX","user_id": 93,"meta_image": "KBMctWDYsTUwlDgNMNcnIRiXG","meta_description": "tTgABcgFJbgAUJqVrNfrjbsHe"}' | http POST "http://localhost:8080/pages" X-Api-User:user123
func AddPages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	pages := &model.Pages{}

	if err := readJSON(r, pages); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := pages.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	pages.Prepare()

	if err := pages.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "pages", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	pages, _, err = dao.AddPages(ctx, pages)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, pages)
}

// UpdatePages Update a single record from pages table in the laravel database
// @Summary Update an record in table pages
// @Description Update a single record from pages table in the laravel database
// @Tags Pages
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  Pages body model.Pages true "Update Pages record"
// @Success 200 {object} model.Pages
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /pages/{argId} [put]
// echo '{"parent_id": 36,"css": "RYKQVIdxLnghdGStVdBmilsoJ","created_at": "2261-02-23T15:14:49.85064681-03:00","updated_at": "2120-07-07T05:00:44.103625424-03:00","id": 58,"type_id": 59,"meta_title": "mEjpZTDGEAFMZdGMYHNuPCMsn","meta_keywords": "YBpGaWLDxYdOJjBRLRodLlaoj","app_id": 65,"slug": "BInbuVLVwenhMlUabpqNsUviB","status": "rWkblgVQMUSdNRavCsdnRikeX","user_id": 93,"meta_image": "KBMctWDYsTUwlDgNMNcnIRiXG","meta_description": "tTgABcgFJbgAUJqVrNfrjbsHe"}' | http PUT "http://localhost:8080/pages/1"  X-Api-User:user123
func UpdatePages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	pages := &model.Pages{}
	if err := readJSON(r, pages); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := pages.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	pages.Prepare()

	if err := pages.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "pages", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	pages, _, err = dao.UpdatePages(ctx,
		argId,
		pages)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, pages)
}

// DeletePages Delete a single record from pages table in the laravel database
// @Summary Delete a record from pages
// @Description Delete a single record from pages table in the laravel database
// @Tags Pages
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.Pages
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /pages/{argId} [delete]
// http DELETE "http://localhost:8080/pages/1" X-Api-User:user123
func DeletePages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "pages", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePages(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
