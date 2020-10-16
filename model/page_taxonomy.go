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


CREATE TABLE `page_taxonomy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `term_id` int(11) NOT NULL,
  `page_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "id": 37,    "term_id": 22,    "page_id": 82,    "created_at": "2236-10-16T03:06:35.026509056-03:00",    "updated_at": "2187-03-04T09:10:22.079736707-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// PageTaxonomy struct is a row record of the page_taxonomy table in the laravel database
type PageTaxonomy struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] term_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TermID int32 `gorm:"column:term_id;type:int;" json:"term_id"`
	//[ 2] page_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PageID int32 `gorm:"column:page_id;type:int;" json:"page_id"`
	//[ 3] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 4] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var page_taxonomyTableInfo = &TableInfo{
	Name: "page_taxonomy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "term_id",
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
			GoFieldName:        "TermID",
			GoFieldType:        "int32",
			JSONFieldName:      "term_id",
			ProtobufFieldName:  "term_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "page_id",
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
			GoFieldName:        "PageID",
			GoFieldType:        "int32",
			JSONFieldName:      "page_id",
			ProtobufFieldName:  "page_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PageTaxonomy) TableName() string {
	return "page_taxonomy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PageTaxonomy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PageTaxonomy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PageTaxonomy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PageTaxonomy) TableInfo() *TableInfo {
	return page_taxonomyTableInfo
}
