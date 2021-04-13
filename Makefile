.PHONY: clean

build:
	go build -o dist/parser cmd/parser/main.go

clean:
	rm -rf dist/*
