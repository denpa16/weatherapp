LOCAL_BIN:=$(CURDIR)/bin

ENV_NAME="weatherapp"

include .env

build:
	go build -o bin/weatherapp ./cmd/weatherapp/ 

run:
	go run ./cmd/weatherapp/
