# The help comments are Based on https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Include shared utility definitions
include utils/colors.mk
include utils/common.mk

.DEFAULT_GOAL := help

.PHONY: buildall testall

buildall: ## run make in all sub directories
	@for dir in client/ server/; do \
		if [ -d "$$dir" ]; then \
			printf "\n=== Building $(C_CYAN)$$dir$(T_RESET)...\n"; \
			$(MAKE) --no-print-directory -C "$$dir" build || exit 1; \
			echo ""; \
		fi; \
	done

testall: ## run test in all sub directories
	@for dir in client/ server/; do \
		if [ -d "$$dir" ]; then \
			printf "\n=== Testing $(C_CYAN)$$dir$(T_RESET)...\n"; \
			$(MAKE) --no-print-directory -C "$$dir" test || exit 1; \
			echo ""; \
		fi; \
	done
