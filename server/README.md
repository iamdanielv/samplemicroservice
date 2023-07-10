# Sample Server

The sample server is implemented in GoLang

## Compiling the Sample Server

You have two options for compiling the sample server

### Compiling natively

In order to compile on your current machine, you can run

```shell
make build
```

it will call

```shell
go build -o bin/sampleserver
```

and compile your source.

>**Note:** You may have to call `go mod download` first

### Compiling inside of a Docker Container
In order to compile inside of a Docker Container, you can run

```shell
make dbuild
```

it will call

```shell
docker build -t samplemicroservice .
```

which will run the build process inside of a docker container. With this command, you don't have to install the golang compiler on your machine.


## Running the Sample

In order to run the sample server, you can run either:

### Native run

This will run a natively compiled executable

```shell
make run
```

### Docker run

This will run a containerized version of the server

```shell
make drun
```

Alternatively, you can run

```shell
make container
```

which will build and run the container for you in one step


### Get Status

A status endpoint allows us to check the current status of our microservice. This comes in handy when using an orchestrator such as Kubernetes or Docker-compose.

The sample exposes a status / health check end point at [http://localhost:8080/status](http://localhost:8080/status). If you navigate to that page using your browser, you should get something similar to `message:"OK"`.

You can also use `curl` with the command:

```shell
curl http://localhost:8080/status
```

and get output similar to:

```
{"message":"OK"}
```


### Get a webpage of all TODO items

Regardless of which option you chose to run (Native or Container), you should now be able to browse to
[http://localhost:8080](http://localhost:8080) and get a webpage with a list of the current TODO items.

### Get a JSON object of all the TODO items

The sample provides an API endpoint at [http://localhost:8080/todos](http://localhost:8080/todos) that exposes all the data as a JSON object. The object will look similar to:

```json
{
    "Title": "TODO List",
    "Todos": [{
        "id": 1,
        "item": "Install GO",
        "isDone": true
    }, {
        "id": 2,
        "item": "Create Microservice",
        "isDone": false
    }]
}
```

### Get a JSON object of a single TODO item

The sample provides an API endpoint at [http://localhost:8080/todo/1](http://localhost:8080/todo/1) that exposes a single todo item as a JSON object. The object will look similar to:

```json
{
  "id": 1,
  "item": "Install GO",
  "isDone": true
}
```

Links to try:

- [http://localhost:8080/todo/1](http://localhost:8080/todo/1)
- [http://localhost:8080/todo/2](http://localhost:8080/todo/2)
- [http://localhost:8080/todo/0](http://localhost:8080/todo/0)
- [http://localhost:8080/todo/asd](http://localhost:8080/todo/asd)
- [http://localhost:8080/todo/-100](http://localhost:8080/todo/-100)


## Exploring the code

You should check out the `Makefile` to see the different options available.

> **Note:** try running `make help` to get a list of options