.PHONY: all readme template
all: readme template

readme: template $(shell find cmd -type f)
	go run cmd/readme/main.go

template: $(shell find cmd -type f)
	go run cmd/template/main.go
