package repository

import (
	"context"

	"github.com/qrave1/lamoda_test/internal/domain/model"
)

type ReservationRepository interface {
	Reservation(ctx context.Context, productID uint, warehouseID uint) (*model.Reservation, error)
	CreateReservation(ctx context.Context, r *model.Reservation) error
	//UpdateReservationQuantity(id uint, quantity int) error
	DeleteReservation(ctx context.Context, id uint) error
}
