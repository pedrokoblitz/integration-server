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


CREATE TABLE `apps` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `logo` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_content` tinyint(4) NOT NULL DEFAULT '0',
  `is_store` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "name": "XJBqIHjlouayrJKyoAYMcCVCu",    "description": "lhoyXdRddLwCCShRTCjcaTKbt",    "is_content": 82,    "is_store": 87,    "id": 65,    "logo": "FsOoanSOJoGCEQuuFojoAdlVh",    "created_at": "2289-07-19T14:34:46.012738659-03:00",    "updated_at": "2101-11-06T21:15:32.21372178-03:00",    "slug": "RVYHDmrRWtrcGQDMmTWUfGjao"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Apps struct is a row record of the apps table in the laravel database
type Apps struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[ 2] slug                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Slug string `gorm:"column:slug;type:varchar;size:255;" json:"slug"`
	//[ 3] logo                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Logo null.String `gorm:"column:logo;type:varchar;size:255;" json:"logo"`
	//[ 4] description                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Description null.String `gorm:"column:description;type:varchar;size:255;" json:"description"`
	//[ 5] is_content                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsContent int32 `gorm:"column:is_content;type:tinyint;default:0;" json:"is_content"`
	//[ 6] is_store                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsStore int32 `gorm:"column:is_store;type:tinyint;default:0;" json:"is_store"`
	//[ 7] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 8] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var appsTableInfo = &TableInfo{
	Name: "apps",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "logo",
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
			GoFieldName:        "Logo",
			GoFieldType:        "null.String",
			JSONFieldName:      "logo",
			ProtobufFieldName:  "logo",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "is_content",
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
			GoFieldName:        "IsContent",
			GoFieldType:        "int32",
			JSONFieldName:      "is_content",
			ProtobufFieldName:  "is_content",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "is_store",
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
			GoFieldName:        "IsStore",
			GoFieldType:        "int32",
			JSONFieldName:      "is_store",
			ProtobufFieldName:  "is_store",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Apps) TableName() string {
	return "apps"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Apps) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Apps) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Apps) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Apps) TableInfo() *TableInfo {
	return appsTableInfo
}
