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


CREATE TABLE `products` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL,
  `type_id` int(10) unsigned NOT NULL,
  `activity` tinyint(4) NOT NULL DEFAULT '1',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `stock_keeping` int(11) NOT NULL DEFAULT '1',
  `keywords` text COLLATE utf8mb4_unicode_ci,
  `meta_tag_description` text COLLATE utf8mb4_unicode_ci,
  `sku` bigint(20) NOT NULL DEFAULT '0',
  `cubicweight` int(11) NOT NULL DEFAULT '0',
  `weight` int(11) NOT NULL DEFAULT '0',
  `height` int(11) NOT NULL DEFAULT '0',
  `length` int(11) NOT NULL DEFAULT '0',
  `width` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "keywords": "iCFiFAfEIGncWXlSQoWBidmna",    "cubicweight": 69,    "height": 41,    "created_at": "2024-02-15T01:46:22.872287981-03:00",    "type_id": 67,    "name": "WiXPmKNiWlvENxnZrRljQoPpi",    "stock_keeping": 65,    "id": 73,    "meta_tag_description": "LuckAthEtyliFTypZYGrOnEbj",    "length": 95,    "updated_at": "2147-11-02T16:34:19.019091582-03:00",    "app_id": 86,    "sku": 54,    "weight": 70,    "width": 28,    "activity": 70}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned



*/

// Products struct is a row record of the products table in the laravel database
type Products struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] app_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AppID uint32 `gorm:"column:app_id;type:uint;" json:"app_id"`
	//[ 2] type_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	TypeID uint32 `gorm:"column:type_id;type:uint;" json:"type_id"`
	//[ 3] activity                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Activity int32 `gorm:"column:activity;type:tinyint;default:1;" json:"activity"`
	//[ 4] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[ 5] stock_keeping                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	StockKeeping int32 `gorm:"column:stock_keeping;type:int;default:1;" json:"stock_keeping"`
	//[ 6] keywords                                       text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Keywords null.String `gorm:"column:keywords;type:text;size:65535;" json:"keywords"`
	//[ 7] meta_tag_description                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	MetaTagDescription null.String `gorm:"column:meta_tag_description;type:text;size:65535;" json:"meta_tag_description"`
	//[ 8] sku                                            bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	Sku int64 `gorm:"column:sku;type:bigint;default:0;" json:"sku"`
	//[ 9] cubicweight                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Cubicweight int32 `gorm:"column:cubicweight;type:int;default:0;" json:"cubicweight"`
	//[10] weight                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Weight int32 `gorm:"column:weight;type:int;default:0;" json:"weight"`
	//[11] height                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Height int32 `gorm:"column:height;type:int;default:0;" json:"height"`
	//[12] length                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Length int32 `gorm:"column:length;type:int;default:0;" json:"length"`
	//[13] width                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Width int32 `gorm:"column:width;type:int;default:0;" json:"width"`
	//[14] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[15] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var productsTableInfo = &TableInfo{
	Name: "products",
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
			Name:               "type_id",
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
			GoFieldName:        "TypeID",
			GoFieldType:        "uint32",
			JSONFieldName:      "type_id",
			ProtobufFieldName:  "type_id",
			ProtobufType:       "uint32",
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
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "stock_keeping",
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
			GoFieldName:        "StockKeeping",
			GoFieldType:        "int32",
			JSONFieldName:      "stock_keeping",
			ProtobufFieldName:  "stock_keeping",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "keywords",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Keywords",
			GoFieldType:        "null.String",
			JSONFieldName:      "keywords",
			ProtobufFieldName:  "keywords",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "meta_tag_description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "MetaTagDescription",
			GoFieldType:        "null.String",
			JSONFieldName:      "meta_tag_description",
			ProtobufFieldName:  "meta_tag_description",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "sku",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "Sku",
			GoFieldType:        "int64",
			JSONFieldName:      "sku",
			ProtobufFieldName:  "sku",
			ProtobufType:       "int64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "cubicweight",
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
			GoFieldName:        "Cubicweight",
			GoFieldType:        "int32",
			JSONFieldName:      "cubicweight",
			ProtobufFieldName:  "cubicweight",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "weight",
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
			GoFieldName:        "Weight",
			GoFieldType:        "int32",
			JSONFieldName:      "weight",
			ProtobufFieldName:  "weight",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "height",
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
			GoFieldName:        "Height",
			GoFieldType:        "int32",
			JSONFieldName:      "height",
			ProtobufFieldName:  "height",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "length",
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
			GoFieldName:        "Length",
			GoFieldType:        "int32",
			JSONFieldName:      "length",
			ProtobufFieldName:  "length",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "width",
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
			GoFieldName:        "Width",
			GoFieldType:        "int32",
			JSONFieldName:      "width",
			ProtobufFieldName:  "width",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Products) TableName() string {
	return "products"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Products) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Products) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Products) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Products) TableInfo() *TableInfo {
	return productsTableInfo
}
