IMAGE="lamoda_test"
TAG=IMAGE+":latest"

.PHONY: run
run:
	@go run cmd/main.go

.PHONY: migrate
migrate:
	@go run cmd/main.go migrate $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
build:
	@docker build -t ($TAG) .

.PHONY: dcup
dcup:
	@docker-compose up -d

.PHONY: dcdn
dcdn:
	@docker-compose down
