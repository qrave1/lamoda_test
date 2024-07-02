package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type reserveBody struct {
	ProductCodes []int `json:"product_codes"`
	WarehouseId  int   `json:"warehouse_id"`
	Quantity     int   `json:"quantity"`
}

func TestReserve(t *testing.T) {
	t.Parallel()

	ts := []struct {
		name     string
		body     reserveBody
		response string
		status   int
	}{
		{
			name: "Резерв на доступном складе",
			body: reserveBody{
				ProductCodes: []int{1234},
				WarehouseId:  1,
				Quantity:     15,
			},
			response: "{\"message\":\"OK\"}",
			status:   http.StatusOK,
		},
		{
			name: "Резерв на недоступном складе",
			body: reserveBody{
				ProductCodes: []int{1},
				WarehouseId:  2,
				Quantity:     1,
			},
			response: "{\"error\":\"warehouse not available\"}",
			status:   http.StatusBadRequest,
		},
		{
			name: "Резерв на несуществующем складе",
			body: reserveBody{
				ProductCodes: []int{1},
				WarehouseId:  1337,
				Quantity:     1,
			},
			response: "",
			status:   http.StatusNotFound,
		},
		{
			name: "Резерв несуществующего товара",
			body: reserveBody{
				ProductCodes: []int{223410},
				WarehouseId:  1,
				Quantity:     1,
			},
			response: "",
			status:   http.StatusNotFound,
		},
		{
			name:     "Резерв с пустым телом",
			body:     reserveBody{},
			response: "{\"error\":\"not valid body\"}",
			status:   http.StatusBadRequest,
		},
		{
			name: "Резерв с ошибкой валидации",
			body: reserveBody{
				ProductCodes: []int{-1},
				WarehouseId:  -1,
				Quantity:     1,
			},
			response: "{\"error\":\"not valid body\"}",
			status:   http.StatusBadRequest,
		},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			reqBody := new(bytes.Buffer)
			_ = json.NewEncoder(reqBody).Encode(tc.body)

			resp, err := http.Post(
				fmt.Sprintf("http://localhost:8080/reserve"),
				"application/json",
				reqBody,
			)
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
