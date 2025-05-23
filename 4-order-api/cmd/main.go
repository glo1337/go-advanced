package main

import (
	"fmt"
	"net/http"
	"order-api/configs"
	"order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(*conf)
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server started on port 8081 ...")
	server.ListenAndServe()
}
