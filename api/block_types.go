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

func configBlockTypesRouter(router *httprouter.Router) {
	router.GET("/blocktypes", GetAllBlockTypes)
	router.POST("/blocktypes", AddBlockTypes)
	router.GET("/blocktypes/:argId", GetBlockTypes)
	router.PUT("/blocktypes/:argId", UpdateBlockTypes)
	router.DELETE("/blocktypes/:argId", DeleteBlockTypes)
}

func configGinBlockTypesRouter(router gin.IRoutes) {
	router.GET("/blocktypes", ConverHttprouterToGin(GetAllBlockTypes))
	router.POST("/blocktypes", ConverHttprouterToGin(AddBlockTypes))
	router.GET("/blocktypes/:argId", ConverHttprouterToGin(GetBlockTypes))
	router.PUT("/blocktypes/:argId", ConverHttprouterToGin(UpdateBlockTypes))
	router.DELETE("/blocktypes/:argId", ConverHttprouterToGin(DeleteBlockTypes))
}

// GetAllBlockTypes is a function to get a slice of record(s) from block_types table in the laravel database
// @Summary Get list of BlockTypes
// @Tags BlockTypes
// @Description GetAllBlockTypes is a handler to get a slice of record(s) from block_types table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.BlockTypes}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blocktypes [get]
// http "http://localhost:8080/blocktypes?page=0&pagesize=20" X-Api-User:user123
func GetAllBlockTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "block_types", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlockTypes(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlockTypes is a function to get a single record from the block_types table in the laravel database
// @Summary Get record from table BlockTypes by  argId
// @Tags BlockTypes
// @ID argId
// @Description GetBlockTypes is a function to get a single record from the block_types table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.BlockTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blocktypes/{argId} [get]
// http "http://localhost:8080/blocktypes/1" X-Api-User:user123
func GetBlockTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "block_types", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlockTypes(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlockTypes add to add a single record to block_types table in the laravel database
// @Summary Add an record to block_types table
// @Description add to add a single record to block_types table in the laravel database
// @Tags BlockTypes
// @Accept  json
// @Produce  json
// @Param BlockTypes body model.BlockTypes true "Add BlockTypes"
// @Success 200 {object} model.BlockTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blocktypes [post]
// echo '{"is_module": 52,"is_page_term": 82,"is_product_term": 39,"is_page": 48,"is_product": 35,"name": "VjGAvNCvONoWBBeCFdRtHwMuu","slug": "seWLFtQXDVUNjiEIinZmhFmpQ","app_id": 31,"created_at": "2300-01-22T21:38:52.284666579-03:00","updated_at": "2081-06-11T03:02:47.556453413-03:00","text_1": 96,"file_2": 12,"date_1": 80,"button_text": 55,"date_2": 50,"id": 89,"file_1": 0,"embed": 63,"line_2": 49,"line_3": 9,"text_2": 59,"url": 74,"is_parent": 7,"is_child": 49,"line_1": 74}' | http POST "http://localhost:8080/blocktypes" X-Api-User:user123
func AddBlockTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blocktypes := &model.BlockTypes{}

	if err := readJSON(r, blocktypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blocktypes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blocktypes.Prepare()

	if err := blocktypes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "block_types", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blocktypes, _, err = dao.AddBlockTypes(ctx, blocktypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blocktypes)
}

// UpdateBlockTypes Update a single record from block_types table in the laravel database
// @Summary Update an record in table block_types
// @Description Update a single record from block_types table in the laravel database
// @Tags BlockTypes
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  BlockTypes body model.BlockTypes true "Update BlockTypes record"
// @Success 200 {object} model.BlockTypes
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blocktypes/{argId} [put]
// echo '{"is_module": 52,"is_page_term": 82,"is_product_term": 39,"is_page": 48,"is_product": 35,"name": "VjGAvNCvONoWBBeCFdRtHwMuu","slug": "seWLFtQXDVUNjiEIinZmhFmpQ","app_id": 31,"created_at": "2300-01-22T21:38:52.284666579-03:00","updated_at": "2081-06-11T03:02:47.556453413-03:00","text_1": 96,"file_2": 12,"date_1": 80,"button_text": 55,"date_2": 50,"id": 89,"file_1": 0,"embed": 63,"line_2": 49,"line_3": 9,"text_2": 59,"url": 74,"is_parent": 7,"is_child": 49,"line_1": 74}' | http PUT "http://localhost:8080/blocktypes/1"  X-Api-User:user123
func UpdateBlockTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blocktypes := &model.BlockTypes{}
	if err := readJSON(r, blocktypes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blocktypes.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blocktypes.Prepare()

	if err := blocktypes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "block_types", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blocktypes, _, err = dao.UpdateBlockTypes(ctx,
		argId,
		blocktypes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blocktypes)
}

// DeleteBlockTypes Delete a single record from block_types table in the laravel database
// @Summary Delete a record from block_types
// @Description Delete a single record from block_types table in the laravel database
// @Tags BlockTypes
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.BlockTypes
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blocktypes/{argId} [delete]
// http DELETE "http://localhost:8080/blocktypes/1" X-Api-User:user123
func DeleteBlockTypes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "block_types", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlockTypes(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
