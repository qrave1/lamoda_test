package persistence

import (
	"database/sql"
	"errors"

	"github.com/qrave1/lamoda_test/internal/domain/model"
	"github.com/qrave1/lamoda_test/internal/domain/repository"
	"github.com/qrave1/lamoda_test/pkg/logger"
)

type WarehousePostgresRepository struct {
	db  *sql.DB
	log logger.Logger
}

var _ repository.WarehouseRepository = (*WarehousePostgresRepository)(nil)

func NewWarehousePostgresRepository(db *sql.DB, log logger.Logger) *WarehousePostgresRepository {
	return &WarehousePostgresRepository{db, log}
}

func (r *WarehousePostgresRepository) WarehouseById(tx *sql.Tx, id uint) (*model.Warehouse, error) {
	var warehouse model.Warehouse
	err := tx.QueryRow("SELECT id, name, is_available FROM warehouses WHERE id = $1", id).
		Scan(
			&warehouse.ID,
			&warehouse.Name,
			&warehouse.IsAvailable,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNoRowsFound
		}
		return nil, err
	}
	return &warehouse, nil
}

func (r *WarehousePostgresRepository) WarehouseByAvailable(tx *sql.Tx) (*model.Warehouse, error) {
	var warehouse model.Warehouse
	err := tx.QueryRow("SELECT id, name, is_available FROM warehouses WHERE is_available = $1", true).
		Scan(
			&warehouse.ID,
			&warehouse.Name,
			&warehouse.IsAvailable,
		)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}
