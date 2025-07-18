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

test: ## Test all runn examples
	@echo "Testing runn examples..."
	@for dir in examples/*/; do \
		echo "Testing $$dir"; \
		for file in $$dir*.yml; do \
			echo "  Running: $$file"; \
			runn run "$$file" || exit 1; \
		done; \
	done
	@echo "All tests passed!"

test-chapter01: ## Test Chapter 01 examples
	@echo "Testing Chapter 01 examples..."
	@cd examples/chapter01 && for file in *.yml; do \
		echo "Running: $$file"; \
		runn run "$$file" || exit 1; \
	done

test-chapter02: ## Test Chapter 02 examples
	@echo "Testing Chapter 02 examples..."
	@cd examples/chapter02 && for file in *.yml; do \
		echo "Running: $$file"; \
		runn run "$$file" || exit 1; \
	done

test-chapter03: ## Test Chapter 03 examples
	@echo "Testing Chapter 03 examples..."
	@cd examples/chapter03 && for file in *.yml; do \
		echo "Running: $$file"; \
		runn run "$$file" || exit 1; \
	done

test-chapter04: ## Test Chapter 04 examples
	@echo "Testing Chapter 04 examples..."
	@cd examples/chapter04 && for file in *.yml; do \
		echo "Running: $$file"; \
		runn run "$$file" || exit 1; \
	done