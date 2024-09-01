#!make
include .env.local

build:
	@go build -tags dev -o bin/github.com/yakubu-llc/jumaah-api cmd/app/main.go  

run: build
	@./bin/github.com/yakubu-llc/jumaah-api

install:
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download

air:
	@air -c .air.toml

docs: 
	@go run . openapi > openapi.yaml

db-status:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${DATABASE_URL} goose -dir=${GOOSE_MIGRATIONS_PATH} status

db-up:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${DATABASE_URL} goose -dir=${GOOSE_MIGRATIONS_PATH} up

db-up-one:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${DATABASE_URL} goose -dir=${GOOSE_MIGRATIONS_PATH} up-by-one 
	
db-down:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${DATABASE_URL} goose -dir=${GOOSE_MIGRATIONS_PATH} down

db-reset:
	@GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${DATABASE_URL} goose -dir=${GOOSE_MIGRATIONS_PATH} reset