package persistence

import (
	"database/sql"
	"errors"

	"github.com/qrave1/lamoda_test/internal/domain/model"
	"github.com/qrave1/lamoda_test/internal/domain/repository"
	"github.com/qrave1/lamoda_test/pkg/logger"
)

type ProductPostgresRepository struct {
	db  *sql.DB
	log logger.Logger
}

var _ repository.ProductRepository = (*ProductPostgresRepository)(nil)

func NewProductPostgresRepository(db *sql.DB, log logger.Logger) *ProductPostgresRepository {
	return &ProductPostgresRepository{db, log}
}

func (pr *ProductPostgresRepository) ProductByUniqueCode(tx *sql.Tx, code uint) (*model.Product, error) {
	var product model.Product
	err := tx.QueryRow("SELECT id, name, size, code FROM products WHERE code = $1", code).
		Scan(
			&product.ID,
			&product.Name,
			&product.Size,
			&product.Code,
		)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNoRowsFound
		}
		return nil, err
	}
	return &product, nil
}

//func (pr *ProductPostgresRepository) ProductsByWarehouse(tx *sql.Tx, warehouseID uint) (
//	res []*model.Product,
//	err error,
//) {
//	rows, err := tx.Query(`SELECT p.id, p.name, p.code, pw.quantity, p.size FROM product_warehouse AS pw
//    								INNER JOIN products p
//    								    ON p.id = pw.product_id
//    								    WHERE pw.warehouse_id = $1`, warehouseID,
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	for rows.Next() {
//		var product model.Product
//		err = rows.Scan(&product.ID, product.Name, &product.Code, &product.Size)
//		if err != nil {
//			return nil, err
//		}
//		res = append(res, &product)
//	}
//
//	return
//}
