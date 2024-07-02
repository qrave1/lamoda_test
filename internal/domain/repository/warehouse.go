package repository

import (
	"database/sql"

	"github.com/qrave1/lamoda_test/internal/domain/model"
)

type WarehouseRepository interface {
	WarehouseById(tx *sql.Tx, id uint) (*model.Warehouse, error)
	WarehouseByAvailable(tx *sql.Tx) (*model.Warehouse, error)
}
