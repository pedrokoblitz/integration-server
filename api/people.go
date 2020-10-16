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

func configPeopleRouter(router *httprouter.Router) {
	router.GET("/people", GetAllPeople)
	router.POST("/people", AddPeople)
	router.GET("/people/:argId", GetPeople)
	router.PUT("/people/:argId", UpdatePeople)
	router.DELETE("/people/:argId", DeletePeople)
}

func configGinPeopleRouter(router gin.IRoutes) {
	router.GET("/people", ConverHttprouterToGin(GetAllPeople))
	router.POST("/people", ConverHttprouterToGin(AddPeople))
	router.GET("/people/:argId", ConverHttprouterToGin(GetPeople))
	router.PUT("/people/:argId", ConverHttprouterToGin(UpdatePeople))
	router.DELETE("/people/:argId", ConverHttprouterToGin(DeletePeople))
}

// GetAllPeople is a function to get a slice of record(s) from people table in the laravel database
// @Summary Get list of People
// @Tags People
// @Description GetAllPeople is a handler to get a slice of record(s) from people table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.People}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /people [get]
// http "http://localhost:8080/people?page=0&pagesize=20" X-Api-User:user123
func GetAllPeople(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "people", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPeople(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetPeople is a function to get a single record from the people table in the laravel database
// @Summary Get record from table People by  argId
// @Tags People
// @ID argId
// @Description GetPeople is a function to get a single record from the people table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.People
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /people/{argId} [get]
// http "http://localhost:8080/people/1" X-Api-User:user123
func GetPeople(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "people", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPeople(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPeople add to add a single record to people table in the laravel database
// @Summary Add an record to people table
// @Description add to add a single record to people table in the laravel database
// @Tags People
// @Accept  json
// @Produce  json
// @Param People body model.People true "Add People"
// @Success 200 {object} model.People
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /people [post]
// echo '{"phone": "WJsqVKIiRmbvTpcUTmfMxysXx","twitter": "DiBTTWjZHXNxtaIepCEEystts","optin": "2110-05-25T15:49:04.578918682-03:00","created_at": "2096-04-06T20:52:13.91492095-03:00","email": "dbTBIVlnpIQFFEvOmWkmVOqSg","birthdate": "2118-02-05T09:34:36.783003302-03:00","optout": "2215-07-31T17:52:57.497178228-03:00","id": 92,"updated_at": "2307-11-27T11:53:49.820816773-03:00","name": "YXEQJoGgfHJBxjXpjeJSkVfjR","instagram": "QBRsrpYqKgayLtqTqbvPWtgvh","facebook": "AcpIwWUyaflawJuBNiWDaKsNL","linkedin": "UqGrnLYkwlQSDXamMVipGudVv","last_campaing_sent": "2202-09-20T07:53:14.221142891-03:00"}' | http POST "http://localhost:8080/people" X-Api-User:user123
func AddPeople(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	people := &model.People{}

	if err := readJSON(r, people); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := people.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	people.Prepare()

	if err := people.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "people", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	people, _, err = dao.AddPeople(ctx, people)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, people)
}

// UpdatePeople Update a single record from people table in the laravel database
// @Summary Update an record in table people
// @Description Update a single record from people table in the laravel database
// @Tags People
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  People body model.People true "Update People record"
// @Success 200 {object} model.People
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /people/{argId} [put]
// echo '{"phone": "WJsqVKIiRmbvTpcUTmfMxysXx","twitter": "DiBTTWjZHXNxtaIepCEEystts","optin": "2110-05-25T15:49:04.578918682-03:00","created_at": "2096-04-06T20:52:13.91492095-03:00","email": "dbTBIVlnpIQFFEvOmWkmVOqSg","birthdate": "2118-02-05T09:34:36.783003302-03:00","optout": "2215-07-31T17:52:57.497178228-03:00","id": 92,"updated_at": "2307-11-27T11:53:49.820816773-03:00","name": "YXEQJoGgfHJBxjXpjeJSkVfjR","instagram": "QBRsrpYqKgayLtqTqbvPWtgvh","facebook": "AcpIwWUyaflawJuBNiWDaKsNL","linkedin": "UqGrnLYkwlQSDXamMVipGudVv","last_campaing_sent": "2202-09-20T07:53:14.221142891-03:00"}' | http PUT "http://localhost:8080/people/1"  X-Api-User:user123
func UpdatePeople(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	people := &model.People{}
	if err := readJSON(r, people); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := people.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	people.Prepare()

	if err := people.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "people", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	people, _, err = dao.UpdatePeople(ctx,
		argId,
		people)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, people)
}

// DeletePeople Delete a single record from people table in the laravel database
// @Summary Delete a record from people
// @Description Delete a single record from people table in the laravel database
// @Tags People
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.People
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /people/{argId} [delete]
// http DELETE "http://localhost:8080/people/1" X-Api-User:user123
func DeletePeople(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "people", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePeople(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
