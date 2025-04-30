package verify

import (
	"fmt"
	"net/http"
	"net/smtp"
	"validation-api/configs"

	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	EmailConfig configs.EmailConfig
}

type VerifyHandler struct {
	EmailConfig configs.EmailConfig
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		EmailConfig: deps.EmailConfig,
	}
	router.HandleFunc("POST /send", handler.send())
	router.HandleFunc("GET /verify/{hash}", handler.verify())
}

func (handler *VerifyHandler) send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("send")

		e := email.NewEmail()
		// config
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	}
}

func (handler *VerifyHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("verify")
	}
}
