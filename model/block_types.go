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


CREATE TABLE `block_types` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_id` int(11) NOT NULL,
  `is_module` tinyint(4) NOT NULL DEFAULT '0',
  `is_page_term` tinyint(4) NOT NULL DEFAULT '0',
  `is_product_term` tinyint(4) NOT NULL DEFAULT '0',
  `is_page` tinyint(4) NOT NULL DEFAULT '0',
  `is_product` tinyint(4) NOT NULL DEFAULT '0',
  `is_parent` tinyint(4) NOT NULL DEFAULT '0',
  `is_child` tinyint(4) NOT NULL DEFAULT '0',
  `line1` tinyint(4) NOT NULL DEFAULT '1',
  `line2` tinyint(4) NOT NULL DEFAULT '0',
  `line3` tinyint(4) NOT NULL DEFAULT '0',
  `text1` tinyint(4) NOT NULL DEFAULT '0',
  `text2` tinyint(4) NOT NULL DEFAULT '0',
  `file1` tinyint(4) NOT NULL DEFAULT '0',
  `file2` tinyint(4) NOT NULL DEFAULT '0',
  `embed` tinyint(4) NOT NULL DEFAULT '0',
  `button_text` tinyint(4) NOT NULL DEFAULT '0',
  `url` tinyint(4) NOT NULL DEFAULT '0',
  `date1` tinyint(4) NOT NULL DEFAULT '0',
  `date2` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "is_page_term": 80,    "file_1": 1,    "updated_at": "2119-10-19T08:24:00.525421887-03:00",    "id": 29,    "app_id": 32,    "is_parent": 85,    "line_3": 55,    "button_text": 5,    "date_1": 40,    "slug": "vpyfGCMjgDufiSIgyKIdRvuQV",    "is_page": 85,    "line_2": 84,    "text_1": 33,    "text_2": 42,    "embed": 77,    "created_at": "2069-09-05T00:04:59.555708001-03:00",    "name": "feakKKVlkMYixQSDTNNnfjrrW",    "is_module": 28,    "is_child": 29,    "line_1": 57,    "file_2": 69,    "url": 64,    "date_2": 83,    "is_product_term": 24,    "is_product": 9}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// BlockTypes struct is a row record of the block_types table in the laravel database
type BlockTypes struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[ 2] slug                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Slug string `gorm:"column:slug;type:varchar;size:255;" json:"slug"`
	//[ 3] app_id                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AppID int32 `gorm:"column:app_id;type:int;" json:"app_id"`
	//[ 4] is_module                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsModule int32 `gorm:"column:is_module;type:tinyint;default:0;" json:"is_module"`
	//[ 5] is_page_term                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsPageTerm int32 `gorm:"column:is_page_term;type:tinyint;default:0;" json:"is_page_term"`
	//[ 6] is_product_term                                tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsProductTerm int32 `gorm:"column:is_product_term;type:tinyint;default:0;" json:"is_product_term"`
	//[ 7] is_page                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsPage int32 `gorm:"column:is_page;type:tinyint;default:0;" json:"is_page"`
	//[ 8] is_product                                     tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsProduct int32 `gorm:"column:is_product;type:tinyint;default:0;" json:"is_product"`
	//[ 9] is_parent                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsParent int32 `gorm:"column:is_parent;type:tinyint;default:0;" json:"is_parent"`
	//[10] is_child                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsChild int32 `gorm:"column:is_child;type:tinyint;default:0;" json:"is_child"`
	//[11] line1                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Line1 int32 `gorm:"column:line1;type:tinyint;default:1;" json:"line_1"`
	//[12] line2                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Line2 int32 `gorm:"column:line2;type:tinyint;default:0;" json:"line_2"`
	//[13] line3                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Line3 int32 `gorm:"column:line3;type:tinyint;default:0;" json:"line_3"`
	//[14] text1                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Text1 int32 `gorm:"column:text1;type:tinyint;default:0;" json:"text_1"`
	//[15] text2                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Text2 int32 `gorm:"column:text2;type:tinyint;default:0;" json:"text_2"`
	//[16] file1                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	File1 int32 `gorm:"column:file1;type:tinyint;default:0;" json:"file_1"`
	//[17] file2                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	File2 int32 `gorm:"column:file2;type:tinyint;default:0;" json:"file_2"`
	//[18] embed                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Embed int32 `gorm:"column:embed;type:tinyint;default:0;" json:"embed"`
	//[19] button_text                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	ButtonText int32 `gorm:"column:button_text;type:tinyint;default:0;" json:"button_text"`
	//[20] url                                            tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	URL int32 `gorm:"column:url;type:tinyint;default:0;" json:"url"`
	//[21] date1                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Date1 int32 `gorm:"column:date1;type:tinyint;default:0;" json:"date_1"`
	//[22] date2                                          tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Date2 int32 `gorm:"column:date2;type:tinyint;default:0;" json:"date_2"`
	//[23] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[24] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var block_typesTableInfo = &TableInfo{
	Name: "block_types",
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
			Name:               "app_id",
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
			GoFieldName:        "AppID",
			GoFieldType:        "int32",
			JSONFieldName:      "app_id",
			ProtobufFieldName:  "app_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "is_module",
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
			GoFieldName:        "IsModule",
			GoFieldType:        "int32",
			JSONFieldName:      "is_module",
			ProtobufFieldName:  "is_module",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "is_page_term",
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
			GoFieldName:        "IsPageTerm",
			GoFieldType:        "int32",
			JSONFieldName:      "is_page_term",
			ProtobufFieldName:  "is_page_term",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "is_product_term",
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
			GoFieldName:        "IsProductTerm",
			GoFieldType:        "int32",
			JSONFieldName:      "is_product_term",
			ProtobufFieldName:  "is_product_term",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "is_page",
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
			GoFieldName:        "IsPage",
			GoFieldType:        "int32",
			JSONFieldName:      "is_page",
			ProtobufFieldName:  "is_page",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "is_product",
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
			GoFieldName:        "IsProduct",
			GoFieldType:        "int32",
			JSONFieldName:      "is_product",
			ProtobufFieldName:  "is_product",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "is_parent",
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
			GoFieldName:        "IsParent",
			GoFieldType:        "int32",
			JSONFieldName:      "is_parent",
			ProtobufFieldName:  "is_parent",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "is_child",
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
			GoFieldName:        "IsChild",
			GoFieldType:        "int32",
			JSONFieldName:      "is_child",
			ProtobufFieldName:  "is_child",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "line1",
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
			GoFieldName:        "Line1",
			GoFieldType:        "int32",
			JSONFieldName:      "line_1",
			ProtobufFieldName:  "line_1",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "line2",
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
			GoFieldName:        "Line2",
			GoFieldType:        "int32",
			JSONFieldName:      "line_2",
			ProtobufFieldName:  "line_2",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "line3",
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
			GoFieldName:        "Line3",
			GoFieldType:        "int32",
			JSONFieldName:      "line_3",
			ProtobufFieldName:  "line_3",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "text1",
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
			GoFieldName:        "Text1",
			GoFieldType:        "int32",
			JSONFieldName:      "text_1",
			ProtobufFieldName:  "text_1",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "text2",
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
			GoFieldName:        "Text2",
			GoFieldType:        "int32",
			JSONFieldName:      "text_2",
			ProtobufFieldName:  "text_2",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "file1",
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
			GoFieldName:        "File1",
			GoFieldType:        "int32",
			JSONFieldName:      "file_1",
			ProtobufFieldName:  "file_1",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "file2",
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
			GoFieldName:        "File2",
			GoFieldType:        "int32",
			JSONFieldName:      "file_2",
			ProtobufFieldName:  "file_2",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "embed",
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
			GoFieldName:        "Embed",
			GoFieldType:        "int32",
			JSONFieldName:      "embed",
			ProtobufFieldName:  "embed",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "button_text",
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
			GoFieldName:        "ButtonText",
			GoFieldType:        "int32",
			JSONFieldName:      "button_text",
			ProtobufFieldName:  "button_text",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "url",
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
			GoFieldName:        "URL",
			GoFieldType:        "int32",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "date1",
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
			GoFieldName:        "Date1",
			GoFieldType:        "int32",
			JSONFieldName:      "date_1",
			ProtobufFieldName:  "date_1",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "date2",
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
			GoFieldName:        "Date2",
			GoFieldType:        "int32",
			JSONFieldName:      "date_2",
			ProtobufFieldName:  "date_2",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BlockTypes) TableName() string {
	return "block_types"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BlockTypes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BlockTypes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BlockTypes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BlockTypes) TableInfo() *TableInfo {
	return block_typesTableInfo
}
