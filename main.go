package main

import (
	"fmt"
	"log"

	"github.com/johnny4young/edteam-go-db/pkg/product"

	"github.com/johnny4young/edteam-go-db/pkg/storage"
)

func main() {
	fmt.Println("it's working")

	storage.NewPostgresDB()
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

}
