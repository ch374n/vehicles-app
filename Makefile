.DEFAULT_GOAL := help


clean:					## Clean project
	go version
	rm -rf target
	rm -rf vendor
	go mod tidy

generate-mock: 
	go generate -v ./...

run:
	go run main.go

coverage:
	go test ./... -coverpkg ./... -coverprofile coverage.out && cat coverage.out | grep -vE "(/mocks/|testenv|/cmd/|/generated/|/cloud/)" > coverage_filtered.out && go tool cover -html coverage_filtered.out

test:					## Run test suite
	go mod vendor
	go build ./...
	go test -cover -race ./...

windows: 
	env GOOS=windows GOARCH=amd64 go build main.go

help:					## Display available targets and documentation
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


