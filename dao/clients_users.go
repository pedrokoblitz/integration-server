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

// GetAllClientsUsers is a function to get a slice of record(s) from clients_users table in the laravel database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllClientsUsers(ctx context.Context, page, pagesize int64, order string) (results []*model.ClientsUsers, totalRows int, err error) {

	resultOrm := DB.Model(&model.ClientsUsers{})
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

// GetClientsUsers is a function to get a single record from the clients_users table in the laravel database
// error - ErrNotFound, db Find error
func GetClientsUsers(ctx context.Context, argId uint64) (record *model.ClientsUsers, err error) {
	record = &model.ClientsUsers{}
	if err = DB.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddClientsUsers is a function to add a single record to clients_users table in the laravel database
// error - ErrInsertFailed, db save call failed
func AddClientsUsers(ctx context.Context, record *model.ClientsUsers) (result *model.ClientsUsers, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateClientsUsers is a function to update a single record from clients_users table in the laravel database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateClientsUsers(ctx context.Context, argId uint64, updated *model.ClientsUsers) (result *model.ClientsUsers, RowsAffected int64, err error) {

	result = &model.ClientsUsers{}
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

// DeleteClientsUsers is a function to delete a single record from clients_users table in the laravel database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteClientsUsers(ctx context.Context, argId uint64) (rowsAffected int64, err error) {

	record := &model.ClientsUsers{}
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
