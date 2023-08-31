# Microservice API

This is a microservice API for managing products. It provides endpoints for listing, retrieving, creating, updating, and deleting products.

## Directory Structure


- `data`: Contains data structures and functions for managing products.
- `handlers`: Contains HTTP handlers for handling product-related operations.
- `main.go`: The main entry point of the microservice that sets up the server and routes.

## Packages

### data

This package provides data structures and functions for managing products.

- `Product`: Defines the structure of a product.
- `Products`: A slice of products.
- `GetProducts()`: Returns a list of all products.
- `GetProductByID(id int)`: Returns a product by its ID.
- `UpdateProduct(p Product)`: Updates an existing product.
- `AddProduct(p Product)`: Adds a new product.
- `DeleteProduct(id int)`: Deletes a product by its ID.
- `findIndexByProductID(id int)`: Finds the index of a product by its ID.
- `ToJSON(i interface{}, w io.Writer)`: Serializes data to JSON.
- `FromJSON(i interface{}, r io.Reader)`: Deserializes JSON data.

### handlers

This package contains HTTP handlers for product operations.

- `Products`: Struct that holds a logger instance.
- `NewProducts(l *log.Logger)`: Creates a new Products instance.
- `getProductID(r *http.Request)`: Extracts the product ID from a request.
- `ListAll(rw http.ResponseWriter, r *http.Request)`: Handler to list all products.
- `ListSingle(rw http.ResponseWriter, r *http.Request)`: Handler to get a single product by ID.
- `MiddlewareValidateProduct(next http.Handler) http.Handler`: Middleware for validating product data.
- `Create(rw http.ResponseWriter, r *http.Request)`: Handler to create a new product.
- `Update(rw http.ResponseWriter, r *http.Request)`: Handler to update an existing product.
- `Delete(rw http.ResponseWriter, r *http.Request)`: Handler to delete a product.

### main

This is the main entry point of the microservice.

- Sets up the server and routes using Gorilla Mux.
- Handles HTTP requests for product operations.
- Integrates with gRPC for communication.
- Gracefully shuts down the server on receiving signals.

## Usage

1. Make sure you have Go installed.
2. Clone this repository.
3. Run the microservice using `go run main.go`.
4. Access the API endpoints as described in the handlers.


## Dependencies

- `github.com/gorilla/mux`
- `github.com/gorilla/handlers`
