package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/internal/product"
	"order-api/internal/user"
	"order-api/middleware"
	"order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(*conf)
	router := http.NewServeMux()

	// Reposotories
	productRepo := product.NewProductRepository(db)
	userRepo := user.NewUserRepository(db)

	// Handler
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: productRepo,
	})
	user.NewUserHandler(router, user.UserHandlerDeps{
		UserRepository: userRepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: middleware.Logging(router),
	}

	fmt.Println("Server started on port 8081 ...")
	server.ListenAndServe()
}
