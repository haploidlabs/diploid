.PHONY: dev
dev:
	@go run ./cmd/server/main.go

.PHONY: build
build:
	@go build -o ./bin/server ./cmd/server/main.go

.PHONY: run
run:
	@./bin/server

.PHONY: test
test:
	@go test -v ./...