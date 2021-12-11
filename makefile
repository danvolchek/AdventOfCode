.PHONY: all clean
all: readme template

readme: $(shell find cmd -type f)
	go build -o readme cmd/readme/main.go

template: $(shell find cmd -type f)
	go build -o template cmd/template/main.go

clean:
	rm readme
	rm template
