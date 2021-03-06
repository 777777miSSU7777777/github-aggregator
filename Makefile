.PHONY: check
check:
	gofmt -d ./$(find . -type f -name '*.go' -not -path "./vendor/*")
	bash build/gofmt_check.sh

	go list ./... | grep -v /vendor/ | xargs -n 1 golint

	# Exclude G104 because it fails on logging middleware.
	gosec -exclude=G104 ./...

	goimports -d ./$(find . -type f -name '*.go' -not -path "./vendor/*")
	bash build/goimports_check.sh

.PHONY: test
test:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...
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