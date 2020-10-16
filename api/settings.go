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

func configSettingsRouter(router *httprouter.Router) {
	router.GET("/settings", GetAllSettings)
	router.POST("/settings", AddSettings)
	router.GET("/settings/:argId", GetSettings)
	router.PUT("/settings/:argId", UpdateSettings)
	router.DELETE("/settings/:argId", DeleteSettings)
}

func configGinSettingsRouter(router gin.IRoutes) {
	router.GET("/settings", ConverHttprouterToGin(GetAllSettings))
	router.POST("/settings", ConverHttprouterToGin(AddSettings))
	router.GET("/settings/:argId", ConverHttprouterToGin(GetSettings))
	router.PUT("/settings/:argId", ConverHttprouterToGin(UpdateSettings))
	router.DELETE("/settings/:argId", ConverHttprouterToGin(DeleteSettings))
}

// GetAllSettings is a function to get a slice of record(s) from settings table in the laravel database
// @Summary Get list of Settings
// @Tags Settings
// @Description GetAllSettings is a handler to get a slice of record(s) from settings table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Settings}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /settings [get]
// http "http://localhost:8080/settings?page=0&pagesize=20" X-Api-User:user123
func GetAllSettings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "settings", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllSettings(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetSettings is a function to get a single record from the settings table in the laravel database
// @Summary Get record from table Settings by  argId
// @Tags Settings
// @ID argId
// @Description GetSettings is a function to get a single record from the settings table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.Settings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /settings/{argId} [get]
// http "http://localhost:8080/settings/1" X-Api-User:user123
func GetSettings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "settings", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetSettings(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddSettings add to add a single record to settings table in the laravel database
// @Summary Add an record to settings table
// @Description add to add a single record to settings table in the laravel database
// @Tags Settings
// @Accept  json
// @Produce  json
// @Param Settings body model.Settings true "Add Settings"
// @Success 200 {object} model.Settings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /settings [post]
// echo '{"created_at": "2188-04-13T13:35:09.151575676-03:00","updated_at": "2121-05-27T22:45:45.254191157-03:00","id": 73,"key": "cUPcnunVtsdeloOQCOeZQidZH","value": "GruQusCfsYeAIfVXvpidVdOxX"}' | http POST "http://localhost:8080/settings" X-Api-User:user123
func AddSettings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	settings := &model.Settings{}

	if err := readJSON(r, settings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := settings.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	settings.Prepare()

	if err := settings.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "settings", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	settings, _, err = dao.AddSettings(ctx, settings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, settings)
}

// UpdateSettings Update a single record from settings table in the laravel database
// @Summary Update an record in table settings
// @Description Update a single record from settings table in the laravel database
// @Tags Settings
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  Settings body model.Settings true "Update Settings record"
// @Success 200 {object} model.Settings
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /settings/{argId} [put]
// echo '{"created_at": "2188-04-13T13:35:09.151575676-03:00","updated_at": "2121-05-27T22:45:45.254191157-03:00","id": 73,"key": "cUPcnunVtsdeloOQCOeZQidZH","value": "GruQusCfsYeAIfVXvpidVdOxX"}' | http PUT "http://localhost:8080/settings/1"  X-Api-User:user123
func UpdateSettings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	settings := &model.Settings{}
	if err := readJSON(r, settings); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := settings.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	settings.Prepare()

	if err := settings.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "settings", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	settings, _, err = dao.UpdateSettings(ctx,
		argId,
		settings)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, settings)
}

// DeleteSettings Delete a single record from settings table in the laravel database
// @Summary Delete a record from settings
// @Description Delete a single record from settings table in the laravel database
// @Tags Settings
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.Settings
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /settings/{argId} [delete]
// http DELETE "http://localhost:8080/settings/1" X-Api-User:user123
func DeleteSettings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "settings", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteSettings(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
