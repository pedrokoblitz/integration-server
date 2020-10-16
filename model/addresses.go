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


CREATE TABLE `addresses` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `type_id` int(10) unsigned NOT NULL DEFAULT '1',
  `name` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `street` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `number` varchar(8) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `address_line` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `district` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `city` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `province` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `country_code` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'BRA',
  `zipcode` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "type_id": 2,    "name": "dxLidIxyKdIWCOlOEWfqJksjL",    "address_line": "exmhGQlQZFYuiDROhgldYNiNM",    "city": "dNJhWEAtVNQnLpSIeJDNKinFR",    "province": "MZVMTHQJmuAsZlqBpNYEePEtg",    "zipcode": "ZDrCNItUBARRRDCgvoBHcyPSx",    "id": 69,    "street": "yPbJZQWPSjkWVWJgNVMLCCjWg",    "number": "BukntDtJgWPmKXaAKCUuSwAyW",    "district": "qcsqyipVgGsoMhQgxjProTobl",    "country_code": "VccJhFuFYfCNWrBqXZPtxMcZf",    "created_at": "2263-10-13T17:58:40.465592086-03:00",    "updated_at": "2259-10-24T18:39:36.915712497-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned



*/

// Addresses struct is a row record of the addresses table in the laravel database
type Addresses struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] type_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	TypeID uint32 `gorm:"column:type_id;type:uint;default:1;" json:"type_id"`
	//[ 2] name                                           varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Name null.String `gorm:"column:name;type:varchar;size:32;" json:"name"`
	//[ 3] street                                         varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Street null.String `gorm:"column:street;type:varchar;size:32;" json:"street"`
	//[ 4] number                                         varchar(8)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 8       default: []
	Number null.String `gorm:"column:number;type:varchar;size:8;" json:"number"`
	//[ 5] address_line                                   varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AddressLine null.String `gorm:"column:address_line;type:varchar;size:255;" json:"address_line"`
	//[ 6] district                                       varchar(16)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 16      default: []
	District null.String `gorm:"column:district;type:varchar;size:16;" json:"district"`
	//[ 7] city                                           varchar(16)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 16      default: []
	City null.String `gorm:"column:city;type:varchar;size:16;" json:"city"`
	//[ 8] province                                       varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	Province null.String `gorm:"column:province;type:varchar;size:2;" json:"province"`
	//[ 9] country_code                                   varchar(3)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 3       default: [BRA]
	CountryCode string `gorm:"column:country_code;type:varchar;size:3;default:BRA;" json:"country_code"`
	//[10] zipcode                                        varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	Zipcode null.String `gorm:"column:zipcode;type:varchar;size:10;" json:"zipcode"`
	//[11] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[12] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var addressesTableInfo = &TableInfo{
	Name: "addresses",
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
			Name:               "type_id",
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
			GoFieldName:        "TypeID",
			GoFieldType:        "uint32",
			JSONFieldName:      "type_id",
			ProtobufFieldName:  "type_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "name",
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
			GoFieldName:        "Name",
			GoFieldType:        "null.String",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "street",
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
			GoFieldName:        "Street",
			GoFieldType:        "null.String",
			JSONFieldName:      "street",
			ProtobufFieldName:  "street",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "number",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(8)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       8,
			GoFieldName:        "Number",
			GoFieldType:        "null.String",
			JSONFieldName:      "number",
			ProtobufFieldName:  "number",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "address_line",
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
			GoFieldName:        "AddressLine",
			GoFieldType:        "null.String",
			JSONFieldName:      "address_line",
			ProtobufFieldName:  "address_line",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "district",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       16,
			GoFieldName:        "District",
			GoFieldType:        "null.String",
			JSONFieldName:      "district",
			ProtobufFieldName:  "district",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "city",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       16,
			GoFieldName:        "City",
			GoFieldType:        "null.String",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "province",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "Province",
			GoFieldType:        "null.String",
			JSONFieldName:      "province",
			ProtobufFieldName:  "province",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "country_code",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(3)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       3,
			GoFieldName:        "CountryCode",
			GoFieldType:        "string",
			JSONFieldName:      "country_code",
			ProtobufFieldName:  "country_code",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "zipcode",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "Zipcode",
			GoFieldType:        "null.String",
			JSONFieldName:      "zipcode",
			ProtobufFieldName:  "zipcode",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Addresses) TableName() string {
	return "addresses"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Addresses) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Addresses) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Addresses) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Addresses) TableInfo() *TableInfo {
	return addressesTableInfo
}
