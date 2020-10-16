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


CREATE TABLE `larametrics_notifications` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `action` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `filter` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `meta` text COLLATE utf8mb4_unicode_ci,
  `notify_by` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'email',
  `last_fired_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "id": 82,    "action": "RXBuLQUnVtdREFcTeqIohdSTI",    "filter": "GTEGtGuXiWOfHZlXZYxVXgheP",    "meta": "fXFEXIBZfacBhIKUZJHLoTRrN",    "notify_by": "YxFIgbZyoApxQRfPilJJveXEi",    "last_fired_at": "2173-07-17T06:23:20.377116228-03:00",    "created_at": "2145-09-22T08:08:49.497645676-03:00",    "updated_at": "2266-05-08T21:59:43.078107934-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LarametricsNotifications struct is a row record of the larametrics_notifications table in the laravel database
type LarametricsNotifications struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] action                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Action string `gorm:"column:action;type:varchar;size:255;" json:"action"`
	//[ 2] filter                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Filter null.String `gorm:"column:filter;type:varchar;size:255;" json:"filter"`
	//[ 3] meta                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Meta null.String `gorm:"column:meta;type:text;size:65535;" json:"meta"`
	//[ 4] notify_by                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: [email]
	NotifyBy string `gorm:"column:notify_by;type:varchar;size:255;default:email;" json:"notify_by"`
	//[ 5] last_fired_at                                  timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	LastFiredAt null.Time `gorm:"column:last_fired_at;type:timestamp;" json:"last_fired_at"`
	//[ 6] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 7] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var larametrics_notificationsTableInfo = &TableInfo{
	Name: "larametrics_notifications",
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
			Name:               "action",
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
			GoFieldName:        "Action",
			GoFieldType:        "string",
			JSONFieldName:      "action",
			ProtobufFieldName:  "action",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "filter",
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
			GoFieldName:        "Filter",
			GoFieldType:        "null.String",
			JSONFieldName:      "filter",
			ProtobufFieldName:  "filter",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "meta",
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
			GoFieldName:        "Meta",
			GoFieldType:        "null.String",
			JSONFieldName:      "meta",
			ProtobufFieldName:  "meta",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "notify_by",
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
			GoFieldName:        "NotifyBy",
			GoFieldType:        "string",
			JSONFieldName:      "notify_by",
			ProtobufFieldName:  "notify_by",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "last_fired_at",
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
			GoFieldName:        "LastFiredAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_fired_at",
			ProtobufFieldName:  "last_fired_at",
			ProtobufType:       "uint64",
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
func (l *LarametricsNotifications) TableName() string {
	return "larametrics_notifications"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LarametricsNotifications) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LarametricsNotifications) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LarametricsNotifications) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LarametricsNotifications) TableInfo() *TableInfo {
	return larametrics_notificationsTableInfo
}
