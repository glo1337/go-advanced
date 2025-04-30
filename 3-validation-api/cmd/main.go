package main

import (
	"fmt"
	"net/http"
	"validation-api/configs"
	"validation-api/internal/verify"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewVerifyHandler(router, verify.VerifyHandlerDeps{
		EmailConfig: conf.EmailConfig,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server started and listening on port 8081 ...")
	server.ListenAndServe()
}
