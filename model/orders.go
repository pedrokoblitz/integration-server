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


CREATE TABLE `orders` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL,
  `ref_code` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `shipping_status_id` int(10) unsigned NOT NULL DEFAULT '1',
  `payment_status_id` int(10) unsigned NOT NULL DEFAULT '1',
  `customer_id` int(10) unsigned NOT NULL DEFAULT '0',
  `shipping_address_id` int(10) unsigned NOT NULL DEFAULT '0',
  `coupon_id` int(10) unsigned NOT NULL DEFAULT '0',
  `value` int(11) NOT NULL DEFAULT '0',
  `shipping` int(11) NOT NULL DEFAULT '0',
  `discount` int(11) NOT NULL DEFAULT '0',
  `total` int(11) NOT NULL DEFAULT '0',
  `receiver` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `payment_tracking` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `shipping_tracking` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `shipped` datetime DEFAULT NULL,
  `expected` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

JSON Sample
-------------------------------------
{    "value": 12,    "receiver": "eFgPyuiBjNlFgGvSAZrRyHvRU",    "shipping_tracking": "GejqyuVQspIuRwSvfalZWpeXT",    "created_at": "2192-04-20T21:48:19.96917947-03:00",    "app_id": 40,    "payment_status_id": 14,    "customer_id": 87,    "id": 96,    "shipping_address_id": 20,    "shipped": "2082-08-16T20:48:23.276826206-03:00",    "discount": 49,    "payment_tracking": "hhXtRfPNxAvWlKdTHDxtpsdWi",    "ref_code": "fCylxirPlEgORXJRYpUfFQCeJ",    "shipping_status_id": 68,    "coupon_id": 4,    "updated_at": "2292-02-20T21:11:29.860446336-03:00",    "shipping": 95,    "total": 82,    "expected": "2265-07-18T15:36:08.296380689-03:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned



*/

// Orders struct is a row record of the orders table in the laravel database
type Orders struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] app_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AppID uint32 `gorm:"column:app_id;type:uint;" json:"app_id"`
	//[ 2] ref_code                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	RefCode string `gorm:"column:ref_code;type:varchar;size:64;" json:"ref_code"`
	//[ 3] shipping_status_id                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	ShippingStatusID uint32 `gorm:"column:shipping_status_id;type:uint;default:1;" json:"shipping_status_id"`
	//[ 4] payment_status_id                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	PaymentStatusID uint32 `gorm:"column:payment_status_id;type:uint;default:1;" json:"payment_status_id"`
	//[ 5] customer_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CustomerID uint32 `gorm:"column:customer_id;type:uint;default:0;" json:"customer_id"`
	//[ 6] shipping_address_id                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ShippingAddressID uint32 `gorm:"column:shipping_address_id;type:uint;default:0;" json:"shipping_address_id"`
	//[ 7] coupon_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CouponID uint32 `gorm:"column:coupon_id;type:uint;default:0;" json:"coupon_id"`
	//[ 8] value                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Value int32 `gorm:"column:value;type:int;default:0;" json:"value"`
	//[ 9] shipping                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Shipping int32 `gorm:"column:shipping;type:int;default:0;" json:"shipping"`
	//[10] discount                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Discount int32 `gorm:"column:discount;type:int;default:0;" json:"discount"`
	//[11] total                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Total int32 `gorm:"column:total;type:int;default:0;" json:"total"`
	//[12] receiver                                       varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Receiver null.String `gorm:"column:receiver;type:varchar;size:255;" json:"receiver"`
	//[13] payment_tracking                               varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	PaymentTracking null.String `gorm:"column:payment_tracking;type:varchar;size:255;" json:"payment_tracking"`
	//[14] shipping_tracking                              varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ShippingTracking null.String `gorm:"column:shipping_tracking;type:varchar;size:255;" json:"shipping_tracking"`
	//[15] shipped                                        datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Shipped null.Time `gorm:"column:shipped;type:datetime;" json:"shipped"`
	//[16] expected                                       datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	Expected null.Time `gorm:"column:expected;type:datetime;" json:"expected"`
	//[17] created_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	CreatedAt null.Time `gorm:"column:created_at;type:timestamp;" json:"created_at"`
	//[18] updated_at                                     timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	UpdatedAt null.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}

var ordersTableInfo = &TableInfo{
	Name: "orders",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ref_code",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "RefCode",
			GoFieldType:        "string",
			JSONFieldName:      "ref_code",
			ProtobufFieldName:  "ref_code",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "shipping_status_id",
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
			GoFieldName:        "ShippingStatusID",
			GoFieldType:        "uint32",
			JSONFieldName:      "shipping_status_id",
			ProtobufFieldName:  "shipping_status_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "payment_status_id",
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
			GoFieldName:        "PaymentStatusID",
			GoFieldType:        "uint32",
			JSONFieldName:      "payment_status_id",
			ProtobufFieldName:  "payment_status_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "customer_id",
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
			GoFieldName:        "CustomerID",
			GoFieldType:        "uint32",
			JSONFieldName:      "customer_id",
			ProtobufFieldName:  "customer_id",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "shipping_address_id",
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
			GoFieldName:        "ShippingAddressID",
			GoFieldType:        "uint32",
			JSONFieldName:      "shipping_address_id",
			ProtobufFieldName:  "shipping_address_id",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "coupon_id",
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
			GoFieldName:        "CouponID",
			GoFieldType:        "uint32",
			JSONFieldName:      "coupon_id",
			ProtobufFieldName:  "coupon_id",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "value",
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
			GoFieldName:        "Value",
			GoFieldType:        "int32",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "shipping",
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
			GoFieldName:        "Shipping",
			GoFieldType:        "int32",
			JSONFieldName:      "shipping",
			ProtobufFieldName:  "shipping",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "discount",
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
			GoFieldName:        "Discount",
			GoFieldType:        "int32",
			JSONFieldName:      "discount",
			ProtobufFieldName:  "discount",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "total",
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
			GoFieldName:        "Total",
			GoFieldType:        "int32",
			JSONFieldName:      "total",
			ProtobufFieldName:  "total",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "receiver",
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
			GoFieldName:        "Receiver",
			GoFieldType:        "null.String",
			JSONFieldName:      "receiver",
			ProtobufFieldName:  "receiver",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "payment_tracking",
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
			GoFieldName:        "PaymentTracking",
			GoFieldType:        "null.String",
			JSONFieldName:      "payment_tracking",
			ProtobufFieldName:  "payment_tracking",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "shipping_tracking",
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
			GoFieldName:        "ShippingTracking",
			GoFieldType:        "null.String",
			JSONFieldName:      "shipping_tracking",
			ProtobufFieldName:  "shipping_tracking",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "shipped",
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
			GoFieldName:        "Shipped",
			GoFieldType:        "null.Time",
			JSONFieldName:      "shipped",
			ProtobufFieldName:  "shipped",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "expected",
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
			GoFieldName:        "Expected",
			GoFieldType:        "null.Time",
			JSONFieldName:      "expected",
			ProtobufFieldName:  "expected",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *Orders) TableName() string {
	return "orders"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Orders) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Orders) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Orders) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *Orders) TableInfo() *TableInfo {
	return ordersTableInfo
}
