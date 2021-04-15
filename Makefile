.PHONY: clean

all: build-parser build-tdef-tree build-find

build-parser:
	go build -o dist/parser cmd/parser/main.go

build-tdef-tree:
	go build -o dist/tdef-tree cmd/tree/main.go

build-find:
	go build -o dist/find cmd/find/main.go

clean:
	rm -rf dist/*
