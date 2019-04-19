hello-banner-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

hello-banner-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

mock:
	mockgen -source=modules/banner/query/query.go -destination=modules/banner/query/mocks/banner_query_mock.go -package=mocks
	mockgen -source=modules/banner/usecase/query.go -destination=modules/banner/usecase/mocks/banner_query_mock.go -package=mocks

test: SHELL := /bin/bash

test:
	go test -race -short ./... | grep -v mocks; exit $${PIPESTATUS[0]}
