# simple-go-ddd
Simple go domain-driven design pattern

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Docker-compose](https://docs.docker.com/compose/) for running MongoDB.

## Setup

Create MongoDB container

```bash
docker-compose up
```

## Run the service

Install packages

```bash
go get ./...
```

Run app

```bash
go run main.go
```
