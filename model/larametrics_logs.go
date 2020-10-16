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


CREATE TABLE `larametrics_logs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `level` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `trace` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "level": "LxfYoZxtRrYInPpMpEfuyKhgq",    "message": "vKepwIcDdjnyoIBWFecTfvxhV",    "user_id": 90,    "email": "iDZQwJjjvPdEROXbkHDFogmRB",    "trace": "IMxgDuDqeZdWsEKpZQiCxZwXF",    "created_at": "2169-10-20T21:55:42.976175825-03:00",    "updated_at": "2068-11-08T13:17:58.340737413-03:00",    "id": 62}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LarametricsLogs struct is a row record of the larametrics_logs table in the laravel database
type LarametricsLogs struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] level                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Level string `gorm:"column:level;type:varchar;size:255;" json:"level"`
	//[ 2] message                                        text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Message string `gorm:"column:message;type:text;size:65535;" json:"message"`
	//[ 3] user_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UserID null.Int `gorm:"column:user_id;type:int;" json:"user_id"`
	//[ 4] email                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Email null.String `gorm:"column:email;type:varchar;size:255;" json:"email"`
	//[ 5] trace                                          text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Trace string `gorm:"column:trace;type:text;size:65535;" json:"trace"`
	//[ 6] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 7] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var larametrics_logsTableInfo = &TableInfo{
	Name: "larametrics_logs",
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
			Name:               "level",
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
			GoFieldName:        "Level",
			GoFieldType:        "string",
			JSONFieldName:      "level",
			ProtobufFieldName:  "level",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "message",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Message",
			GoFieldType:        "string",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
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
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "null.String",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "trace",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Trace",
			GoFieldType:        "string",
			JSONFieldName:      "trace",
			ProtobufFieldName:  "trace",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LarametricsLogs) TableName() string {
	return "larametrics_logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LarametricsLogs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LarametricsLogs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LarametricsLogs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LarametricsLogs) TableInfo() *TableInfo {
	return larametrics_logsTableInfo
}
