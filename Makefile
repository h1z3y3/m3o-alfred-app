.PHONY: build

app ?= not_exist

build:
	go build -v -o ./$(app)/bin/gif-search $(app)/main.go