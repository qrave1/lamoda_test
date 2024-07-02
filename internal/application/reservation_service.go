package application

import (
	"context"
	"database/sql"
	"errors"

	"github.com/qrave1/lamoda_test/internal/domain/repository"
	"github.com/qrave1/lamoda_test/internal/domain/service"
	"github.com/qrave1/lamoda_test/pkg/logger"
)

type ReservationServiceImpl struct {
	db                   *sql.DB
	productRepo          repository.ProductRepository
	warehouseRepo        repository.WarehouseRepository
	productWarehouseRepo repository.ProductWarehouseRepository
	log                  logger.Logger
}

func NewReservationServiceImpl(
	db *sql.DB,
	warehouseRepo repository.WarehouseRepository,
	productRepo repository.ProductRepository,
	productWarehouseRepo repository.ProductWarehouseRepository,
	log logger.Logger,
) *ReservationServiceImpl {
	return &ReservationServiceImpl{
		db:                   db,
		warehouseRepo:        warehouseRepo,
		productRepo:          productRepo,
		productWarehouseRepo: productWarehouseRepo,
		log:                  log,
	}
}

func (r *ReservationServiceImpl) ReserveProducts(ctx context.Context, codes []uint, warehouseID, quantity uint) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return service.ErrBeginTx
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

	wh, err := r.warehouseRepo.WarehouseById(tx, warehouseID)
	if err != nil {
		if errors.Is(err, repository.ErrNoRowsFound) {
			return service.NewApplicationError("warehouse not found", 404)
		} else {
			return err
		}
	}

	if !wh.IsAvailable {
		return service.NewApplicationError("warehouse not available", 400)
	}

	for _, code := range codes {
		product, err := r.productRepo.ProductByUniqueCode(tx, code)
		if err != nil {
			if errors.Is(err, repository.ErrNoRowsFound) {
				return service.NewApplicationError("product not found", 404)
			} else {
				return err
			}
		}

		err = r.productWarehouseRepo.Reserve(tx, product.ID, warehouseID, quantity)
		if err != nil {
			return err
		}
	}

	return err
}

func (r *ReservationServiceImpl) ReleaseProducts(ctx context.Context, codes []uint, warehouseID uint) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return service.ErrBeginTx
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
			if errors.Is(err, repository.ErrNoRowsFound) {
				return service.NewApplicationError("product not found", 404)
			} else {
				return err
			}
		}

		err = r.productWarehouseRepo.Release(tx, product.ID, warehouseID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ReservationServiceImpl) Inventory(ctx context.Context, warehouseID uint) (uint, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, service.ErrBeginTx
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

	wh, err := r.warehouseRepo.WarehouseById(tx, warehouseID)
	if err != nil {
		if errors.Is(err, repository.ErrNoRowsFound) {
			return 0, service.NewApplicationError("warehouse not found", 404)
		} else {
			return 0, err
		}
	}

	if !wh.IsAvailable {
		return 0, service.NewApplicationError("warehouse not available", 400)
	}

	return r.productWarehouseRepo.ProductsInWarehouse(tx, warehouseID)
}
