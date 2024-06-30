package repository

import (
	"context"

	"github.com/qrave1/lamoda_test/internal/domain/model"
)

type ProductRepository interface {
	ProductByUniqueCode(ctx context.Context, code uint) (*model.Product, error)
	UpdateQuantity(ctx context.Context, id uint, quantity int) error
}
