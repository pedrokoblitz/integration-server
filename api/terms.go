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

func configTermsRouter(router *httprouter.Router) {
	router.GET("/terms", GetAllTerms)
	router.POST("/terms", AddTerms)
	router.GET("/terms/:argId", GetTerms)
	router.PUT("/terms/:argId", UpdateTerms)
	router.DELETE("/terms/:argId", DeleteTerms)
}

func configGinTermsRouter(router gin.IRoutes) {
	router.GET("/terms", ConverHttprouterToGin(GetAllTerms))
	router.POST("/terms", ConverHttprouterToGin(AddTerms))
	router.GET("/terms/:argId", ConverHttprouterToGin(GetTerms))
	router.PUT("/terms/:argId", ConverHttprouterToGin(UpdateTerms))
	router.DELETE("/terms/:argId", ConverHttprouterToGin(DeleteTerms))
}

// GetAllTerms is a function to get a slice of record(s) from terms table in the laravel database
// @Summary Get list of Terms
// @Tags Terms
// @Description GetAllTerms is a handler to get a slice of record(s) from terms table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Terms}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /terms [get]
// http "http://localhost:8080/terms?page=0&pagesize=20" X-Api-User:user123
func GetAllTerms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "terms", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllTerms(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetTerms is a function to get a single record from the terms table in the laravel database
// @Summary Get record from table Terms by  argId
// @Tags Terms
// @ID argId
// @Description GetTerms is a function to get a single record from the terms table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 200 {object} model.Terms
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /terms/{argId} [get]
// http "http://localhost:8080/terms/1" X-Api-User:user123
func GetTerms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "terms", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTerms(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddTerms add to add a single record to terms table in the laravel database
// @Summary Add an record to terms table
// @Description add to add a single record to terms table in the laravel database
// @Tags Terms
// @Accept  json
// @Produce  json
// @Param Terms body model.Terms true "Add Terms"
// @Success 200 {object} model.Terms
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /terms [post]
// echo '{"created_at": "2256-03-22T15:48:05.280801806-03:00","updated_at": "2250-09-21T21:15:20.681197925-03:00","id": 80,"slug": "OuLmGQpUZDMZtXptiwIyEpIcJ","app_id": 84,"type_id": 2,"term": "mKwgjeygPNhFxokrakVnpcGJf"}' | http POST "http://localhost:8080/terms" X-Api-User:user123
func AddTerms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	terms := &model.Terms{}

	if err := readJSON(r, terms); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := terms.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	terms.Prepare()

	if err := terms.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "terms", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	terms, _, err = dao.AddTerms(ctx, terms)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, terms)
}

// UpdateTerms Update a single record from terms table in the laravel database
// @Summary Update an record in table terms
// @Description Update a single record from terms table in the laravel database
// @Tags Terms
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Param  Terms body model.Terms true "Update Terms record"
// @Success 200 {object} model.Terms
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /terms/{argId} [put]
// echo '{"created_at": "2256-03-22T15:48:05.280801806-03:00","updated_at": "2250-09-21T21:15:20.681197925-03:00","id": 80,"slug": "OuLmGQpUZDMZtXptiwIyEpIcJ","app_id": 84,"type_id": 2,"term": "mKwgjeygPNhFxokrakVnpcGJf"}' | http PUT "http://localhost:8080/terms/1"  X-Api-User:user123
func UpdateTerms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	terms := &model.Terms{}
	if err := readJSON(r, terms); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := terms.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	terms.Prepare()

	if err := terms.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "terms", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	terms, _, err = dao.UpdateTerms(ctx,
		argId,
		terms)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, terms)
}

// DeleteTerms Delete a single record from terms table in the laravel database
// @Summary Delete a record from terms
// @Description Delete a single record from terms table in the laravel database
// @Tags Terms
// @Accept  json
// @Produce  json
// @Param  argId path int true "id"
// @Success 204 {object} model.Terms
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /terms/{argId} [delete]
// http DELETE "http://localhost:8080/terms/1" X-Api-User:user123
func DeleteTerms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint32(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "terms", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteTerms(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
