package application

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/qrave1/lamoda_test/internal/domain/model"
	"github.com/qrave1/lamoda_test/internal/domain/repository"
	"github.com/qrave1/lamoda_test/pkg/logger"
)

type ReservationServiceImpl struct {
	db              *sql.DB
	productRepo     repository.ProductRepository
	warehouseRepo   repository.WarehouseRepository
	reservationRepo repository.ReservationRepository
	log             logger.Logger
}

func NewReservationServiceImpl(
	db *sql.DB,
	warehouseRepo repository.WarehouseRepository,
	productRepo repository.ProductRepository,
	reservationRepo repository.ReservationRepository,
	log logger.Logger,
) *ReservationServiceImpl {
	return &ReservationServiceImpl{
		db:              db,
		warehouseRepo:   warehouseRepo,
		productRepo:     productRepo,
		reservationRepo: reservationRepo,
		log:             log,
	}
}

func (r *ReservationServiceImpl) ReserveProducts(ctx context.Context, codes []uint, quantity uint) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return ErrBeginTx
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

	for _, code := range codes {
		product, err := r.productRepo.ProductByUniqueCode(tx, code)
		if err != nil {
			return err
		}

		if product.Quantity <= 0 || product.Quantity <= quantity {
			return errors.New("not enough quantity for product with code: " + strconv.Itoa(int(code)))
		}

		warehouse, err := r.warehouseRepo.WarehouseByAvailable(tx)
		if err != nil {
			return err
		}

		// todo проверить отрицательные значения
		product.Quantity -= quantity
		err = r.productRepo.UpdateQuantityByUniqueCode(tx, code, product.Quantity)
		if err != nil {
			return err
		}

		err = r.reservationRepo.CreateReservation(tx, product.ID, warehouse.ID, product.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ReservationServiceImpl) ReleaseProducts(ctx context.Context, codes []uint, warehouseID uint) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return ErrBeginTx
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

	for _, code := range codes {
		product, err := r.productRepo.ProductByUniqueCode(tx, code)
		if err != nil {
			return err
		}

		reserved, err := r.reservationRepo.DeleteReservation(tx, product.ID)
		if err != nil {
			return err
		}

		product.Quantity += reserved
		err = r.productRepo.UpdateQuantityByUniqueCode(tx, code, product.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ReservationServiceImpl) GetInventory(ctx context.Context, warehouseID uint) ([]model.Product, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, ErrBeginTx
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

	r.warehouseRepo.
}
