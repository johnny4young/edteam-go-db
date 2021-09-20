package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/johnny4young/edteam-go-db/pkg/invoiceheader"
	"github.com/johnny4young/edteam-go-db/pkg/invoiceitem"
	"github.com/johnny4young/edteam-go-db/pkg/product"

	"github.com/johnny4young/edteam-go-db/storage"
)

func main() {
	fmt.Println("it's working")

	storage.NewPostgresDB()
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v", err)
	}

	// CREATE SQL
	// m := &product.Model{
	// 	Name:  "OOP GO",
	// 	Price: 50,
	// }
	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// fmt.Printf("%+v\n", m)

	// GET ALL
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}
	fmt.Println(ms)

	m, err := serviceProduct.GetByID(2)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println("there is not product found")
	case err != nil:
		log.Fatalf(("product.GetByID: %v"), err)
	default:
		fmt.Println(m)
	}

	storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)

	if err := serviceInvoiceHeader.Migrate(); err != nil {
		log.Fatalf("invoiceHeader.Migrate: %v", err)
	}

	storageInvoiceItem := storage.NewPsqlInvoiceItem((storage.Pool()))
	serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	if err := serviceInvoiceItem.Migrate(); err != nil {
		log.Fatalf("invoiceItem.Migrate: %v", err)
	}

}
