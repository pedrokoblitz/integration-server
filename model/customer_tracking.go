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


CREATE TABLE `customer_tracking` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `person_id` int(10) unsigned NOT NULL DEFAULT '0',
  `campaign_id` int(10) unsigned NOT NULL DEFAULT '0',
  `is_confirmed` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "user_token": "QrqnEyRrRFFdRkeBXQwjleZHl",    "person_id": 9,    "campaign_id": 49,    "is_confirmed": 46,    "created_at": "2121-08-12T17:42:36.274308744-03:00",    "updated_at": "2241-05-25T08:03:26.030502081-03:00",    "id": 18}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// CustomerTracking struct is a row record of the customer_tracking table in the laravel database
type CustomerTracking struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] user_token                                     varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	UserToken string `gorm:"column:user_token;type:varchar;size:255;" json:"user_token"`
	//[ 2] person_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PersonID uint32 `gorm:"column:person_id;type:uint;default:0;" json:"person_id"`
	//[ 3] campaign_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CampaignID uint32 `gorm:"column:campaign_id;type:uint;default:0;" json:"campaign_id"`
	//[ 4] is_confirmed                                   tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsConfirmed int32 `gorm:"column:is_confirmed;type:tinyint;default:0;" json:"is_confirmed"`
	//[ 5] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 6] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var customer_trackingTableInfo = &TableInfo{
	Name: "customer_tracking",
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
			Name:               "user_token",
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
			GoFieldName:        "UserToken",
			GoFieldType:        "string",
			JSONFieldName:      "user_token",
			ProtobufFieldName:  "user_token",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "is_confirmed",
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
			GoFieldName:        "IsConfirmed",
			GoFieldType:        "int32",
			JSONFieldName:      "is_confirmed",
			ProtobufFieldName:  "is_confirmed",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CustomerTracking) TableName() string {
	return "customer_tracking"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CustomerTracking) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CustomerTracking) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CustomerTracking) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CustomerTracking) TableInfo() *TableInfo {
	return customer_trackingTableInfo
}
