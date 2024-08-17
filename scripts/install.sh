#!/bin/sh

go install github.com/automation-co/husky@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1
go install github.com/securego/gosec/v2/cmd/gosec@latest
go get
