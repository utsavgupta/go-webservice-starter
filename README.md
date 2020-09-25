# Go Web Service Starter Project

[![Build Status](https://travis-ci.org/utsavgupta/go-webservice-starter.svg?branch=master)](https://travis-ci.org/utsavgupta/go-webservice-starter)

The project provides a basic setup for writing web services in Go. 

It uses the following external dependencies.

- [Zap](https://github.com/uber-go/zap)
- [Http Router](https://github.com/JulienSchmidt/httprouter)

## Running the Application

The project can be run either as a native application on your system or as a container using the provided Docker file.

### Running the Application Natively

```bash
$ export APP_STAGE=local
$ export APP_PORT=8080
$ go run .
```

### Running the Application as a Container

```bash
$ docker build . -t go-webservice-starter:latest
$ docker run -d -p 8080:8080 -e APP_STAGE=local -e APP_PORT=8080 go-webservice-starter:latest
```
