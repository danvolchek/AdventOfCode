.PHONY: all readme template

all: template readme

readme:
	go run cmd/readme/main.go

template:
	go run cmd/template/main.go
