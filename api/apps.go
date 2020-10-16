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

func configAppsRouter(router *httprouter.Router) {
	router.GET("/apps", GetAllApps)
	router.POST("/apps", AddApps)
	router.GET("/apps/:argId", GetApps)
	router.PUT("/apps/:argId", UpdateApps)
	router.DELETE("/apps/:argId", DeleteApps)
}

func configGinAppsRouter(router gin.IRoutes) {
	router.GET("/apps", ConverHttprouterToGin(GetAllApps))
	router.POST("/apps", ConverHttprouterToGin(AddApps))
	router.GET("/apps/:argId", ConverHttprouterToGin(GetApps))
	router.PUT("/apps/:argId", ConverHttprouterToGin(UpdateApps))
	router.DELETE("/apps/:argId", ConverHttprouterToGin(DeleteApps))
}

// GetAllApps is a function to get a slice of record(s) from apps table in the laravel database
// @Summary Get list of Apps
// @Tags Apps
// @Description GetAllApps is a handler to get a slice of record(s) from apps table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Apps}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /apps [get]
// http "http://localhost:8080/apps?page=0&pagesize=20" X-Api-User:user123
func GetAllApps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "apps", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllApps(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetApps is a function to get a single record from the apps table in the laravel database
// @Summary Get record from table Apps by  argId
// @Tags Apps
// @ID argId
// @Description GetApps is a function to get a single record from the apps table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.Apps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /apps/{argId} [get]
// http "http://localhost:8080/apps/1" X-Api-User:user123
func GetApps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "apps", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetApps(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddApps add to add a single record to apps table in the laravel database
// @Summary Add an record to apps table
// @Description add to add a single record to apps table in the laravel database
// @Tags Apps
// @Accept  json
// @Produce  json
// @Param Apps body model.Apps true "Add Apps"
// @Success 200 {object} model.Apps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /apps [post]
// echo '{"logo": "ckwGoYUTtMTvBdAHDwkVSEOSo","created_at": "2231-05-26T06:10:00.114676915-03:00","updated_at": "2221-07-20T08:19:11.827487845-03:00","is_store": 31,"id": 45,"name": "CGssXavdjSdXyYrvNhqhpduCs","slug": "FdPmCTJJlIumZrLqLHejaMZUj","description": "JZtteaRuFVkcPpiEeTiyKNFCM","is_content": 34}' | http POST "http://localhost:8080/apps" X-Api-User:user123
func AddApps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	apps := &model.Apps{}

	if err := readJSON(r, apps); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := apps.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	apps.Prepare()

	if err := apps.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "apps", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	apps, _, err = dao.AddApps(ctx, apps)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, apps)
}

// UpdateApps Update a single record from apps table in the laravel database
// @Summary Update an record in table apps
// @Description Update a single record from apps table in the laravel database
// @Tags Apps
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  Apps body model.Apps true "Update Apps record"
// @Success 200 {object} model.Apps
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /apps/{argId} [put]
// echo '{"logo": "ckwGoYUTtMTvBdAHDwkVSEOSo","created_at": "2231-05-26T06:10:00.114676915-03:00","updated_at": "2221-07-20T08:19:11.827487845-03:00","is_store": 31,"id": 45,"name": "CGssXavdjSdXyYrvNhqhpduCs","slug": "FdPmCTJJlIumZrLqLHejaMZUj","description": "JZtteaRuFVkcPpiEeTiyKNFCM","is_content": 34}' | http PUT "http://localhost:8080/apps/1"  X-Api-User:user123
func UpdateApps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	apps := &model.Apps{}
	if err := readJSON(r, apps); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := apps.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	apps.Prepare()

	if err := apps.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "apps", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	apps, _, err = dao.UpdateApps(ctx,
		argId,
		apps)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, apps)
}

// DeleteApps Delete a single record from apps table in the laravel database
// @Summary Delete a record from apps
// @Description Delete a single record from apps table in the laravel database
// @Tags Apps
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.Apps
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /apps/{argId} [delete]
// http DELETE "http://localhost:8080/apps/1" X-Api-User:user123
func DeleteApps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "apps", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteApps(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
