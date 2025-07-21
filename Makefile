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

test-basics: ## Test Chapter 01 examples
	go test -v -run TestBasics

test-scenario: ## Test Chapter 02 examples
	go test -v -run TestChapter02

test-chapter03: ## Test Chapter 03 examples
	go test -v -run TestChapter03

test-chapter04: ## Test Chapter 04 examples
	go test -v -run TestChapter04
