.DEFAULT_GOAL := help

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

clean:        ## Clean project
	go version
	rm -rf target
	rm -rf vendor
	go mod tidy

format: 
	go fmt ./...

generate-mock: 
	go generate -v ./...

run:    ## Run application
	go run main.go

run-docker:
	docker run -it -p 8081:8081 -e APP_MONGO_URI=$(APP_MONGO_URI) chetannimbalkar98/vehicles-app

coverage:   ## Run code coverage
	go test ./... -coverpkg ./... -coverprofile coverage.out && cat coverage.out | grep -vE '(/app/|configs|/logger/|/generated/|/cloud/)' > coverage_filtered.out

docker-compose-up:
	docker-compose up 

docker-compose-down:
	docker-compose down

minikube: ## Start minkube
	minikube start
	

start-cluster: ## Deploy vehicles-app
	kubectl apply -f manifests/

port-forward: ## Expose a service
	kubectl port-forward svc/vehiclesapp 8081:8081

stop-cluster: ## Stop vehicles-app
	kubectl delete -f manifests/
	minikube stop 


test:        ## Run test suite
	go mod vendor
	go build ./...
	go test -cover -race ./...

windows:
	env GOOS=windows GOARCH=amd64 go build main.go