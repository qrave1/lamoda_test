package repository

import (
	"context"

	"github.com/qrave1/lamoda_test/internal/domain/model"
)

type WarehouseRepository interface {
	Warehouse(ctx context.Context, id uint) (*model.Warehouse, error)
	CreateWarehouse(ctx context.Context, wh *model.Warehouse) error
}
