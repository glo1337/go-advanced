package main

import (
	"fmt"
	"net/http"
	"validation-api/configs"
	"validation-api/internal"
	"validation-api/internal/verify"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	storage := &internal.HashStorage{
		Path: "hashstorage.json",
	}
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		EmailConfig: conf.EmailConfig,
		Storage:     *storage,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server started and listening on port 8081 ...")
	server.ListenAndServe()
}
