package persistence

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/qrave1/lamoda_test/internal/domain/model"
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

func (r *ReservationPostgresRepository) Reservation(ctx context.Context, productID uint, warehouseID uint) (*model.Reservation, error) {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: true,
	})
	if err != nil {
		return nil, TransactionStartError
	}
	defer func() {
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				r.log.Error("Transaction rollback error", "error", txErr)
			}
		} else {
			txErr := tx.Commit()
			if txErr != nil {
				r.log.Error("Transaction commit failed", "error", txErr)
			}
		}
	}()

	var reservation model.Reservation
	err = tx.QueryRow(
		`SELECT id, product_id, warehouse_id, reserved_quantity FROM reservations 
                                                       WHERE product_id = $1 AND warehouse_id = $2`,
		productID,
		warehouseID,
	).
		Scan(
			&reservation.ID,
			&reservation.ProductID,
			&reservation.WarehouseID,
			&reservation.ReservedQuantity,
		)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NoRowsFoundError
	}

	return &reservation, nil
}

func (r *ReservationPostgresRepository) CreateReservation(ctx context.Context, reservation *model.Reservation) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return TransactionStartError
	}
	defer func() {
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				r.log.Error("Transaction rollback error", "error", txErr)
			}
		} else {
			txErr := tx.Commit()
			if txErr != nil {
				r.log.Error("Transaction commit failed", "error", txErr)
			}
		}
	}()

	query, err := tx.Query("SELECT warehouses.id")
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		"INSERT INTO reservations (product_id, warehouse_id, reserved_quantity) VALUES ($1, $2, $3)",
		reservation.ProductID,
		reservation.WarehouseID,
		reservation.ReservedQuantity,
	)
	return err
}

//func (r *ReservationPostgresRepository) UpdateReservationQuantity(id uint, quantity int) error {
//	tx, err := r.db.Begin()
//	if err != nil {
//		return TransactionStartError
//	}
//	defer func() {
//		if err != nil {
//			txErr := tx.Rollback()
//			if txErr != nil {
//				r.log.Error("Transaction rollback error", "error", txErr)
//			}
//		} else {
//			txErr := tx.Commit()
//			if txErr != nil {
//				r.log.Error("Transaction commit failed", "error", txErr)
//			}
//		}
//	}()
//	_, err = tx.Exec(
//		"UPDATE reservations SET reserved_quantity = $1 WHERE id = $2",
//		quantity,
//		id,
//	)
//	return err
//}

func (r *ReservationPostgresRepository) DeleteReservation(ctx context.Context, id uint) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		ReadOnly: false,
	})
	if err != nil {
		return TransactionStartError
	}
	defer func() {
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				r.log.Error("Transaction rollback error", "error", txErr)
			}
		} else {
			txErr := tx.Commit()
			if txErr != nil {
				r.log.Error("Transaction commit failed", "error", txErr)
			}
		}
	}()

	_, err = tx.Exec("DELETE FROM reservations WHERE id = $1", id)
	return err
}
