package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/qrave1/lamoda_test/internal/domain/repository"
	"github.com/qrave1/lamoda_test/internal/domain/service"
)

const (
	InternalServerError = "internal server error"
)

type Handler struct {
	service *service.WarehouseService
}

func NewWarehouseHandler(db *sql.DB) *Handler {
	productRepo := repository.NewProductRepository(db)
	reservationRepo := repository.NewReservationRepository(db)
	warehouseService := service.NewWarehouseService(productRepo, reservationRepo)
	return &Handler{warehouseService}
}

func (h *Handler) ReserveProducts(w http.ResponseWriter, r *http.Request) {
	var request ProductsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.ReserveProducts(request.ProductCodes, request.WarehouseID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := json.NewEncoder(w).Encode(map[string]string{"message": "Products reserved successfully"})
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ReleaseProducts(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ProductCodes []string `json:"product_codes"`
		WarehouseID  uint     `json:"warehouse_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.ReleaseProducts(request.ProductCodes, request.WarehouseID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Products released successfully"})
}

func (h *Handler) GetInventory(w http.ResponseWriter, r *http.Request) {
	warehouseID := r.URL.Query().Get("warehouse_id")
	if warehouseID == "" {
		http.Error(w, "warehouse_id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(warehouseID)
	if err != nil {
		http.Error(w, "invalid warehouse_id", http.StatusBadRequest)
		return
	}

	products, err := h.service.GetInventory(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"inventory": products})
}
