# Makefile

SHELL := /bin/bash

compile:
	go build -o build/server ./cmd/server
vet: fmt
	go vet ./...
fmt:
	go fmt ./...