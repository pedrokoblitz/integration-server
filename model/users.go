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


CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `person_id` int(10) unsigned NOT NULL DEFAULT '0',
  `role` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "email_verified_at": "2209-02-26T20:43:23.361572122-03:00",    "remember_token": "IHKnffxsMXARlZateGSGHVSWC",    "updated_at": "2293-01-20T05:09:01.491400595-03:00",    "id": 14,    "person_id": 97,    "role": "nNctsBOSaoVnwhYVudAHboBRH",    "name": "pLOhyuUgcEExwDWjkNoQjpeKV",    "email": "cjjgZXMDkoBwPmWWrVcmHOZPI",    "password": "ovpkfsFYilAJbTcWhMdGMfXEO",    "created_at": "2175-05-25T16:13:37.565684538-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Users struct is a row record of the users table in the laravel database
type Users struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] person_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PersonID uint32 `gorm:"column:person_id;type:uint;default:0;" json:"person_id"`
	//[ 2] role                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: [customer]
	Role string `gorm:"column:role;type:varchar;size:255;default:customer;" json:"role"`
	//[ 3] name                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Name string `gorm:"column:name;type:varchar;size:255;" json:"name"`
	//[ 4] email                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Email string `gorm:"column:email;type:varchar;size:255;" json:"email"`
	//[ 5] email_verified_at                              timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	EmailVerifiedAt null.Time `gorm:"column:email_verified_at;type:timestamp;" json:"email_verified_at"`
	//[ 6] password                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Password string `gorm:"column:password;type:varchar;size:255;" json:"password"`
	//[ 7] remember_token                                 varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RememberToken null.String `gorm:"column:remember_token;type:varchar;size:100;" json:"remember_token"`
	//[ 8] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[ 9] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var usersTableInfo = &TableInfo{
	Name: "users",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "role",
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
			GoFieldName:        "Role",
			GoFieldType:        "string",
			JSONFieldName:      "role",
			ProtobufFieldName:  "role",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "email_verified_at",
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
			GoFieldName:        "EmailVerifiedAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "email_verified_at",
			ProtobufFieldName:  "email_verified_at",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "password",
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
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "remember_token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RememberToken",
			GoFieldType:        "null.String",
			JSONFieldName:      "remember_token",
			ProtobufFieldName:  "remember_token",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *Users) TableName() string {
	return "users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Users) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Users) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Users) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *Users) TableInfo() *TableInfo {
	return usersTableInfo
}
