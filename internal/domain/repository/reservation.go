package repository

import (
	"database/sql"
)

type ReservationRepository interface {
	CreateReservation(tx *sql.Tx, productID, warehouseID, quantity uint) error
	DeleteReservation(tx *sql.Tx, productID uint) (uint, error)
}
