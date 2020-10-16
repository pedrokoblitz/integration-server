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

func configBlocksRouter(router *httprouter.Router) {
	router.GET("/blocks", GetAllBlocks)
	router.POST("/blocks", AddBlocks)
	router.GET("/blocks/:argId", GetBlocks)
	router.PUT("/blocks/:argId", UpdateBlocks)
	router.DELETE("/blocks/:argId", DeleteBlocks)
}

func configGinBlocksRouter(router gin.IRoutes) {
	router.GET("/blocks", ConverHttprouterToGin(GetAllBlocks))
	router.POST("/blocks", ConverHttprouterToGin(AddBlocks))
	router.GET("/blocks/:argId", ConverHttprouterToGin(GetBlocks))
	router.PUT("/blocks/:argId", ConverHttprouterToGin(UpdateBlocks))
	router.DELETE("/blocks/:argId", ConverHttprouterToGin(DeleteBlocks))
}

// GetAllBlocks is a function to get a slice of record(s) from blocks table in the laravel database
// @Summary Get list of Blocks
// @Tags Blocks
// @Description GetAllBlocks is a handler to get a slice of record(s) from blocks table in the laravel database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Blocks}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blocks [get]
// http "http://localhost:8080/blocks?page=0&pagesize=20" X-Api-User:user123
func GetAllBlocks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "blocks", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllBlocks(ctx, page, pagesize, order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: totalRows}
	writeJSON(ctx, w, result)
}

// GetBlocks is a function to get a single record from the blocks table in the laravel database
// @Summary Get record from table Blocks by  argId
// @Tags Blocks
// @ID argId
// @Description GetBlocks is a function to get a single record from the blocks table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 200 {object} model.Blocks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /blocks/{argId} [get]
// http "http://localhost:8080/blocks/1" X-Api-User:user123
func GetBlocks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blocks", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetBlocks(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddBlocks add to add a single record to blocks table in the laravel database
// @Summary Add an record to blocks table
// @Description add to add a single record to blocks table in the laravel database
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param Blocks body model.Blocks true "Add Blocks"
// @Success 200 {object} model.Blocks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blocks [post]
// echo '{"app_id": 93,"date_2": "2163-01-04T09:03:26.129579507-03:00","file_2": "jeXKOvWLVoLuQoutFCekJfwOj","line_3": "XVitSbsHOlThfrtLLuycVhCkc","embed": "amtTQivLmpeuJrwKEHWclcxQc","type_id": 11,"order": 47,"item_id": 39,"text_2": "KqnhoBTBVcnjcpUkJDAeFNavJ","page_id": 22,"grid_size": 44,"grid_size_mobile": 93,"css": "YdXQNrkRIOYpJoWrHfSigbxpq","url": "nmaxrKbYegrrUtndurhyKTFIJ","view": "JEbjsQsGGaLdUHgWkMygRNhEg","line_1": "LsQBRnbZJRiVTRYQKuyOcmxWT","line_2": "NESXUBrepnMCGJNLtXiqfCYug","date_1": "2309-07-24T10:20:01.169301386-03:00","term_id": 78,"file_1": "GIirvdUOCpIsdOcUrrJnhqZvw","button_text": "iLdujlMdjTswCmPjfYHipEaUX","created_at": "2076-06-02T14:14:55.49864723-03:00","id": 16,"parent_id": 39,"limit": 42,"text_1": "LddnLQADhUtgTjSVbdFhrkoYM","updated_at": "2271-10-18T11:32:55.648692566-03:00"}' | http POST "http://localhost:8080/blocks" X-Api-User:user123
func AddBlocks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	blocks := &model.Blocks{}

	if err := readJSON(r, blocks); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blocks.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blocks.Prepare()

	if err := blocks.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blocks", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	blocks, _, err = dao.AddBlocks(ctx, blocks)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blocks)
}

// UpdateBlocks Update a single record from blocks table in the laravel database
// @Summary Update an record in table blocks
// @Description Update a single record from blocks table in the laravel database
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Param  Blocks body model.Blocks true "Update Blocks record"
// @Success 200 {object} model.Blocks
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /blocks/{argId} [put]
// echo '{"app_id": 93,"date_2": "2163-01-04T09:03:26.129579507-03:00","file_2": "jeXKOvWLVoLuQoutFCekJfwOj","line_3": "XVitSbsHOlThfrtLLuycVhCkc","embed": "amtTQivLmpeuJrwKEHWclcxQc","type_id": 11,"order": 47,"item_id": 39,"text_2": "KqnhoBTBVcnjcpUkJDAeFNavJ","page_id": 22,"grid_size": 44,"grid_size_mobile": 93,"css": "YdXQNrkRIOYpJoWrHfSigbxpq","url": "nmaxrKbYegrrUtndurhyKTFIJ","view": "JEbjsQsGGaLdUHgWkMygRNhEg","line_1": "LsQBRnbZJRiVTRYQKuyOcmxWT","line_2": "NESXUBrepnMCGJNLtXiqfCYug","date_1": "2309-07-24T10:20:01.169301386-03:00","term_id": 78,"file_1": "GIirvdUOCpIsdOcUrrJnhqZvw","button_text": "iLdujlMdjTswCmPjfYHipEaUX","created_at": "2076-06-02T14:14:55.49864723-03:00","id": 16,"parent_id": 39,"limit": 42,"text_1": "LddnLQADhUtgTjSVbdFhrkoYM","updated_at": "2271-10-18T11:32:55.648692566-03:00"}' | http PUT "http://localhost:8080/blocks/1"  X-Api-User:user123
func UpdateBlocks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blocks := &model.Blocks{}
	if err := readJSON(r, blocks); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := blocks.BeforeSave(); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	blocks.Prepare()

	if err := blocks.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "blocks", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	blocks, _, err = dao.UpdateBlocks(ctx,
		argId,
		blocks)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, blocks)
}

// DeleteBlocks Delete a single record from blocks table in the laravel database
// @Summary Delete a record from blocks
// @Description Delete a single record from blocks table in the laravel database
// @Tags Blocks
// @Accept  json
// @Produce  json
// @Param  argId path int64 true "id"
// @Success 204 {object} model.Blocks
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /blocks/{argId} [delete]
// http DELETE "http://localhost:8080/blocks/1" X-Api-User:user123
func DeleteBlocks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argId, err := parseUint64(ps, "argId")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "blocks", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteBlocks(ctx, argId)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
