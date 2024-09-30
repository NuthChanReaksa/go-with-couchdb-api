package main

import (
	"context"
	"log"
	"net/http" // standard http package

	customHTTP "cmd/internal/http" // renamed internal http package to avoid conflict
	"cmd/internal/infrastructure"
	"cmd/internal/repository"
	"cmd/internal/usecase"
)

func main() {
	ctx := context.Background() // Reintroducing the context
	dsn := "http://localhost:5984/"
	dbName := "products"

	// Connect to CouchDB
	client, err := infrastructure.NewCouchDBConnection(ctx, dsn)
	if err != nil {
		log.Fatalf("Failed to connect to CouchDB: %v", err)
	}

	// Open the database
	db, err := infrastructure.OpenDatabase(client, dbName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Set up repository, usecase, and handler
	repo := repository.NewCouchDBProductRepository(db)
	productUsecase := usecase.NewProductUsecase(repo)
	productHandler := customHTTP.NewProductHandler(productUsecase) // Use customHTTP to avoid conflict

	// Define the routes and handlers
	http.HandleFunc("/products", productHandler.Create)
	http.HandleFunc("/products/bulk", productHandler.BulkCreateOrUpdate)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
