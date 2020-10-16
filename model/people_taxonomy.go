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


CREATE TABLE `people_taxonomy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `person_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "updated_at": "2180-10-15T03:44:25.773426007-03:00",    "id": 1,    "person_id": 40,    "product_id": 30,    "created_at": "2170-10-28T21:08:57.007968305-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// PeopleTaxonomy struct is a row record of the people_taxonomy table in the laravel database
type PeopleTaxonomy struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] person_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PersonID int32 `gorm:"column:person_id;type:int;" json:"person_id"`
	//[ 2] product_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ProductID int32 `gorm:"column:product_id;type:int;" json:"product_id"`
	//[ 3] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 4] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var people_taxonomyTableInfo = &TableInfo{
	Name: "people_taxonomy",
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
			Name:               "person_id",
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
			GoFieldName:        "PersonID",
			GoFieldType:        "int32",
			JSONFieldName:      "person_id",
			ProtobufFieldName:  "person_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "product_id",
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
			GoFieldName:        "ProductID",
			GoFieldType:        "int32",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
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
func (p *PeopleTaxonomy) TableName() string {
	return "people_taxonomy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PeopleTaxonomy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PeopleTaxonomy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PeopleTaxonomy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PeopleTaxonomy) TableInfo() *TableInfo {
	return people_taxonomyTableInfo
}
