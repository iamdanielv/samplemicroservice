# The help comments are Based on https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.DEFAULT_GOAL := help

.PHONY: help view buildall testall

help: ## show help contents
	@printf "$(C_BLUE)%-12s$(T_RESET) %s\n" "Target" "Description"
	@printf "%-12s %s\n" "------" "-----------"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

view: ## view folder structure
	@printf "📁 Directory structure:\n"
	@# Check if tree command is available, if not print an error message
	@sh -c 'command -v tree > /dev/null || (printf "\033[1;31mtree is not installed\033[0m\n" && exit 1)'
	@tree

build: ## run make in all sub directories
	@for dir in client/ server/; do \
		if [ -d "$$dir" ]; then \
			printf "\n=== Building \033[36m$$dir \033[0m...\n"; \
			$(MAKE) --no-print-directory -C "$$dir" build || exit 1; \
			echo ""; \
		fi; \
	done


test: ## run test in all sub directories
	@for dir in client/ server/; do \
		if [ -d "$$dir" ]; then \
			printf "\n=== Testing \033[36m$$dir \033[0m...\n"; \
			$(MAKE) --no-print-directory -C "$$dir" test || exit 1; \
			echo ""; \
		fi; \
	done
