package persistence

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/qrave1/lamoda_test/internal/domain/model"
	"github.com/qrave1/lamoda_test/internal/domain/repository"
)

type PostgresProductRepository struct {
	db  *sql.DB
	log *slog.Logger
}

var _ repository.ProductRepository = (*PostgresProductRepository)(nil)

func NewPostgresProductRepository(db *sql.DB, log *slog.Logger) *PostgresProductRepository {
	return &PostgresProductRepository{db, log}
}

func (r *PostgresProductRepository) ProductByUniqueCode(ctx context.Context, code uint) (*model.Product, error) {
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

	var product model.Product
	err = tx.QueryRow("SELECT id, name, size, unique_code, quantity FROM products WHERE unique_code = $1", code).
		Scan(
			&product.ID,
			&product.Name,
			&product.Size,
			&product.UniqueCode,
			&product.Quantity,
		)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *PostgresProductRepository) UpdateQuantity(ctx context.Context, id uint, quantity int) error {
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

	_, err = tx.Exec("UPDATE products SET quantity = $1 WHERE id = $2", quantity, id)
	return err
}
