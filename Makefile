.PHONY: check
check:
	gofmt -d ./$(find . -type f -name '*.go' -not -path "./vendor/*")
	bash build/gofmt_check.sh

	go list ./... | grep -v /vendor/ | xargs -n 1 golint

	gosec ./... $(find . -type f -name '*.go' -not -path "./vendor/*")

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