.PHONY: dev
dev:
	@go run ./cmd/diploid/main.go

.PHONY: build
build:
	@go build -o ./.bin/diploid ./cmd/diploid/main.go

.PHONY: run
run:
	@./.bin/server

.PHONY: test
test:
	@go test -v ./...

.PHONY: db-up
db-up:
	@dbmate up

.PHONY: db-gen
db-gen:
	@sqlc generate
