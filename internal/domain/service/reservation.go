package service

import (
	"context"
)

type ReservationService interface {
	ReserveProducts(ctx context.Context, codes []uint, warehouseID, quantity uint) error
	ReleaseProducts(ctx context.Context, codes []uint, warehouseID uint) error
	Inventory(ctx context.Context, warehouseID uint) (uint, error)
}
