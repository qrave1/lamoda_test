package persistence

import (
	"context"
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

func (r *WarehousePostgresRepository) Warehouse(ctx context.Context, id uint) (*model.Warehouse, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})
	if err != nil {
		return nil, TransactionStartError
	}
	defer func() {
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				r.log.Error("Transaction rollback error", "error", txErr)
			}
		} else {
			txErr := tx.Commit()
			if txErr != nil {
				r.log.Error("Transaction commit failed", "error", txErr)
			}
		}
	}()

	var warehouse model.Warehouse
	query := "SELECT id, name, is_available FROM warehouses WHERE id = $1"
	err = tx.QueryRow(query, id).Scan(&warehouse.ID, &warehouse.Name, &warehouse.IsAvailable)
	if err != nil {
		return nil, err
	}
	return &warehouse, nil
}

func (r *WarehousePostgresRepository) CreateWarehouse(ctx context.Context, wh *model.Warehouse) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return TransactionStartError
	}
	defer func() {
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				r.log.Error("Transaction rollback error", "error", txErr)
			}
		} else {
			txErr := tx.Commit()
			if txErr != nil {
				r.log.Error("Transaction commit failed", "error", txErr)
			}
		}
	}()

	query := "INSERT INTO warehouses (name, is_available) VALUES ($1, $2)"
	err = tx.QueryRow(query, wh.Name, wh.IsAvailable).Err()
	return err
}
