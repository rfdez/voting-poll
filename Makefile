.PHONY: build watch clean test lint-dockerfile lint-go lint-yaml help
.DEFAULT_GOAL := help

DEPS := docker docker-compose
$(foreach bin,$(DEPS),\
		$(if $(shell command -v $(bin) 2> /dev/null),$(@info Found `$(bin)`),$(error Please install `$(bin)`)))

build: ## Build your project and run it
	@docker-compose up -d db mq
	@sleep 5
	@docker-compose up --build app && docker-compose down --remove-orphans --rmi local

watch: ## Run the code with docker-compose as development mode
	@docker-compose up -d db mq
	@sleep 5
	@docker-compose up --build dev-app && docker-compose down --remove-orphans --rmi local

clean: ## Remove build related file
	@rm -f ./coverage.txt

test: ## Run the project tests
	@docker run --rm -v $(shell pwd):/app -w /app golang:1.18 go test -short -race -covermode atomic -coverprofile coverage.txt ./... \
		&& go tool cover -func=coverage.txt

lint-dockerfile: ## Run the linter on the dockerfile
	@docker run --rm -i hadolint/hadolint:latest < Dockerfile

lint-go: ## Use golintci-lint on your project
	@docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run ./...

lint-yaml: ## Use yamllint on the yaml file of your projects
	@docker run --rm -it -v $(shell pwd):/data cytopia/yamllint:latest $(shell git ls-files '*.yml' '*.yaml')

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
