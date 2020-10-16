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


CREATE TABLE `page_types` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL,
  `is_email` tinyint(4) NOT NULL DEFAULT '0',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "updated_at": "2143-12-12T03:15:16.650039767-03:00",    "id": 22,    "app_id": 11,    "is_email": 0,    "name": "LbUFKJgeEmddrJbSfPJUgSnTr",    "slug": "DqulhMGBCHGVCcOJDVmZMQiqw",    "created_at": "2065-11-02T05:49:33.888784383-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// PageTypes struct is a row record of the page_types table in the laravel database
type PageTypes struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] app_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AppID uint32 `gorm:"column:app_id;type:uint;" json:"app_id"`
	//[ 2] is_email                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsEmail int32 `gorm:"column:is_email;type:tinyint;default:0;" json:"is_email"`
	//[ 3] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[ 4] slug                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Slug string `gorm:"column:slug;type:varchar;size:255;" json:"slug"`
	//[ 5] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 6] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var page_typesTableInfo = &TableInfo{
	Name: "page_types",
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
			Name:               "is_email",
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
			GoFieldName:        "IsEmail",
			GoFieldType:        "int32",
			JSONFieldName:      "is_email",
			ProtobufFieldName:  "is_email",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "slug",
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
			GoFieldName:        "Slug",
			GoFieldType:        "string",
			JSONFieldName:      "slug",
			ProtobufFieldName:  "slug",
			ProtobufType:       "string",
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
func (p *PageTypes) TableName() string {
	return "page_types"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PageTypes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PageTypes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PageTypes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PageTypes) TableInfo() *TableInfo {
	return page_typesTableInfo
}
