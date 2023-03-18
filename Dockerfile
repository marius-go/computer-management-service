FROM golang:1.20 as builder

# install go-swagger, see also: https://goswagger.io/install.html 
RUN curl --output /usr/local/bin/swagger -L https://github.com/go-swagger/go-swagger/releases/download/v0.30.4/swagger_linux_amd64
RUN chmod +x /usr/local/bin/swagger


WORKDIR /computer-management-service

# install dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# generate rest api code
COPY api api
RUN mkdir -p gen/controller/rest
RUN swagger generate server -t gen/controller/rest -f api/v1/computer-management.yaml

# build service
COPY gen/controller/rest/restapi/configure_computer_management.go gen/controller/rest/restapi/configure_computer_management.go
COPY internal internal
RUN CGO_ENABLED=0 go build -o computer-management-service gen/controller/rest/cmd/computer-management-server/main.go


FROM scratch
COPY --from=builder /computer-management-service/computer-management-service /bin/computer-management-service

ENTRYPOINT ["/bin/computer-management-service"]
CMD ["--host", "0.0.0.0", "--port", "5000"]
# default port of the service
EXPOSE 5000
