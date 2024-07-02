package persistence

import (
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

func (r *PostgresProductRepository) ProductByUniqueCode(tx *sql.Tx, code uint) (*model.Product, error) {
	var product model.Product
	err := tx.QueryRow("SELECT id, name, size, code, quantity FROM products WHERE code = $1", code).
		Scan(
			&product.ID,
			&product.Name,
			&product.Size,
			&product.Code,
			&product.Quantity,
		)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *PostgresProductRepository) UpdateQuantityByUniqueCode(tx *sql.Tx, code uint, quantity uint) error {
	res, err := tx.Exec("UPDATE products SET quantity = $1 WHERE code = $2", quantity, code)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 && err == nil {
		return ErrNoRowsAffected
	}
	return err
}
