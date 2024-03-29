#!/usr/bin/env bash

go fmt ./...
go mod tidy
go build ./...
go test -cover -race -count=1 -timeout=30s ./...