package persistence

import (
	"database/sql"
	"log/slog"

	"github.com/qrave1/lamoda_test/internal/domain/model"
	"github.com/qrave1/lamoda_test/internal/domain/repository"
)

type WarehousePostgresRepository struct {
	db  *sql.DB
	log *slog.Logger
}

var _ repository.WarehouseRepository = (*WarehousePostgresRepository)(nil)

func NewWarehousePostgresRepository(db *sql.DB, log *slog.Logger) *WarehousePostgresRepository {
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
