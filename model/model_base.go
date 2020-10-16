package model

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["address_types"] = address_typesTableInfo
	tables["addresses"] = addressesTableInfo
	tables["api_key_types"] = api_key_typesTableInfo
	tables["api_keys"] = api_keysTableInfo
	tables["apps"] = appsTableInfo
	tables["block_types"] = block_typesTableInfo
	tables["blocks"] = blocksTableInfo
	tables["campaings"] = campaingsTableInfo
	tables["clients"] = clientsTableInfo
	tables["clients_apps"] = clients_appsTableInfo
	tables["clients_users"] = clients_usersTableInfo
	tables["coupons"] = couponsTableInfo
	tables["customer_tracking"] = customer_trackingTableInfo
	tables["failed_jobs"] = failed_jobsTableInfo
	tables["jobs"] = jobsTableInfo
	tables["larametrics_logs"] = larametrics_logsTableInfo
	tables["larametrics_models"] = larametrics_modelsTableInfo
	tables["larametrics_notifications"] = larametrics_notificationsTableInfo
	tables["larametrics_requests"] = larametrics_requestsTableInfo
	tables["migrations"] = migrationsTableInfo
	tables["order_items"] = order_itemsTableInfo
	tables["orders"] = ordersTableInfo
	tables["page_taxonomy"] = page_taxonomyTableInfo
	tables["page_types"] = page_typesTableInfo
	tables["pages"] = pagesTableInfo
	tables["password_resets"] = password_resetsTableInfo
	tables["payment_status"] = payment_statusTableInfo
	tables["payment_status_history"] = payment_status_historyTableInfo
	tables["people"] = peopleTableInfo
	tables["people_taxonomy"] = people_taxonomyTableInfo
	tables["person_actions"] = person_actionsTableInfo
	tables["person_history"] = person_historyTableInfo
	tables["prices"] = pricesTableInfo
	tables["product_images"] = product_imagesTableInfo
	tables["product_taxonomy"] = product_taxonomyTableInfo
	tables["product_types"] = product_typesTableInfo
	tables["products"] = productsTableInfo
	tables["progress"] = progressTableInfo
	tables["settings"] = settingsTableInfo
	tables["shipping_status"] = shipping_statusTableInfo
	tables["shipping_status_history"] = shipping_status_historyTableInfo
	tables["tasks"] = tasksTableInfo
	tables["term_types"] = term_typesTableInfo
	tables["terms"] = termsTableInfo
	tables["users"] = usersTableInfo
	tables["users_in_apps"] = users_in_appsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
