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


CREATE TABLE `larametrics_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `model` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `model_id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `method` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `original` text COLLATE utf8mb4_unicode_ci,
  `changes` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "created_at": "2121-01-20T22:29:30.279270535-03:00",    "id": 52,    "model": "jVghTqwXActYcuJwJCwQdNGHl",    "model_id": 77,    "user_id": 83,    "method": "RhAQiLwkmopUDyWamkwBHOUls",    "original": "sHiwPwTKAGRvCkBmaCbxnVcnN",    "changes": "mgvlyhSGcPoOoZXBWTicwCQjd",    "updated_at": "2073-10-29T10:52:27.320111816-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LarametricsModels struct is a row record of the larametrics_models table in the laravel database
type LarametricsModels struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] model                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Model string `gorm:"column:model;type:varchar;size:255;" json:"model"`
	//[ 2] model_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ModelID int32 `gorm:"column:model_id;type:int;" json:"model_id"`
	//[ 3] user_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UserID null.Int `gorm:"column:user_id;type:int;" json:"user_id"`
	//[ 4] method                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Method string `gorm:"column:method;type:varchar;size:255;" json:"method"`
	//[ 5] original                                       text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Original null.String `gorm:"column:original;type:text;size:65535;" json:"original"`
	//[ 6] changes                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Changes null.String `gorm:"column:changes;type:text;size:65535;" json:"changes"`
	//[ 7] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 8] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var larametrics_modelsTableInfo = &TableInfo{
	Name: "larametrics_models",
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
			Name:               "model",
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
			GoFieldName:        "Model",
			GoFieldType:        "string",
			JSONFieldName:      "model",
			ProtobufFieldName:  "model",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "model_id",
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
			GoFieldName:        "ModelID",
			GoFieldType:        "int32",
			JSONFieldName:      "model_id",
			ProtobufFieldName:  "model_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "user_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "null.Int",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "method",
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
			GoFieldName:        "Method",
			GoFieldType:        "string",
			JSONFieldName:      "method",
			ProtobufFieldName:  "method",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "original",
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
			GoFieldName:        "Original",
			GoFieldType:        "null.String",
			JSONFieldName:      "original",
			ProtobufFieldName:  "original",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "changes",
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
			GoFieldName:        "Changes",
			GoFieldType:        "null.String",
			JSONFieldName:      "changes",
			ProtobufFieldName:  "changes",
			ProtobufType:       "string",
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
func (l *LarametricsModels) TableName() string {
	return "larametrics_models"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LarametricsModels) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LarametricsModels) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LarametricsModels) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LarametricsModels) TableInfo() *TableInfo {
	return larametrics_modelsTableInfo
}
