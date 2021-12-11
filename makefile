.PHONY: all both readme template
all: both

both: template readme

readme:
	go run cmd/readme/main.go

template:
	go run cmd/template/main.go
