# This file contains common Makefile targets that can be shared across the project.

.PHONY: help view check-docker

help: ## show help contents
	@printf "$(C_BLUE)%-12s$(T_RESET) %s\n" "Target" "Description"
	@printf "%-12s %s\n" "------" "-----------"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "$(C_BLUE)%-12s$(T_RESET) %s\n", $$1, $$2}'

view: ## view folder structure
	@printf "📁 Directory structure:\n"
	@# Check if tree command is available, if not print an error message
	@sh -c 'command -v tree > /dev/null || (printf "$(T_BOLD)$(C_RED)tree is not installed$(T_RESET)\n" && exit 1)'
	@tree

check-docker: ## check if docker is installed
	@command -v docker >/dev/null 2>&1 || { echo >&2 "$(T_BOLD)$(C_RED)Docker is not installed. Please install Docker and try again.$(T_RESET)"; exit 1; }