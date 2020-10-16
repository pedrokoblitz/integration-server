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


CREATE TABLE `coupons` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL,
  `product_id` int(10) unsigned NOT NULL DEFAULT '0',
  `fixed_discount` int(11) NOT NULL DEFAULT '0',
  `percent_discount` int(11) NOT NULL DEFAULT '0',
  `free_shipping` tinyint(4) NOT NULL DEFAULT '0',
  `ref_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `begins` datetime DEFAULT NULL,
  `expires` datetime DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "created_at": "2021-02-18T17:32:43.030293573-03:00",    "app_id": 38,    "fixed_discount": 51,    "free_shipping": 87,    "begins": "2070-05-23T07:59:03.692302905-03:00",    "deleted_at": "2070-02-06T22:52:31.216692114-03:00",    "updated_at": "2031-06-10T19:59:04.078202256-03:00",    "percent_discount": 14,    "ref_code": "RnxwTphmCyGSsCqScbknhkbYZ",    "name": "iUnyXAlaWSsOnkJPKuhHDMSOq",    "id": 53,    "product_id": 64,    "description": "kKEgKRIDvYfwHdUGYxsZGygDH",    "expires": "2234-07-26T03:03:23.289025469-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// Coupons struct is a row record of the coupons table in the laravel database
type Coupons struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] app_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AppID uint32 `gorm:"column:app_id;type:uint;" json:"app_id"`
	//[ 2] product_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ProductID uint32 `gorm:"column:product_id;type:uint;default:0;" json:"product_id"`
	//[ 3] fixed_discount                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FixedDiscount int32 `gorm:"column:fixed_discount;type:int;default:0;" json:"fixed_discount"`
	//[ 4] percent_discount                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PercentDiscount int32 `gorm:"column:percent_discount;type:int;default:0;" json:"percent_discount"`
	//[ 5] free_shipping                                  tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	FreeShipping int32 `gorm:"column:free_shipping;type:tinyint;default:0;" json:"free_shipping"`
	//[ 6] ref_code                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	RefCode string `gorm:"column:ref_code;type:varchar;size:255;" json:"ref_code"`
	//[ 7] name                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name null.String `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[ 8] description                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Description null.String `gorm:"column:description;type:varchar;size:255;" json:"description"`
	//[ 9] begins                                         datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Begins null.Time `gorm:"column:begins;type:datetime;" json:"begins"`
	//[10] expires                                        datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Expires null.Time `gorm:"column:expires;type:datetime;" json:"expires"`
	//[11] deleted_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	DeletedAt null.Time `gorm:"column:deleted_at;type:timestamp;" json:"deleted_at"`
	//[12] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[13] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var couponsTableInfo = &TableInfo{
	Name: "coupons",
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
			Name:               "app_id",
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
			GoFieldName:        "AppID",
			GoFieldType:        "uint32",
			JSONFieldName:      "app_id",
			ProtobufFieldName:  "app_id",
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "ref_code",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "RefCode",
			GoFieldType:        "string",
			JSONFieldName:      "ref_code",
			ProtobufFieldName:  "ref_code",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Description",
			GoFieldType:        "null.String",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "deleted_at",
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
			GoFieldName:        "DeletedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "deleted_at",
			ProtobufFieldName:  "deleted_at",
			ProtobufType:       "uint64",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Coupons) TableName() string {
	return "coupons"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Coupons) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Coupons) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Coupons) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Coupons) TableInfo() *TableInfo {
	return couponsTableInfo
}
