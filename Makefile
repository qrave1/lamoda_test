IMAGE="lamoda_test"
TAG=$(IMAGE)":latest"

.PHONY: init
init:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: migrate
migrate:
	@go run cmd/main.go migrate $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
build:
	@docker build -t $(TAG) .

.PHONY: dcup
dcup:
	@docker-compose up -d --build

.PHONY: dcdn
dcdn:
	@docker-compose down

.PHONY: swagger-gen
swagger_gen:
	@oapi-codegen \
	-generate fiber,types,strict-server,spec \
	-package gen -o internal/interface/http/gen/openapi_gen.go ./api/api.yaml

.PHONY: tests
tests:
	@go run tests/random_data/random_data.go truncate
	@go run tests/random_data/random_data.go insert
	@go test ./tests -v
	@go run tests/random_data/random_data.go truncate

.PHONY: insert
insert:
	@go run tests/random_data/random_data.go insert
