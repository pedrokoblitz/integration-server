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


CREATE TABLE `pages` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `type_id` int(10) unsigned NOT NULL,
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0',
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_image` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta_keywords` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `css` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "user_id": 25,    "slug": "rHBwMBNGsJFKsYivhjtlIUJuf",    "meta_image": "kxNOkadKxiOmaNbncSwDSwiaa",    "app_id": 56,    "updated_at": "2114-07-29T20:27:34.547714751-03:00",    "parent_id": 27,    "type_id": 83,    "meta_description": "DtFIcOCYjQhmkKLwODIdkgUTv",    "status": "cgJCqyYXXjKGDHXfxbcSFBlrZ",    "css": "cDvjjeGcZRfvLJiHlQWwZvpUx",    "id": 89,    "meta_keywords": "WsdyIbSPleTJGXgSvYdQJuRwx",    "created_at": "2094-06-08T21:37:47.171627177-03:00",    "meta_title": "VnfQniIjXGZuhMnudjUhQYxMj"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// Pages struct is a row record of the pages table in the laravel database
type Pages struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] app_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AppID uint32 `gorm:"column:app_id;type:uint;" json:"app_id"`
	//[ 2] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	UserID uint32 `gorm:"column:user_id;type:uint;" json:"user_id"`
	//[ 3] type_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	TypeID uint32 `gorm:"column:type_id;type:uint;" json:"type_id"`
	//[ 4] parent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ParentID uint32 `gorm:"column:parent_id;type:uint;default:0;" json:"parent_id"`
	//[ 5] slug                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Slug null.String `gorm:"column:slug;type:varchar;size:255;" json:"slug"`
	//[ 6] meta_image                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	MetaImage null.String `gorm:"column:meta_image;type:varchar;size:255;" json:"meta_image"`
	//[ 7] meta_title                                     varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	MetaTitle null.String `gorm:"column:meta_title;type:varchar;size:255;" json:"meta_title"`
	//[ 8] meta_description                               varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	MetaDescription null.String `gorm:"column:meta_description;type:varchar;size:255;" json:"meta_description"`
	//[ 9] meta_keywords                                  varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	MetaKeywords null.String `gorm:"column:meta_keywords;type:varchar;size:255;" json:"meta_keywords"`
	//[10] status                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Status null.String `gorm:"column:status;type:varchar;size:255;" json:"status"`
	//[11] css                                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CSS null.String `gorm:"column:css;type:varchar;size:255;" json:"css"`
	//[12] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[13] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var pagesTableInfo = &TableInfo{
	Name: "pages",
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
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "uint32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "parent_id",
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
			GoFieldName:        "ParentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "parent_id",
			ProtobufFieldName:  "parent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "slug",
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
			GoFieldName:        "Slug",
			GoFieldType:        "null.String",
			JSONFieldName:      "slug",
			ProtobufFieldName:  "slug",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "meta_image",
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
			GoFieldName:        "MetaImage",
			GoFieldType:        "null.String",
			JSONFieldName:      "meta_image",
			ProtobufFieldName:  "meta_image",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "meta_title",
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
			GoFieldName:        "MetaTitle",
			GoFieldType:        "null.String",
			JSONFieldName:      "meta_title",
			ProtobufFieldName:  "meta_title",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "meta_description",
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
			GoFieldName:        "MetaDescription",
			GoFieldType:        "null.String",
			JSONFieldName:      "meta_description",
			ProtobufFieldName:  "meta_description",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "meta_keywords",
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
			GoFieldName:        "MetaKeywords",
			GoFieldType:        "null.String",
			JSONFieldName:      "meta_keywords",
			ProtobufFieldName:  "meta_keywords",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "status",
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
			GoFieldName:        "Status",
			GoFieldType:        "null.String",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "css",
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
			GoFieldName:        "CSS",
			GoFieldType:        "null.String",
			JSONFieldName:      "css",
			ProtobufFieldName:  "css",
			ProtobufType:       "string",
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
func (p *Pages) TableName() string {
	return "pages"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Pages) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Pages) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Pages) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Pages) TableInfo() *TableInfo {
	return pagesTableInfo
}
