.PHONY: help install build serve deploy clean test

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install MkDocs and dependencies
	pip install -r requirements.txt

build: ## Build the MkDocs site
	mkdocs build

serve: ## Serve the documentation locally (with auto-reload)
	mkdocs serve

deploy: ## Deploy to GitHub Pages (requires push permissions)
	mkdocs gh-deploy

clean: ## Clean the build directory
	rm -rf site/
	rm -rf *.out
	rm -rf *.stdout
	rm -rf *.stderr

test: ## Test all runn examples
	go test ./...
