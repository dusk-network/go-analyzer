lint: ## Lint the files
	GOBIN=$(PWD)/bin go run scripts/build.go lint
