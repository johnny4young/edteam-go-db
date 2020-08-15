package main

import (
	"fmt"

	"github.com/johnny4young/edteam-go-db/pkg/storage"
)

func main() {
	fmt.Println("it's working")

	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
	storage.NewPostgresDB()
}
