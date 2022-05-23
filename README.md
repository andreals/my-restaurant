# My Restaurant
## _Code Challenge: My Restaurant_

My Restaurant is a "code challenge" offered by Checkmarx as a step in a Software Engineering selection process.

The CLI application is responsible for creating/loading a restaurant, placing orders and preparing table orders.

## Stack

My Restaurant uses as a technology stack:
- Golang
- Docker (optional)
- Make (optional)

## Installation

The application needs an environment with [Golang](https://go.dev/doc/install) 1.17+.

Install the dependency and run the command:
```sh
go run cmd/cli/main.go
```

## Tests

To run the application tests, just run the command in the project root:
```sh
go test ./...
```

## Docker

My Restaurant is very easy to install and run in a Docker container.
To do so, follow the steps below to succeed:

To build the container:
```sh
docker build -t restaurant . -f ./build/Dockerfile
```

To run the container:

```sh
docker run -it restaurant
```

## Makefile

An alternative to run it is using the Makefile that is in the root of the project.

Existing commands:
```sh
make run
```
This command will run the My Restaurant application in Go on your local machine.

```sh
make build
```
This command will generate a My Restaurant application binary to run anywhere.

```sh
make test
```
This command will run all existing tests in the My Restaurant application to validate that everything is OK with its execution.

```sh
make docker-build
```
This command will build a new docker image with the environment ready to run the My Restaurant application.

```sh
make docker-run
```
This command will run the generated docker image to run the My Restaurant application.

## Notes

As it is a language in which there is no architectural "rule", I used some community premises and adhered to some good market practices, which I have been improving since 2018 when I had my first contact with the language in a monolith.

## License

MIT