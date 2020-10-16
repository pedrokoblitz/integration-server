package api

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/pedrokoblitz/comm/dao"
	"github.com/pedrokoblitz/comm/model"

	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

var (
	_             = time.Second // import time.Second for unknown usage in api
	crudEndpoints map[string]*CrudAPI
)

// CrudAPI describes requests available for tables in the database
type CrudAPI struct {
	Name            string           `json:"name"`
	CreateURL       string           `json:"create_url"`
	RetrieveOneURL  string           `json:"retrieve_one_url"`
	RetrieveManyURL string           `json:"retrieve_many_url"`
	UpdateURL       string           `json:"update_url"`
	DeleteURL       string           `json:"delete_url"`
	FetchDDLURL     string           `json:"fetch_ddl_url"`
	TableInfo       *model.TableInfo `json:"table_info"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ConfigRouter configure http.Handler router
func ConfigRouter() http.Handler {
	router := httprouter.New()
	configAppsRouter(router)
	configBlockTypesRouter(router)
	configBlocksRouter(router)
	configCampaingsRouter(router)
	configPageTaxonomyRouter(router)
	configPageTypesRouter(router)
	configPagesRouter(router)
	configPeopleRouter(router)
	configPeopleTaxonomyRouter(router)
	configPersonActionsRouter(router)
	configPersonHistoryRouter(router)
	configTermTypesRouter(router)
	configTermsRouter(router)

	router.GET("/ddl/:argID", GetDdl)
	router.GET("/ddl", GetDdlEndpoints)
	return router
}

// ConfigGinRouter configure gin router
func ConfigGinRouter(router gin.IRoutes) {
	configGinAppsRouter(router)
	configGinBlockTypesRouter(router)
	configGinBlocksRouter(router)
	configGinCampaingsRouter(router)
	configGinPageTaxonomyRouter(router)
	configGinPageTypesRouter(router)
	configGinPagesRouter(router)
	configGinPeopleRouter(router)
	configGinPeopleTaxonomyRouter(router)
	configGinPersonActionsRouter(router)
	configGinPersonHistoryRouter(router)
	configGinTermTypesRouter(router)
	configGinTermsRouter(router)

	router.GET("/ddl/:argID", ConverHttprouterToGin(GetDdl))
	router.GET("/ddl", ConverHttprouterToGin(GetDdlEndpoints))
	return
}

// ConverHttprouterToGin wrap httprouter.Handle to gin.HandlerFunc
func ConverHttprouterToGin(f httprouter.Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params httprouter.Params
		_len := len(c.Params)
		if _len == 0 {
			params = nil
		} else {
			params = ((*[1 << 10]httprouter.Param)(unsafe.Pointer(&c.Params[0])))[:_len]
		}

		f(c.Writer, c.Request, params)
	}
}

func initializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var RequestValidator RequestValidatorFunc

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var ContextInitializer ContextInitializerFunc

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}

	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func writeRowsAffected(w http.ResponseWriter, rowsAffected int64) {
	data, _ := json.Marshal(rowsAffected)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func returnError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	status := 0
	switch err {
	case dao.ErrNotFound:
		status = http.StatusBadRequest
	case dao.ErrUnableToMarshalJSON:
		status = http.StatusBadRequest
	case dao.ErrUpdateFailed:
		status = http.StatusBadRequest
	case dao.ErrInsertFailed:
		status = http.StatusBadRequest
	case dao.ErrDeleteFailed:
		status = http.StatusBadRequest
	case dao.ErrBadParams:
		status = http.StatusBadRequest
	default:
		status = http.StatusBadRequest
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	SendJSON(w, r, er.Code, er)
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

func parseUint8(ps httprouter.Params, key string) (uint8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return uint8(id), err
	}
	return uint8(id), err
}
func parseUint16(ps httprouter.Params, key string) (uint16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return uint16(id), err
	}
	return uint16(id), err
}
func parseUint32(ps httprouter.Params, key string) (uint32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return uint32(id), err
	}
	return uint32(id), err
}
func parseUint64(ps httprouter.Params, key string) (uint64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), err
}
func parseInt(ps httprouter.Params, key string) (int, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return -1, err
	}
	return int(id), err
}
func parseInt8(ps httprouter.Params, key string) (int8, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 8)
	if err != nil {
		return -1, err
	}
	return int8(id), err
}
func parseInt16(ps httprouter.Params, key string) (int16, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil {
		return -1, err
	}
	return int16(id), err
}
func parseInt32(ps httprouter.Params, key string) (int32, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return -1, err
	}
	return int32(id), err
}
func parseInt64(ps httprouter.Params, key string) (int64, error) {
	idStr := ps.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 54)
	if err != nil {
		return -1, err
	}
	return id, err
}
func parseString(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}
func parseUUID(ps httprouter.Params, key string) (string, error) {
	idStr := ps.ByName(key)
	return idStr, nil
}

// GetDdl is a function to get table info for a table in the laravel database
// @Summary Get table info for a table in the laravel database by argID
// @Tags TableInfo
// @ID argID
// @Description GetDdl is a function to get table info for a table in the laravel database
// @Accept  json
// @Produce  json
// @Param  argID path int true "id"
// @Success 200 {object} api.CrudAPI
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /ddl/{argID} [get]
// http "http://localhost:8080/ddl/xyz" X-Api-User:user123
func GetDdl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID := ps.ByName("argID")

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, ok := crudEndpoints[argID]
	if !ok {
		returnError(ctx, w, r, fmt.Errorf("unable to find table: %s", argID))
		return
	}

	writeJSON(ctx, w, record)
}

// GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the laravel database
// @Summary Gets a list of ddl endpoints available for tables in the laravel database
// @Tags TableInfo
// @Description GetDdlEndpoints is a function to get a list of ddl endpoints available for tables in the laravel database
// @Accept  json
// @Produce  json
// @Success 200 {object} api.CrudAPI
// @Router /ddl [get]
// http "http://localhost:8080/ddl" X-Api-User:user123
func GetDdlEndpoints(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "ddl", model.FetchDDL); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, crudEndpoints)
}

func init() {
	crudEndpoints = make(map[string]*CrudAPI)

	var tmp *CrudAPI

	tmp = &CrudAPI{
		Name:            "apps",
		CreateURL:       "/apps",
		RetrieveOneURL:  "/apps",
		RetrieveManyURL: "/apps",
		UpdateURL:       "/apps",
		DeleteURL:       "/apps",
		FetchDDLURL:     "/ddl/apps",
	}

	tmp.TableInfo, _ = model.GetTableInfo("apps")
	crudEndpoints["apps"] = tmp

	tmp = &CrudAPI{
		Name:            "block_types",
		CreateURL:       "/blocktypes",
		RetrieveOneURL:  "/blocktypes",
		RetrieveManyURL: "/blocktypes",
		UpdateURL:       "/blocktypes",
		DeleteURL:       "/blocktypes",
		FetchDDLURL:     "/ddl/block_types",
	}

	tmp.TableInfo, _ = model.GetTableInfo("block_types")
	crudEndpoints["block_types"] = tmp

	tmp = &CrudAPI{
		Name:            "blocks",
		CreateURL:       "/blocks",
		RetrieveOneURL:  "/blocks",
		RetrieveManyURL: "/blocks",
		UpdateURL:       "/blocks",
		DeleteURL:       "/blocks",
		FetchDDLURL:     "/ddl/blocks",
	}

	tmp.TableInfo, _ = model.GetTableInfo("blocks")
	crudEndpoints["blocks"] = tmp

	tmp = &CrudAPI{
		Name:            "campaings",
		CreateURL:       "/campaings",
		RetrieveOneURL:  "/campaings",
		RetrieveManyURL: "/campaings",
		UpdateURL:       "/campaings",
		DeleteURL:       "/campaings",
		FetchDDLURL:     "/ddl/campaings",
	}

	tmp.TableInfo, _ = model.GetTableInfo("campaings")
	crudEndpoints["campaings"] = tmp

	tmp = &CrudAPI{
		Name:            "page_taxonomy",
		CreateURL:       "/pagetaxonomy",
		RetrieveOneURL:  "/pagetaxonomy",
		RetrieveManyURL: "/pagetaxonomy",
		UpdateURL:       "/pagetaxonomy",
		DeleteURL:       "/pagetaxonomy",
		FetchDDLURL:     "/ddl/page_taxonomy",
	}

	tmp.TableInfo, _ = model.GetTableInfo("page_taxonomy")
	crudEndpoints["page_taxonomy"] = tmp

	tmp = &CrudAPI{
		Name:            "page_types",
		CreateURL:       "/pagetypes",
		RetrieveOneURL:  "/pagetypes",
		RetrieveManyURL: "/pagetypes",
		UpdateURL:       "/pagetypes",
		DeleteURL:       "/pagetypes",
		FetchDDLURL:     "/ddl/page_types",
	}

	tmp.TableInfo, _ = model.GetTableInfo("page_types")
	crudEndpoints["page_types"] = tmp

	tmp = &CrudAPI{
		Name:            "pages",
		CreateURL:       "/pages",
		RetrieveOneURL:  "/pages",
		RetrieveManyURL: "/pages",
		UpdateURL:       "/pages",
		DeleteURL:       "/pages",
		FetchDDLURL:     "/ddl/pages",
	}

	tmp.TableInfo, _ = model.GetTableInfo("pages")
	crudEndpoints["pages"] = tmp


	tmp = &CrudAPI{
		Name:            "people",
		CreateURL:       "/people",
		RetrieveOneURL:  "/people",
		RetrieveManyURL: "/people",
		UpdateURL:       "/people",
		DeleteURL:       "/people",
		FetchDDLURL:     "/ddl/people",
	}

	tmp.TableInfo, _ = model.GetTableInfo("people")
	crudEndpoints["people"] = tmp

	tmp = &CrudAPI{
		Name:            "people_taxonomy",
		CreateURL:       "/peopletaxonomy",
		RetrieveOneURL:  "/peopletaxonomy",
		RetrieveManyURL: "/peopletaxonomy",
		UpdateURL:       "/peopletaxonomy",
		DeleteURL:       "/peopletaxonomy",
		FetchDDLURL:     "/ddl/people_taxonomy",
	}

	tmp.TableInfo, _ = model.GetTableInfo("people_taxonomy")
	crudEndpoints["people_taxonomy"] = tmp

	tmp = &CrudAPI{
		Name:            "person_actions",
		CreateURL:       "/personactions",
		RetrieveOneURL:  "/personactions",
		RetrieveManyURL: "/personactions",
		UpdateURL:       "/personactions",
		DeleteURL:       "/personactions",
		FetchDDLURL:     "/ddl/person_actions",
	}

	tmp.TableInfo, _ = model.GetTableInfo("person_actions")
	crudEndpoints["person_actions"] = tmp

	tmp = &CrudAPI{
		Name:            "person_history",
		CreateURL:       "/personhistory",
		RetrieveOneURL:  "/personhistory",
		RetrieveManyURL: "/personhistory",
		UpdateURL:       "/personhistory",
		DeleteURL:       "/personhistory",
		FetchDDLURL:     "/ddl/person_history",
	}

	tmp.TableInfo, _ = model.GetTableInfo("person_history")
	crudEndpoints["person_history"] = tmp

	tmp = &CrudAPI{
		Name:            "prices",
		CreateURL:       "/prices",
		RetrieveOneURL:  "/prices",
		RetrieveManyURL: "/prices",
		UpdateURL:       "/prices",
		DeleteURL:       "/prices",
		FetchDDLURL:     "/ddl/prices",
	}

	tmp = &CrudAPI{
		Name:            "term_types",
		CreateURL:       "/termtypes",
		RetrieveOneURL:  "/termtypes",
		RetrieveManyURL: "/termtypes",
		UpdateURL:       "/termtypes",
		DeleteURL:       "/termtypes",
		FetchDDLURL:     "/ddl/term_types",
	}

	tmp.TableInfo, _ = model.GetTableInfo("term_types")
	crudEndpoints["term_types"] = tmp

	tmp = &CrudAPI{
		Name:            "terms",
		CreateURL:       "/terms",
		RetrieveOneURL:  "/terms",
		RetrieveManyURL: "/terms",
		UpdateURL:       "/terms",
		DeleteURL:       "/terms",
		FetchDDLURL:     "/ddl/terms",
	}

}
