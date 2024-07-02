package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type releaseBody struct {
	ProductCodes []int `json:"product_codes"`
	WarehouseId  int   `json:"warehouse_id"`
}

func TestRelease(t *testing.T) {
	t.Parallel()

	ts := []struct {
		name     string
		body     releaseBody
		response string
		status   int
	}{
		{
			name: "Освобождение товаров на доступном складе",
			body: releaseBody{
				ProductCodes: []int{1234},
				WarehouseId:  3,
			},
			response: "{\"message\":\"OK\"}",
			status:   http.StatusOK,
		},
		{
			name: "Освобождение товаров на несуществующем складе",
			body: releaseBody{
				ProductCodes: []int{1},
				WarehouseId:  1337,
			},
			response: "",
			status:   http.StatusNotFound,
		},
		{
			name: "Освобождение несуществующего товара",
			body: releaseBody{
				ProductCodes: []int{223410},
				WarehouseId:  1,
			},
			response: "",
			status:   http.StatusNotFound,
		},
		{
			name:     "Освобождение товаров с пустым телом",
			body:     releaseBody{},
			response: "{\"error\":\"not valid body\"}",
			status:   http.StatusBadRequest,
		},
		{
			name: "Освобождение товаров с ошибкой валидации",
			body: releaseBody{
				ProductCodes: []int{-1},
				WarehouseId:  -1,
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
				fmt.Sprintf("http://localhost:8080/release"),
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
