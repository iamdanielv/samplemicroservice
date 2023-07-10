# The help comments are Based on https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.DEFAULT_GOAL := help

.PHONY: help drun dbuild dstop container dstat dlog view
help: ## show help contents
	@printf "\033[36m%-10s\033[0m %s\n" "Target" "Description"
	@printf "%-10s %s\n" "------" "-----------"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

build: ## build the sample server
	@echo "Building the Sample server"
	@go build -o bin/sampleserver

dbuild: ## build the sample server using a docker container
	@echo "Building the Sample server in docker"
	@docker build -t samplemicroservice .

run: build ## build and run the sample
	@echo "Running the Sample server"
	@bin/sampleserver

drun: ## run the sample server in a docker container
	@echo "Running the Sample server in docker"
	@# the following command will remove a container if it exists
	@# and ignore any errors
	-@docker rm -f sampleserver
	@docker run -p 8080:8080 -d --name sampleserver samplemicroservice

dstop: ## stop the sample server docker container
	@echo "Stoping the Sample server running in docker"
	@docker stop sampleserver
	@docker rm -f sampleserver

container: dbuild drun ## build and run the container using docker
	@docker ps --filter 'name=sampleserver' --format "{{.Names}} {{.Status}}\n {{.Ports}}"

dstat: ## get the status of the sample server container
	@docker ps --filter 'name=sampleserver' --format "{{.Names}} {{.Status}}\n {{.Ports}}"

dlog: ## tail the logs of the sample server container
	@docker logs sampleserver -f

view: ## view folder structure
	@tree
