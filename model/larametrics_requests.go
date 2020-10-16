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


CREATE TABLE `larametrics_requests` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `method` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `headers` text COLLATE utf8mb4_unicode_ci,
  `start_time` double(16,4) DEFAULT NULL,
  `end_time` double(16,4) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "method": "YIAAIGOvNjQKPRDlXVdyDUCgA",    "ip": "PNIiuaqliXCKHURYfhPIaqXVx",    "start_time": 0.008995498798409373,    "end_time": 0.9395966036654653,    "created_at": "2065-03-03T00:36:39.423434275-03:00",    "updated_at": "2292-12-26T14:36:40.361644711-03:00",    "id": 34,    "uri": "COfPSOsSomMUdjGwkoACFmpmZ",    "headers": "NDuZrIPDbEtWVWQPkpOHgIEnO"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LarametricsRequests struct is a row record of the larametrics_requests table in the laravel database
type LarametricsRequests struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] method                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Method string `gorm:"column:method;type:varchar;size:255;" json:"method"`
	//[ 2] uri                                            text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	URI string `gorm:"column:uri;type:text;size:65535;" json:"uri"`
	//[ 3] ip                                             varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	IP null.String `gorm:"column:ip;type:varchar;size:255;" json:"ip"`
	//[ 4] headers                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Headers null.String `gorm:"column:headers;type:text;size:65535;" json:"headers"`
	//[ 5] start_time                                     double               null: true   primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	StartTime null.Float `gorm:"column:start_time;type:double;" json:"start_time"`
	//[ 6] end_time                                       double               null: true   primary: false  isArray: false  auto: false  col: double          len: -1      default: []
	EndTime null.Float `gorm:"column:end_time;type:double;" json:"end_time"`
	//[ 7] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 8] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var larametrics_requestsTableInfo = &TableInfo{
	Name: "larametrics_requests",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "uri",
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
			GoFieldName:        "URI",
			GoFieldType:        "string",
			JSONFieldName:      "uri",
			ProtobufFieldName:  "uri",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ip",
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
			GoFieldName:        "IP",
			GoFieldType:        "null.String",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "headers",
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
			GoFieldName:        "Headers",
			GoFieldType:        "null.String",
			JSONFieldName:      "headers",
			ProtobufFieldName:  "headers",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "start_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "null.Float",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "float",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "end_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "double",
			DatabaseTypePretty: "double",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "double",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "null.Float",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "float",
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
func (l *LarametricsRequests) TableName() string {
	return "larametrics_requests"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LarametricsRequests) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LarametricsRequests) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LarametricsRequests) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LarametricsRequests) TableInfo() *TableInfo {
	return larametrics_requestsTableInfo
}
