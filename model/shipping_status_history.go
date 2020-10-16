package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `shipping_status_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(11) NOT NULL,
  `old_status_id` int(11) NOT NULL DEFAULT '0',
  `new_status_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=298 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "new_status_id": 14,    "created_at": "2304-07-13T21:57:01.651124065-03:00",    "updated_at": "2126-10-25T01:59:56.280583738-03:00",    "id": 6,    "order_id": 97,    "old_status_id": 8}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// ShippingStatusHistory struct is a row record of the shipping_status_history table in the laravel database
type ShippingStatusHistory struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] order_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	OrderID int32 `gorm:"column:order_id;type:int;" json:"order_id"`
	//[ 2] old_status_id                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OldStatusID int32 `gorm:"column:old_status_id;type:int;default:0;" json:"old_status_id"`
	//[ 3] new_status_id                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	NewStatusID int32 `gorm:"column:new_status_id;type:int;" json:"new_status_id"`
	//[ 4] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 5] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var shipping_status_historyTableInfo = &TableInfo{
	Name: "shipping_status_history",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "order_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OrderID",
			GoFieldType:        "int32",
			JSONFieldName:      "order_id",
			ProtobufFieldName:  "order_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "old_status_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OldStatusID",
			GoFieldType:        "int32",
			JSONFieldName:      "old_status_id",
			ProtobufFieldName:  "old_status_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "new_status_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "NewStatusID",
			GoFieldType:        "int32",
			JSONFieldName:      "new_status_id",
			ProtobufFieldName:  "new_status_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShippingStatusHistory) TableName() string {
	return "shipping_status_history"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShippingStatusHistory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShippingStatusHistory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShippingStatusHistory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShippingStatusHistory) TableInfo() *TableInfo {
	return shipping_status_historyTableInfo
}
