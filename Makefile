build:
	@go build -o bin/api cmd/api/main.go
run: build
	@./bin/api
migrate-create:
	@migrate create -ext sql -dir internal/migrations $(filter-out $@,$(MAKECMDGOALS))
migrate-up:
	@migrate -database $(GO_REST_POSTGRES_URL) -path internal/migrations up
migrate-down:
	@migrate -database $(GO_REST_POSTGRES_URL) -path internal/migrations down