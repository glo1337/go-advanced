package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/internal/product"
	"order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(*conf)
	router := http.NewServeMux()

	// Reposotories
	productRepo := product.NewProductRepository(db)

	// Handler
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server started on port 8081 ...")
	server.ListenAndServe()
}
