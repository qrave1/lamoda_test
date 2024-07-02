package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestInventory(t *testing.T) {
	t.Parallel()

	ts := []struct {
		name        string
		warehouseID string
		response    string
		status      int
	}{
		{
			name:        "Товары на доступном складе",
			warehouseID: "5",
			response:    "1",
			status:      http.StatusOK,
		},
		{
			name:        "Товары на недоступном складе",
			warehouseID: "2",
			response:    "{\"error\":\"warehouse not available\"}",
			status:      http.StatusBadRequest,
		},
		{
			name:        "Товары на несуществующем складе",
			warehouseID: "1337",
			response:    "",
			status:      http.StatusNotFound,
		},
		{
			name:        "Пустой warehouse_id",
			warehouseID: "",
			response:    "Query argument warehouse_id is required, but not found",
			status:      http.StatusInternalServerError,
		},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("http://localhost:8080/inventory?warehouse_id=%s", tc.warehouseID))
			if err != nil {
				t.Error(err)
			}

			if resp.StatusCode != tc.status {
				t.Errorf("expected status %d, got %d", tc.status, resp.StatusCode)
			}

			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Error(err)
			}

			if string(body) != tc.response {
				t.Errorf("expected %s, got %s", tc.response, string(body))
			}
		})
	}
}
