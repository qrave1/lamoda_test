package repository

import (
	"database/sql"

	"github.com/qrave1/lamoda_test/internal/domain/model"
)

type ProductRepository interface {
	ProductByUniqueCode(tx *sql.Tx, code uint) (*model.Product, error)
	UpdateQuantityByUniqueCode(tx *sql.Tx, code uint, quantity uint) error
}
