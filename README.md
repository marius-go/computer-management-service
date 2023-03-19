# Computer Management Service

A small service managing computers of a company. It allows a system administrator to keep track of the computers issued by the company.

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

In order build the service on you local pc without docker you need to [install go](https://go.dev/doc/install). If you want to change the REST API you need to install [go-swagger](#code-generation-via-go-swagger) as well.

To build the service execute

```sh
go build -o computer-management-service gen/controller/rest/cmd/computer-management-server/main.go
```

Then you can run the service with 
```sh
./computer-management-service --port 5000 # run with flag --help for further command line options
```

You can then access the service as in the docker setup.

In doubt see the steps in the [Dockerfile](Dockerfile) (you can ignore the `COPY` commands as the files are already in their correct location).

### Run tests

Execute 

```sh
go test ./...
```

## Code generation via go-swagger

This project uses [go-swagger](https://github.com/go-swagger/go-swagger) to generate server code from the [OpenApi 2.0 specs](api/v1/computer-management.yaml).  Install according to the [official instructions](https://goswagger.io/install.html).

Afterwards generate code with 

```sh
swagger generate server -t gen/controller/rest -f api/v1/computer-management.yaml
```

This generates the required server code for the REST API. The glue code to connect the REST API calls to the core needs to be added to [configure_computer_management.go](gen/controller/rest/restapi/configure_computer_management.go).

Note that this file is not overwritten and needs to be deleted if you need to update it due to api changes. Then simply reapply your changes afterwards.
