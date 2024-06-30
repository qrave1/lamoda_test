package service

import "github.com/qrave1/lamoda_test/internal/domain/model"

type ReservationService interface {
	ReserveProducts(codes []uint, warehouseID uint) error
	ReleaseProducts(codes []uint, warehouseID uint) error
	GetInventory(warehouseID uint) ([]model.Product, error)
}
