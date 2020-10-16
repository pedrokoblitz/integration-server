package dao

import (
	"context"
	"time"

	"github.com/pedrokoblitz/comm/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllAddressTypes is a function to get a slice of record(s) from address_types table in the laravel database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllAddressTypes(ctx context.Context, page, pagesize int64, order string) (results []*model.AddressTypes, totalRows int, err error) {

	resultOrm := DB.Model(&model.AddressTypes{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetAddressTypes is a function to get a single record from the address_types table in the laravel database
// error - ErrNotFound, db Find error
func GetAddressTypes(ctx context.Context, argId uint64) (record *model.AddressTypes, err error) {
	record = &model.AddressTypes{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddAddressTypes is a function to add a single record to address_types table in the laravel database
// error - ErrInsertFailed, db save call failed
func AddAddressTypes(ctx context.Context, record *model.AddressTypes) (result *model.AddressTypes, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateAddressTypes is a function to update a single record from address_types table in the laravel database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateAddressTypes(ctx context.Context, argId uint64, updated *model.AddressTypes) (result *model.AddressTypes, RowsAffected int64, err error) {

	result = &model.AddressTypes{}
	db := DB.First(result, argId)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteAddressTypes is a function to delete a single record from address_types table in the laravel database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteAddressTypes(ctx context.Context, argId uint64) (rowsAffected int64, err error) {

	record := &model.AddressTypes{}
	db := DB.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
