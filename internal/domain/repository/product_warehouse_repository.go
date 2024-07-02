package repository

import (
	"database/sql"
)

type ProductWarehouseRepository interface {
	Reserve(tx *sql.Tx, productID, warehouseID, quantity uint) error
	Release(tx *sql.Tx, productID, warehouseID uint) error
	ProductsInWarehouse(tx *sql.Tx, warehouseID uint) (uint, error)
}
