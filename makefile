.PHONY: all readme template

all: template readme

template:
	go run cmd/template/main.go

readme:
	go run cmd/readme/main.go

test:
	go test ./cmd/*
