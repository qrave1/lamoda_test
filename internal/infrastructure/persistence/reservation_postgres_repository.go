package persistence

import (
	"database/sql"
	"log/slog"

	"github.com/qrave1/lamoda_test/internal/domain/repository"
)

type ReservationPostgresRepository struct {
	db  *sql.DB
	log *slog.Logger
}

var _ repository.ReservationRepository = (*ReservationPostgresRepository)(nil)

func NewReservationPostgresRepository(db *sql.DB, log *slog.Logger) *ReservationPostgresRepository {
	return &ReservationPostgresRepository{db: db, log: log}
}

func (r *ReservationPostgresRepository) CreateReservation(tx *sql.Tx, productID, warehouseID, quantity uint) error {
	res, err := tx.Exec(
		"INSERT INTO reservations (product_id, warehouse_id, reserved_quantity) VALUES ($1, $2, $3)",
		productID,
		warehouseID,
		quantity,
	)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 && err == nil {
		return ErrNoRowsAffected
	}
	return err
}

func (r *ReservationPostgresRepository) DeleteReservation(tx *sql.Tx, productID uint) (
	reserved uint,
	err error,
) {
	err = tx.QueryRow(
		"DELETE FROM reservations WHERE product_id = $1 RETURNING reserved_quantity",
		productID,
	).Scan(&reserved)
	return reserved, err
}
