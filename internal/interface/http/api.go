package http

import (
	"context"
	"errors"
	"net/http"

	"github.com/qrave1/lamoda_test/internal/domain/service"
	"github.com/qrave1/lamoda_test/internal/interface/http/gen"
	"github.com/qrave1/lamoda_test/pkg/logger"
)

type API struct {
	reservationService service.ReservationService
	log                logger.Logger
}

func NewAPI(rs service.ReservationService, log logger.Logger) *API {
	return &API{
		reservationService: rs,
		log:                log,
	}
}

func (a *API) GetInventory(ctx context.Context, r gen.GetInventoryRequestObject) (gen.GetInventoryResponseObject, error) {
	requestID := RequestIDFromContext(ctx)

	count, err := a.reservationService.Inventory(ctx, uint(r.Params.WarehouseId))
	if err != nil {
		a.log.Warn("error getting inventory", "error", err, "requestID", requestID)

		var appErr service.ApplicationError
		if errors.As(err, &appErr) {
			switch appErr.Status() {
			case http.StatusBadRequest:
				return gen.GetInventory400JSONResponse{Error: err.Error()}, nil
			case http.StatusNotFound:
				return gen.GetInventory404Response{}, nil
			default:
				return gen.GetInventory500Response{}, nil
			}
		}
	}

	return gen.GetInventory200JSONResponse(count), nil
}

func (a *API) PostRelease(ctx context.Context, r gen.PostReleaseRequestObject) (gen.PostReleaseResponseObject, error) {
	requestID := RequestIDFromContext(ctx)

	if !r.Valid() {
		return gen.PostRelease400JSONResponse{Error: "not valid body"}, nil
	}

	err := a.reservationService.ReleaseProducts(ctx, castToUint(r.Body.ProductCodes), uint(r.Body.WarehouseId))
	if err != nil {
		a.log.Warn("error posting release", "error", err, "requestID", requestID)

		var appErr service.ApplicationError
		if errors.As(err, &appErr) {
			switch appErr.Status() {
			case http.StatusBadRequest:
				return gen.PostRelease400JSONResponse{Error: err.Error()}, nil
			case http.StatusNotFound:
				return gen.PostRelease404Response{}, nil
			default:
				return gen.PostRelease500Response{}, nil
			}
		}
	}

	return gen.PostRelease200JSONResponse{Message: "OK"}, nil
}

func (a *API) PostReserve(ctx context.Context, r gen.PostReserveRequestObject) (gen.PostReserveResponseObject, error) {
	requestID := RequestIDFromContext(ctx)

	if !r.Valid() {
		return gen.PostReserve400JSONResponse{Error: "not valid body"}, nil
	}

	err := a.reservationService.ReserveProducts(
		ctx,
		castToUint(r.Body.ProductCodes),
		uint(r.Body.WarehouseId),
		uint(r.Body.Quantity),
	)
	if err != nil {
		a.log.Warn("error posting reserve", "error", err, "requestID", requestID)

		var appErr service.ApplicationError
		if errors.As(err, &appErr) {
			switch appErr.Status() {
			case http.StatusBadRequest:
				return gen.PostReserve400JSONResponse{Error: err.Error()}, nil
			case http.StatusNotFound:
				return gen.PostReserve404Response{}, nil
			default:
				return gen.PostReserve500Response{}, nil
			}
		}
	}

	return gen.PostReserve200JSONResponse{Message: "OK"}, nil
}
