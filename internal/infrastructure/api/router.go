package api

import (
	"net/http"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	//warehouseHandler := handlers.NewWarehouseHandler()
	mux.HandleFunc("POST /reserve", warehouseHandler.ReserveProducts)
	mux.HandleFunc("POST /release", warehouseHandler.ReleaseProducts)
	mux.HandleFunc("GET /inventory", warehouseHandler.GetInventory)

	return mux
}
