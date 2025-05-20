package verify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"validation-api/configs"
	"validation-api/internal"

	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	EmailConfig configs.EmailConfig
	Storage     internal.HashStorage
}

type VerifyHandler struct {
	EmailConfig configs.EmailConfig
	Storage     internal.HashStorage
}

type SendRequest struct {
	Email string `json:"email"`
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		EmailConfig: deps.EmailConfig,
		Storage:     deps.Storage,
	}
	router.HandleFunc("POST /send", handler.send())
	router.HandleFunc("GET /verify/{hash}", handler.verify())
}

func (handler *VerifyHandler) send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestData SendRequest
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil || requestData.Email == "" {
			http.Error(w, "Invalid or missing email", http.StatusBadRequest)
			return
		}

		hash, err := internal.GenerateHash()
		if err != nil {
			http.Error(w, "Failed to generate hash", http.StatusInternalServerError)
			return
		}

		items, err := handler.Storage.ReadItems()
		if err != nil {
			http.Error(w, "Failed to read items", http.StatusInternalServerError)
			return
		}

		items = append(items, internal.StorageItem{
			Email: requestData.Email,
			Hash:  hash,
		})
		err = handler.Storage.WriteItems(items)
		if err != nil {
			http.Error(w, "Failed to write items", http.StatusInternalServerError)
			return
		}

		link := fmt.Sprintf("http://localhost:8081/verify/%s", hash)

		e := email.NewEmail()
		e.From = "Sender <" + handler.EmailConfig.Email + ">"
		e.To = []string{requestData.Email}
		e.Subject = "Verification link From Go Api"
		e.Text = []byte("Click this link to verify: " + link)
		err = e.Send("smtp.mail.ru:587", smtp.PlainAuth("", handler.EmailConfig.Email, handler.EmailConfig.Password, "smtp.mail.ru"))
		if err != nil {
			http.Error(w, "Failed to send email", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Email sent with hash: %s\n", link)
	}
}

func (handler *VerifyHandler) verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if len(hash) == 0 {
			http.Error(w, "Hash not valid", http.StatusBadRequest)
			return
		}

		items, err := handler.Storage.ReadItems()
		if err != nil {
			http.Error(w, "Failed to read items", http.StatusInternalServerError)
			return
		}

		var updatedItems []internal.StorageItem
		found := false

		for _, item := range items {
			if item.Hash == hash {
				found = true
				break
			}
			updatedItems = append(updatedItems, item)
		}

		if found {
			err := handler.Storage.WriteItems(updatedItems)
			if err != nil {
				http.Error(w, "Failed to update items", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Verified successfully"))
			return
		}

		http.Error(w, "Hash not found", http.StatusNotFound)
	}
}
