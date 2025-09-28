# The help comments are Based on https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

# Include shared color definitions
include ../utils/colors.mk
include ../utils/common.mk

.DEFAULT_GOAL := help

.PHONY: build dbuild run drun dstop container dstat dlog test

build: ## build the sample server
	@echo "🏗️  Building the Sample server..."
	@go build -o bin/sampleserver

dbuild: check-docker  ## build the sample server using a Docker container
	@printf "🐳 Building the Sample server in Docker...\n"
	@docker build -t samplemicroservice .

run: build ## build and run the sample
	@printf "🚀 Running the Sample server...\n"
	@bin/sampleserver

drun: check-docker ## run the sample server in a Docker container
	@printf "🐳🚀 Running the Sample server in Docker...\n"
	@# the following command will remove a container if it exists
	@# and ignore any errors
	-@docker rm -f sampleserver
	@docker run -p 8080:8080 -d --name sampleserver samplemicroservice

dstop: check-docker ## stop the sample server Docker container
	@printf "🛑 Stopping the Sample server in Docker...\n"
	@docker stop sampleserver
	@docker rm -f sampleserver

container: check-docker dbuild drun ## build and run the container using Docker
	@printf "🐳 Sample Server container status:\n"
	@docker ps --filter 'name=sampleserver' --format "  $(C_GREEN){{.Names}}$(T_RESET): {{.Status}} (Ports: {{.Ports}})\n" || printf "  Container not found.\n"


dstat: check-docker ## get the status of the sample server container
	@docker ps --filter 'name=sampleserver' --format "{{.Names}} {{.Status}}\n {{.Ports}}\n"

dlog: check-docker ## tail the logs of the sample server container
	@docker logs sampleserver -f

test: ## run Go tests
	@printf "Running $(C_YELLOW)tests...$(T_RESET)\n"
	@go test ./...
