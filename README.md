# Computer Management Service

A small service managing computers of a company.

## Prerequisites

Needed tools:
    - [docker](https://docs.docker.com/engine/install/)
    - [docker compose v2](https://docs.docker.com/compose/install/)
    
If you are having only compose v1 installed, just replace in all commands `docker compose` with `docker-compose`.

## Quickstart

To run all the required components execute

```sh
docker compose up --build --abort-on-container-exit
```
You can then access the api docs at http://localhost:5000/api/v1/docs.

## Local development

To build the service on you local pc without docker you need to [install go](https://go.dev/doc/install).

For further required setup see the steps in the [Dockerfile](Dockerfile) (you can ignore the `COPY` commands as the files are already in their correct location).

## Code generation via go-swagger

This project uses [go-swagger](https://github.com/go-swagger/go-swagger) to generate server code from the [OpenApi 2.0 specs](api/v1/computer-management.yaml).
