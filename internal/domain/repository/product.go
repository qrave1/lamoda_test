package repository

import (
	"database/sql"

	"github.com/qrave1/lamoda_test/internal/domain/model"
)

type ProductRepository interface {
	ProductByUniqueCode(tx *sql.Tx, code uint) (*model.Product, error)

	// если захочется возвращать не кол-во, а сами товары
	//ProductsByWarehouse(tx *sql.Tx, warehouseID uint) ([]*model.Product, error)
}
