.PHONY: build

app ?= not_exist

build:
	go build -v -o ./$(app)/bin/$(app) $(app)/main.go