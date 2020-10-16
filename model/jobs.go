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


CREATE TABLE `jobs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `queue` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `payload` longtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `attempts` tinyint(3) unsigned NOT NULL,
  `reserved_at` int(10) unsigned DEFAULT NULL,
  `available_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `jobs_queue_index` (`queue`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "queue": "RVjsYmZcywhCDoMjYOCeAtuXa",    "payload": "ZhMDpAlBLsdLFAMHxFVPcyECe",    "attempts": 15,    "reserved_at": 44,    "available_at": 0,    "created_at": 86,    "id": 10}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// Jobs struct is a row record of the jobs table in the laravel database
type Jobs struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] queue                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Queue string `gorm:"column:queue;type:varchar;size:255;" json:"queue"`
	//[ 2] payload                                        text(4294967295)     null: false  primary: false  isArray: false  auto: false  col: text            len: 4294967295 default: []
	Payload string `gorm:"column:payload;type:text;size:4294967295;" json:"payload"`
	//[ 3] attempts                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: []
	Attempts uint32 `gorm:"column:attempts;type:utinyint;" json:"attempts"`
	//[ 4] reserved_at                                    uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	ReservedAt null.Int `gorm:"column:reserved_at;type:uint;" json:"reserved_at"`
	//[ 5] available_at                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AvailableAt uint32 `gorm:"column:available_at;type:uint;" json:"available_at"`
	//[ 6] created_at                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	CreatedAt uint32 `gorm:"column:created_at;type:uint;" json:"created_at"`
}

var jobsTableInfo = &TableInfo{
	Name: "jobs",
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
			Name:               "queue",
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
			GoFieldName:        "Queue",
			GoFieldType:        "string",
			JSONFieldName:      "queue",
			ProtobufFieldName:  "queue",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "payload",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(4294967295)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       4294967295,
			GoFieldName:        "Payload",
			GoFieldType:        "string",
			JSONFieldName:      "payload",
			ProtobufFieldName:  "payload",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "attempts",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Attempts",
			GoFieldType:        "uint32",
			JSONFieldName:      "attempts",
			ProtobufFieldName:  "attempts",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "reserved_at",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ReservedAt",
			GoFieldType:        "null.Int",
			JSONFieldName:      "reserved_at",
			ProtobufFieldName:  "reserved_at",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "available_at",
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
			GoFieldName:        "AvailableAt",
			GoFieldType:        "uint32",
			JSONFieldName:      "available_at",
			ProtobufFieldName:  "available_at",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "created_at",
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
			GoFieldName:        "CreatedAt",
			GoFieldType:        "uint32",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (j *Jobs) TableName() string {
	return "jobs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (j *Jobs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (j *Jobs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (j *Jobs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (j *Jobs) TableInfo() *TableInfo {
	return jobsTableInfo
}
