# Microservice API and Server

This repository contains a microservice API for managing products and a server for handling currency conversion rate requests.

## Products Microservice

The `data` and `handlers` packages provide the functionality to manage products via HTTP requests. The `main.go` sets up the server and routes for handling these operations.

### Usage

1. Clone this repository.
2. Navigate to the `data` directory and run `go run .` to start the products microservice.
3. Access the API endpoints as described in the handlers.

## Currency Server

The `protos` directory contains the protocol buffer definitions for currency conversion rate requests. The `server` package implements the gRPC server to handle these requests.

### Usage

1. Navigate to the `server` directory and run `go run .` to start the currency conversion rate server.
2. The server listens for gRPC requests and provides rate conversion and subscription functionalities.

## Dependencies

- `google.golang.org/grpc`
- `google.golang.org/grpc/credentials`
- `github.com/gorilla/mux`
- `github.com/gorilla/handlers`
