.PHONY: check
check:
	gofmt -d ./$(find . -type f -name '*.go' -not -path "./vendor/*")
	bash build/gofmt_check.sh

	go list ./... | grep -v /vendor/ | xargs -n 1 golint

	gosec ./...

	goimports -d ./$(find . -type f -name '*.go' -not -path "./vendor/*")
	bash build/goimports_check.sh

.PHONY: test
test:
	go test -race -coverprofile=coverage.out.tmp -covermode=atomic ./...
	cat coverage.out.tmp | grep -v "internal/api" | grep -v "pkg/query/datasource" > coverage.out
	go tool cover -html=coverage.out

.PHONY: build
build: check test
	go build cmd/web-app/main.go

.PHONY: run
run: build
	./main

.PHONY: fmt
fmt:
	gofmt -w ./
	goimports -w ./