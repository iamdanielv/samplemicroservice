# Sample Microservice Implementation with GoLang and Docker

This repository contains a sample microservice implementation using GoLang and Docker. It includes both server and client portions, though the client is not yet implemented.

## Code Layout

The code is organized into two main directories:

- `server`: Contains the server logic.
- `client`: Placeholder for client-side code (currently not implemented).

## Makefile Targets

This command will provide a list of available Make commands to get you started:

```sh
make help
```

### Server Directory (`server/Makefile`)

- **help**: Show help contents.
- **check-docker**: Check if Docker is installed.
- **build**: Build the sample server.
- **dbuild**: Build the sample server using a Docker container.
- **run**: Build and run the sample.
- **drun**: Run the sample server in a Docker container.
- **dstop**: Stop the sample server Docker container.
- **container**: Build and run the container using Docker.
- **dstat**: Get the status of the sample server container.
- **dlog**: Tail the logs of the sample server container.
- **view**: View folder structure.
- **test**: Run Go tests.

### Client Directory (`client/Makefile`)

- **help**: Show help contents.
- **check-docker**: Check if Docker is installed.
- **build**: Build the sample client (not yet implemented).
- **test**: Run tests for the sample client (not yet implemented).

## Folder structure

To view the folder structure, you can use:

```sh
make view
```

## Next Steps

**Have fun exploring!** Your journey into the world of microservices with GoLang and Docker is just beginning!

Let me know if you have any questions. I can be reached at [@IAmDanielV](https://twitter.com/IAmDanielV) or [@iamdanielv.bsky.social](https://bsky.app/profile/iamdanielv.bsky.social).
