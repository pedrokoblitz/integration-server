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


CREATE TABLE `prices` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `product_id` int(10) unsigned NOT NULL,
  `price` int(11) NOT NULL DEFAULT '1',
  `activity` tinyint(4) NOT NULL DEFAULT '1',
  `default` tinyint(4) NOT NULL DEFAULT '1',
  `fixed_discount` int(11) NOT NULL DEFAULT '0',
  `percent_discount` int(11) NOT NULL DEFAULT '0',
  `free_shipping` tinyint(4) NOT NULL DEFAULT '0',
  `begins` datetime DEFAULT NULL,
  `expires` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "product_id": 73,    "activity": 97,    "begins": "2033-05-30T14:42:47.846037436-03:00",    "expires": "2094-01-02T13:55:53.940334569-03:00",    "created_at": "2214-04-09T02:43:46.195708775-03:00",    "id": 83,    "price": 3,    "default": 73,    "fixed_discount": 57,    "percent_discount": 54,    "free_shipping": 98,    "updated_at": "2154-04-15T12:36:16.647033164-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Prices struct is a row record of the prices table in the laravel database
type Prices struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] product_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ProductID uint32 `gorm:"column:product_id;type:uint;" json:"product_id"`
	//[ 2] price                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Price int32 `gorm:"column:price;type:int;default:1;" json:"price"`
	//[ 3] activity                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Activity int32 `gorm:"column:activity;type:tinyint;default:1;" json:"activity"`
	//[ 4] default                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Default int32 `gorm:"column:default;type:tinyint;default:1;" json:"default"`
	//[ 5] fixed_discount                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FixedDiscount int32 `gorm:"column:fixed_discount;type:int;default:0;" json:"fixed_discount"`
	//[ 6] percent_discount                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PercentDiscount int32 `gorm:"column:percent_discount;type:int;default:0;" json:"percent_discount"`
	//[ 7] free_shipping                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	FreeShipping int32 `gorm:"column:free_shipping;type:tinyint;default:0;" json:"free_shipping"`
	//[ 8] begins                                         datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Begins null.Time `gorm:"column:begins;type:datetime;" json:"begins"`
	//[ 9] expires                                        datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Expires null.Time `gorm:"column:expires;type:datetime;" json:"expires"`
	//[10] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[11] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var pricesTableInfo = &TableInfo{
	Name: "prices",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "price",
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
			GoFieldName:        "Price",
			GoFieldType:        "int32",
			JSONFieldName:      "price",
			ProtobufFieldName:  "price",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "activity",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Activity",
			GoFieldType:        "int32",
			JSONFieldName:      "activity",
			ProtobufFieldName:  "activity",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "default",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Default",
			GoFieldType:        "int32",
			JSONFieldName:      "default",
			ProtobufFieldName:  "default",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "fixed_discount",
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
			GoFieldName:        "FixedDiscount",
			GoFieldType:        "int32",
			JSONFieldName:      "fixed_discount",
			ProtobufFieldName:  "fixed_discount",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "percent_discount",
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
			GoFieldName:        "PercentDiscount",
			GoFieldType:        "int32",
			JSONFieldName:      "percent_discount",
			ProtobufFieldName:  "percent_discount",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "free_shipping",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "FreeShipping",
			GoFieldType:        "int32",
			JSONFieldName:      "free_shipping",
			ProtobufFieldName:  "free_shipping",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "begins",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "Begins",
			GoFieldType:        "null.Time",
			JSONFieldName:      "begins",
			ProtobufFieldName:  "begins",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "expires",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "Expires",
			GoFieldType:        "null.Time",
			JSONFieldName:      "expires",
			ProtobufFieldName:  "expires",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Prices) TableName() string {
	return "prices"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Prices) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Prices) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Prices) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Prices) TableInfo() *TableInfo {
	return pricesTableInfo
}
