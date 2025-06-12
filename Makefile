# The help comments are Based on https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.DEFAULT_GOAL := help

.PHONY: help view buildall
help: ## show help contents
	@printf "\033[36m%-10s\033[0m %s\n" "Target" "Description"
	@printf "%-10s %s\n" "------" "-----------"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'


view: ## view folder structure
	@sh -c 'command -v tree > /dev/null || (printf "\033[1;31mtree is not installed\033[0m\n" && exit 1)&& tree -a'


buildall: ## run make in all sub directories
	@for dir in client/ server/; do \
		if [ -d "$$dir" ]; then \
			printf "\n=== Building \033[36m$$dir \033[0m===\n"; \
			$(MAKE) --no-print-directory -C "$$dir" build || exit 1; \
		fi; \
	done