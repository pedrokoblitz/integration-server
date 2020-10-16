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


CREATE TABLE `order_items` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` int(10) unsigned NOT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `price_id` int(10) unsigned NOT NULL,
  `quantity` int(11) NOT NULL DEFAULT '1',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=199 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "order_id": 33,    "product_id": 23,    "price_id": 79,    "quantity": 54,    "created_at": "2196-01-26T08:58:19.411206837-03:00",    "updated_at": "2102-05-15T02:15:55.305143596-03:00",    "id": 20}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// OrderItems struct is a row record of the order_items table in the laravel database
type OrderItems struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] order_id                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	OrderID uint32 `gorm:"column:order_id;type:uint;" json:"order_id"`
	//[ 2] product_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ProductID uint32 `gorm:"column:product_id;type:uint;" json:"product_id"`
	//[ 3] price_id                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	PriceID uint32 `gorm:"column:price_id;type:uint;" json:"price_id"`
	//[ 4] quantity                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Quantity int32 `gorm:"column:quantity;type:int;default:1;" json:"quantity"`
	//[ 5] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 6] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var order_itemsTableInfo = &TableInfo{
	Name: "order_items",
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
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "OrderID",
			GoFieldType:        "uint32",
			JSONFieldName:      "order_id",
			ProtobufFieldName:  "order_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "product_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "uint32",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "price_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PriceID",
			GoFieldType:        "uint32",
			JSONFieldName:      "price_id",
			ProtobufFieldName:  "price_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "quantity",
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
			GoFieldName:        "Quantity",
			GoFieldType:        "int32",
			JSONFieldName:      "quantity",
			ProtobufFieldName:  "quantity",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrderItems) TableName() string {
	return "order_items"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderItems) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderItems) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderItems) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrderItems) TableInfo() *TableInfo {
	return order_itemsTableInfo
}
