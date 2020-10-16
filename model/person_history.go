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


CREATE TABLE `person_history` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `action_id` int(10) unsigned NOT NULL,
  `app_id` int(10) unsigned NOT NULL,
  `person_id` int(10) unsigned NOT NULL DEFAULT '0',
  `campaign_id` int(10) unsigned NOT NULL DEFAULT '0',
  `user_identifier` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "app_id": 27,    "id": 60,    "person_id": 47,    "campaign_id": 85,    "user_identifier": "oHdIvaemvNMOIwamSRVwnNvQG",    "deleted_at": "2286-08-17T22:48:31.582332-03:00",    "created_at": "2200-04-25T12:42:07.68413435-03:00",    "updated_at": "2223-06-20T11:25:58.860682274-03:00",    "action_id": 15}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// PersonHistory struct is a row record of the person_history table in the laravel database
type PersonHistory struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] action_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ActionID uint32 `gorm:"column:action_id;type:uint;" json:"action_id"`
	//[ 2] app_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AppID uint32 `gorm:"column:app_id;type:uint;" json:"app_id"`
	//[ 3] person_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PersonID uint32 `gorm:"column:person_id;type:uint;default:0;" json:"person_id"`
	//[ 4] campaign_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CampaignID uint32 `gorm:"column:campaign_id;type:uint;default:0;" json:"campaign_id"`
	//[ 5] user_identifier                                varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	UserIdentifier null.String `gorm:"column:user_identifier;type:varchar;size:32;" json:"user_identifier"`
	//[ 6] deleted_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	DeletedAt null.Time `gorm:"column:deleted_at;type:timestamp;" json:"deleted_at"`
	//[ 7] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 8] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var person_historyTableInfo = &TableInfo{
	Name: "person_history",
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
			Name:               "action_id",
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
			GoFieldName:        "ActionID",
			GoFieldType:        "uint32",
			JSONFieldName:      "action_id",
			ProtobufFieldName:  "action_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "person_id",
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
			GoFieldName:        "PersonID",
			GoFieldType:        "uint32",
			JSONFieldName:      "person_id",
			ProtobufFieldName:  "person_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "campaign_id",
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
			GoFieldName:        "CampaignID",
			GoFieldType:        "uint32",
			JSONFieldName:      "campaign_id",
			ProtobufFieldName:  "campaign_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "user_identifier",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "UserIdentifier",
			GoFieldType:        "null.String",
			JSONFieldName:      "user_identifier",
			ProtobufFieldName:  "user_identifier",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
func (p *PersonHistory) TableName() string {
	return "person_history"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PersonHistory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PersonHistory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PersonHistory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PersonHistory) TableInfo() *TableInfo {
	return person_historyTableInfo
}
