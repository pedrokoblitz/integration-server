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


CREATE TABLE `people` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `instagram` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `facebook` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `twitter` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `linkedin` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `birthdate` date DEFAULT NULL,
  `optin` datetime DEFAULT NULL,
  `optout` datetime DEFAULT NULL,
  `last_campaing_sent` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "facebook": "JpMTOQyaQkLFhpeQaKwNlMqjk",    "optout": "2133-06-27T12:15:53.13614602-03:00",    "created_at": "2299-08-24T05:06:11.734829332-03:00",    "id": 34,    "phone": "wRmTXdrTidhkCDqYogObniiDE",    "instagram": "JWVRNIepLByZruWwQrniGvFnU",    "last_campaing_sent": "2312-04-06T02:00:38.635438036-03:00",    "email": "icFFJHsEachiwlxIxjEkleqMd",    "name": "bEiApWqUeIvWCvsDgDymphjNJ",    "twitter": "AyOaHDGjHnnFKxpEpgZXbexPb",    "optin": "2063-09-13T22:10:23.521945907-03:00",    "linkedin": "KIZuYcLApduuvBNsmAmfHAKJS",    "birthdate": "2041-05-01T15:40:55.933900379-03:00",    "updated_at": "2064-11-08T13:21:14.925022763-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// People struct is a row record of the people table in the laravel database
type People struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] email                                          varchar(128)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	Email null.String `gorm:"column:email;type:varchar;size:128;" json:"email"`
	//[ 2] name                                           varchar(128)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 128     default: []
	Name null.String `gorm:"column:name;type:varchar;size:128;" json:"name"`
	//[ 3] phone                                          varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Phone null.String `gorm:"column:phone;type:varchar;size:32;" json:"phone"`
	//[ 4] instagram                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Instagram null.String `gorm:"column:instagram;type:varchar;size:255;" json:"instagram"`
	//[ 5] facebook                                       varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Facebook null.String `gorm:"column:facebook;type:varchar;size:255;" json:"facebook"`
	//[ 6] twitter                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Twitter null.String `gorm:"column:twitter;type:varchar;size:255;" json:"twitter"`
	//[ 7] linkedin                                       varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Linkedin null.String `gorm:"column:linkedin;type:varchar;size:255;" json:"linkedin"`
	//[ 8] birthdate                                      date                 null: true   primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	Birthdate null.Time `gorm:"column:birthdate;type:date;" json:"birthdate"`
	//[ 9] optin                                          datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Optin null.Time `gorm:"column:optin;type:datetime;" json:"optin"`
	//[10] optout                                         datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Optout null.Time `gorm:"column:optout;type:datetime;" json:"optout"`
	//[11] last_campaing_sent                             datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	LastCampaingSent null.Time `gorm:"column:last_campaing_sent;type:datetime;" json:"last_campaing_sent"`
	//[12] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[13] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var peopleTableInfo = &TableInfo{
	Name: "people",
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
			Name:               "email",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "Email",
			GoFieldType:        "null.String",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(128)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       128,
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "phone",
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
			GoFieldName:        "Phone",
			GoFieldType:        "null.String",
			JSONFieldName:      "phone",
			ProtobufFieldName:  "phone",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "instagram",
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
			GoFieldName:        "Instagram",
			GoFieldType:        "null.String",
			JSONFieldName:      "instagram",
			ProtobufFieldName:  "instagram",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "facebook",
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
			GoFieldName:        "Facebook",
			GoFieldType:        "null.String",
			JSONFieldName:      "facebook",
			ProtobufFieldName:  "facebook",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "twitter",
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
			GoFieldName:        "Twitter",
			GoFieldType:        "null.String",
			JSONFieldName:      "twitter",
			ProtobufFieldName:  "twitter",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "linkedin",
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
			GoFieldName:        "Linkedin",
			GoFieldType:        "null.String",
			JSONFieldName:      "linkedin",
			ProtobufFieldName:  "linkedin",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "birthdate",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "Birthdate",
			GoFieldType:        "null.Time",
			JSONFieldName:      "birthdate",
			ProtobufFieldName:  "birthdate",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "optin",
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
			GoFieldName:        "Optin",
			GoFieldType:        "null.Time",
			JSONFieldName:      "optin",
			ProtobufFieldName:  "optin",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "optout",
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
			GoFieldName:        "Optout",
			GoFieldType:        "null.Time",
			JSONFieldName:      "optout",
			ProtobufFieldName:  "optout",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "last_campaing_sent",
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
			GoFieldName:        "LastCampaingSent",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_campaing_sent",
			ProtobufFieldName:  "last_campaing_sent",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *People) TableName() string {
	return "people"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *People) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *People) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *People) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *People) TableInfo() *TableInfo {
	return peopleTableInfo
}
