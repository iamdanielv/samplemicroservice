# The help comments are Based on https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.DEFAULT_GOAL := help

.PHONY: check-docker help build dbuild run drun dstop container dstat dlog view test

check-docker:
	@command -v docker >/dev/null 2>&1 || { echo >&2 "\033[1;31mDocker is not installed. Please install Docker and try again.\033[0m"; exit 1; }

help: ## show help contents
	@printf "\033[36m%-10s\033[0m %s\n" "Target" "Description"
	@printf "%-10s %s\n" "------" "-----------"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: ## build the sample server
	@echo "Building the Sample server"
	@go build -o bin/sampleserver

dbuild: check-docker ## build the sample server using a Docker container
	@echo "Building the Sample server in Docker"
	@docker build -t samplemicroservice .

run: build ## build and run the sample
	@echo "Running the Sample server"
	@bin/sampleserver

drun: check-docker ## run the sample server in a Docker container
	@echo "Running the Sample server in Docker"
	@# the following command will remove a container if it exists
	@# and ignore any errors
	-@docker rm -f sampleserver
	@docker run -p 8080:8080 -d --name sampleserver samplemicroservice

dstop: check-docker ## stop the sample server Docker container
	@echo "Stopping the Sample server running in Docker"
	@docker stop sampleserver
	@docker rm -f sampleserver

container: check-docker dbuild drun ## build and run the container using Docker
	@docker ps --filter 'name=sampleserver' --format "{{.Names}} {{.Status}}\n {{.Ports}}"

dstat: check-docker ## get the status of the sample server container
	@docker ps --filter 'name=sampleserver' --format "{{.Names}} {{.Status}}\n {{.Ports}}"

dlog: check-docker ## tail the logs of the sample server container
	@docker logs sampleserver -f

view: ## view folder structure
	@sh -c 'command -v tree > /dev/null || (printf "\033[1;31mtree is not installed\033[0m\n" && exit 1)&& tree -a'

test: ## run Go tests
	@printf "Running \033[33mtests...\033[0m\n"
	@go test ./...
