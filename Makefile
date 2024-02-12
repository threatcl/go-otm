GO_CMD?=go 
GOFMT_FILES?=$$(find . -name '*.go')

default: help

fmt: ## Checks go formatting
	goimports -w $(GOFMT_FILES)

vet: ## Run go vet
	$(GO_CMD) vet ./pkg/otm

test: ## Run go test
	$(GO_CMD) test ./pkg/otm

testvet: vet test ## Run go vet and test

testcover: ## Run go test and go tool cover
	$(GO_CMD) test -coverprofile=cover.txt ./pkg/otm; go tool cover -html=cover.txt

help: ## Output make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help fmt vet test testvet testcover

