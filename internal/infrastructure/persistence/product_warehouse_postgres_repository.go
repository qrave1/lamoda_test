package persistence

import (
	"database/sql"
	"errors"

	"github.com/qrave1/lamoda_test/internal/domain/repository"
	"github.com/qrave1/lamoda_test/pkg/logger"
)

type ProductWarehousePostgresRepository struct {
	db  *sql.DB
	log logger.Logger
}

var _ repository.ProductWarehouseRepository = (*ProductWarehousePostgresRepository)(nil)

func NewProductWarehousePostgresRepository(db *sql.DB, log logger.Logger) *ProductWarehousePostgresRepository {
	return &ProductWarehousePostgresRepository{db: db, log: log}
}

func (pw *ProductWarehousePostgresRepository) Reserve(tx *sql.Tx, productID, warehouseID, quantity uint) error {
	res, err := tx.Exec(
		`UPDATE product_warehouse SET 
                             quantity = quantity - $1,
                             reserved_quantity = reserved_quantity + $1,
                             updated_at = CURRENT_TIMESTAMP
								WHERE product_id = $2 AND warehouse_id = $3`,
		quantity,
		productID,
		warehouseID,
	)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 && err == nil {
		return repository.ErrNoRowsAffected
	} else if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrNoRowsFound
	}
	return err
}

func (pw *ProductWarehousePostgresRepository) Release(tx *sql.Tx, productID, warehouseID uint) error {
	res, err := tx.Exec(
		`UPDATE product_warehouse SET 
                             quantity = quantity + reserved_quantity,
                             reserved_quantity = 0,
                             updated_at = CURRENT_TIMESTAMP
								WHERE product_id = $1 AND warehouse_id = $2`,
		productID,
		warehouseID,
	)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 && err == nil {
		return repository.ErrNoRowsAffected
	} else if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrNoRowsFound
	}
	return err
}

func (pw *ProductWarehousePostgresRepository) ProductsInWarehouse(tx *sql.Tx, warehouseID uint) (count uint, err error) {
	err = tx.QueryRow("SELECT sum(quantity) FROM product_warehouse WHERE warehouse_id = $1", warehouseID).
		Scan(
			&count,
		)
	return
}
