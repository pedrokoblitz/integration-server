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


CREATE TABLE `blocks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(11) NOT NULL,
  `type_id` int(10) unsigned NOT NULL,
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0',
  `page_id` int(10) unsigned NOT NULL DEFAULT '0',
  `order` int(11) NOT NULL DEFAULT '0',
  `limit` int(11) NOT NULL DEFAULT '4',
  `grid_size` int(11) NOT NULL DEFAULT '12',
  `grid_size_mobile` int(11) NOT NULL DEFAULT '12',
  `view` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `css` text COLLATE utf8mb4_unicode_ci,
  `item_id` int(10) unsigned NOT NULL DEFAULT '0',
  `term_id` int(10) unsigned NOT NULL DEFAULT '0',
  `line1` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `line2` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `line3` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `text1` text COLLATE utf8mb4_unicode_ci,
  `text2` text COLLATE utf8mb4_unicode_ci,
  `date1` datetime DEFAULT NULL,
  `date2` datetime DEFAULT NULL,
  `file1` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `file2` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `embed` text COLLATE utf8mb4_unicode_ci,
  `button_text` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "grid_size": 31,    "grid_size_mobile": 86,    "updated_at": "2115-02-22T23:52:52.065119948-03:00",    "order": 11,    "css": "SwoNDMdXDoJWXuLRQeqlmVIQT",    "text_1": "hPCOMdxuPMFxjypwoyYjwQFRP",    "date_2": "2041-11-08T14:24:52.381231424-03:00",    "embed": "fYGdyKYswNwNsoVNDvHwJqYGF",    "button_text": "rebkmBEsdRlcRGXDPVLMtCOqZ",    "limit": 85,    "line_2": "vDyFEKtdaWQHwfPHfoDjvbGuH",    "line_3": "jOecebYRJVvdVnPyLtWmOIOJS",    "url": "VrQsasIpjBvPfZNJCJmORBEgL",    "created_at": "2269-01-08T19:55:55.535369771-03:00",    "parent_id": 79,    "view": "QfFqTbLhfcGNTMvnqdyUQDNMN",    "term_id": 46,    "type_id": 30,    "item_id": 7,    "file_1": "DWmJIUAtaATCkghLlyuVxAWpv",    "line_1": "rAZvTdBdMQmEfLdhWuTNsYfxb",    "file_2": "ysDxhoclPDhwUsiMZucHaZRLe",    "date_1": "2065-09-03T22:17:17.769750058-03:00",    "text_2": "BQOxZOfAebejfYUOfkWelFFPA",    "id": 7,    "app_id": 70,    "page_id": 91}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned



*/

// Blocks struct is a row record of the blocks table in the laravel database
type Blocks struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] app_id                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AppID int32 `gorm:"column:app_id;type:int;" json:"app_id"`
	//[ 2] type_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	TypeID uint32 `gorm:"column:type_id;type:uint;" json:"type_id"`
	//[ 3] parent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ParentID uint32 `gorm:"column:parent_id;type:uint;default:0;" json:"parent_id"`
	//[ 4] page_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PageID uint32 `gorm:"column:page_id;type:uint;default:0;" json:"page_id"`
	//[ 5] order                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Order int32 `gorm:"column:order;type:int;default:0;" json:"order"`
	//[ 6] limit                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [4]
	Limit int32 `gorm:"column:limit;type:int;default:4;" json:"limit"`
	//[ 7] grid_size                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [12]
	GridSize int32 `gorm:"column:grid_size;type:int;default:12;" json:"grid_size"`
	//[ 8] grid_size_mobile                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [12]
	GridSizeMobile int32 `gorm:"column:grid_size_mobile;type:int;default:12;" json:"grid_size_mobile"`
	//[ 9] view                                           varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	View null.String `gorm:"column:view;type:varchar;size:255;" json:"view"`
	//[10] css                                            text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	CSS null.String `gorm:"column:css;type:text;size:65535;" json:"css"`
	//[11] item_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemID uint32 `gorm:"column:item_id;type:uint;default:0;" json:"item_id"`
	//[12] term_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TermID uint32 `gorm:"column:term_id;type:uint;default:0;" json:"term_id"`
	//[13] line1                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Line1 null.String `gorm:"column:line1;type:varchar;size:255;" json:"line_1"`
	//[14] line2                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Line2 null.String `gorm:"column:line2;type:varchar;size:255;" json:"line_2"`
	//[15] line3                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Line3 null.String `gorm:"column:line3;type:varchar;size:255;" json:"line_3"`
	//[16] text1                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Text1 null.String `gorm:"column:text1;type:text;size:65535;" json:"text_1"`
	//[17] text2                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Text2 null.String `gorm:"column:text2;type:text;size:65535;" json:"text_2"`
	//[18] date1                                          datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Date1 null.Time `gorm:"column:date1;type:datetime;" json:"date_1"`
	//[19] date2                                          datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Date2 null.Time `gorm:"column:date2;type:datetime;" json:"date_2"`
	//[20] file1                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	File1 null.String `gorm:"column:file1;type:varchar;size:255;" json:"file_1"`
	//[21] file2                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	File2 null.String `gorm:"column:file2;type:varchar;size:255;" json:"file_2"`
	//[22] embed                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Embed null.String `gorm:"column:embed;type:text;size:65535;" json:"embed"`
	//[23] button_text                                    varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ButtonText null.String `gorm:"column:button_text;type:varchar;size:255;" json:"button_text"`
	//[24] url                                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	URL null.String `gorm:"column:url;type:varchar;size:255;" json:"url"`
	//[25] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[26] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var blocksTableInfo = &TableInfo{
	Name: "blocks",
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "page_id",
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
			GoFieldName:        "PageID",
			GoFieldType:        "uint32",
			JSONFieldName:      "page_id",
			ProtobufFieldName:  "page_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "order",
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
			GoFieldName:        "Order",
			GoFieldType:        "int32",
			JSONFieldName:      "order",
			ProtobufFieldName:  "order",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "limit",
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
			GoFieldName:        "Limit",
			GoFieldType:        "int32",
			JSONFieldName:      "limit",
			ProtobufFieldName:  "limit",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "grid_size",
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
			GoFieldName:        "GridSize",
			GoFieldType:        "int32",
			JSONFieldName:      "grid_size",
			ProtobufFieldName:  "grid_size",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "grid_size_mobile",
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
			GoFieldName:        "GridSizeMobile",
			GoFieldType:        "int32",
			JSONFieldName:      "grid_size_mobile",
			ProtobufFieldName:  "grid_size_mobile",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "view",
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
			GoFieldName:        "View",
			GoFieldType:        "null.String",
			JSONFieldName:      "view",
			ProtobufFieldName:  "view",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "css",
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
			GoFieldName:        "CSS",
			GoFieldType:        "null.String",
			JSONFieldName:      "css",
			ProtobufFieldName:  "css",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "term_id",
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
			GoFieldName:        "TermID",
			GoFieldType:        "uint32",
			JSONFieldName:      "term_id",
			ProtobufFieldName:  "term_id",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "line1",
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
			GoFieldName:        "Line1",
			GoFieldType:        "null.String",
			JSONFieldName:      "line_1",
			ProtobufFieldName:  "line_1",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "line2",
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
			GoFieldName:        "Line2",
			GoFieldType:        "null.String",
			JSONFieldName:      "line_2",
			ProtobufFieldName:  "line_2",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "line3",
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
			GoFieldName:        "Line3",
			GoFieldType:        "null.String",
			JSONFieldName:      "line_3",
			ProtobufFieldName:  "line_3",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "text1",
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
			GoFieldName:        "Text1",
			GoFieldType:        "null.String",
			JSONFieldName:      "text_1",
			ProtobufFieldName:  "text_1",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "text2",
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
			GoFieldName:        "Text2",
			GoFieldType:        "null.String",
			JSONFieldName:      "text_2",
			ProtobufFieldName:  "text_2",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "date1",
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
			GoFieldName:        "Date1",
			GoFieldType:        "null.Time",
			JSONFieldName:      "date_1",
			ProtobufFieldName:  "date_1",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "date2",
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
			GoFieldName:        "Date2",
			GoFieldType:        "null.Time",
			JSONFieldName:      "date_2",
			ProtobufFieldName:  "date_2",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "file1",
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
			GoFieldName:        "File1",
			GoFieldType:        "null.String",
			JSONFieldName:      "file_1",
			ProtobufFieldName:  "file_1",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "file2",
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
			GoFieldName:        "File2",
			GoFieldType:        "null.String",
			JSONFieldName:      "file_2",
			ProtobufFieldName:  "file_2",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "embed",
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
			GoFieldName:        "Embed",
			GoFieldType:        "null.String",
			JSONFieldName:      "embed",
			ProtobufFieldName:  "embed",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "button_text",
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
			GoFieldName:        "ButtonText",
			GoFieldType:        "null.String",
			JSONFieldName:      "button_text",
			ProtobufFieldName:  "button_text",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "url",
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
			GoFieldName:        "URL",
			GoFieldType:        "null.String",
			JSONFieldName:      "url",
			ProtobufFieldName:  "url",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
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
			ProtobufPos:        27,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *Blocks) TableName() string {
	return "blocks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *Blocks) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *Blocks) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *Blocks) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *Blocks) TableInfo() *TableInfo {
	return blocksTableInfo
}
