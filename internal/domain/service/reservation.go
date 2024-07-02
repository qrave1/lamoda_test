package service

import (
	"context"

	"github.com/qrave1/lamoda_test/internal/domain/model"
)

type ReservationService interface {
	ReserveProducts(ctx context.Context, codes []uint, warehouseID uint) error
	ReleaseProducts(ctx context.Context, codes []uint, warehouseID uint) error
	GetInventory(ctx context.Context, warehouseID uint) ([]model.Product, error)
}
