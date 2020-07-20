# Go API

A simple Go API following concepts of Domain Driven Design for educational purposes.

Open API specification can be found [here](docs/reference/openapi.v1.yaml)

## Installation

Clone the repo

```bash
$ git clone git@github.com:JustSteveKing/go-api.git
```

```bash
$ cd go-api
```

## Running

Once installed, simply spin up the docker container:

```bash
$ docker-compose up -d --build
```

This will give you a single endpoint for now under: `http://localhost:8080/ping` which is a healthcheck URL, and will return the following:

```json
{
    "message": "Service Online"
}
```
