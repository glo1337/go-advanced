package user

import (
	"errors"
	"net/http"
	"order-api/request"
	"order-api/response"
	"slices"

	"gorm.io/gorm"
)

type UserHandler struct {
	UserRepository *UserRepository
}

type UserHandlerDeps struct {
	UserRepository *UserRepository
}

var allowedCodes = []uint{1234, 3245, 0000}

func NewUserHandler(router *http.ServeMux, deps UserHandlerDeps) {
	handler := &UserHandler{
		UserRepository: deps.UserRepository,
	}
	router.HandleFunc("POST /auth/phone", handler.AuthByPhone())
	router.HandleFunc("POST /auth/verify", handler.VerifyAuth())
}

func (handler *UserHandler) AuthByPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[UserAuthByPhoneRequest](&w, r)
		if err != nil {
			return
		}

		user, err := handler.UserRepository.FindByPhone(body.Phone)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Ошибка", http.StatusBadRequest)
			return
		}

		if user == nil {
			user = NewUser(body.Phone)
			user, err = handler.UserRepository.Create(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			user.GenerateSessionId()
			_, err := handler.UserRepository.UpdateSessionId(user.ID, user.SessionId)
			if err != nil {
				http.Error(w, "Произошла ошибка", http.StatusInternalServerError)
				return
			}
		}

		authResponse := UserAuthByPhoneResponse{
			SessionId: user.SessionId,
		}

		response.Json(w, authResponse, 200)
	}
}

func (handler *UserHandler) VerifyAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[UserVerifyAuthRequest](&w, r)
		if err != nil {
			return
		}

		user, err := handler.UserRepository.FindBySessionId(body.SessionId)
		if err != nil {
			http.Error(w, "Ошибка", http.StatusBadRequest)
			return
		}

		if !isCodeAllowed(body.Code, allowedCodes) {
			http.Error(w, "Код недействителен", http.StatusInternalServerError)
			return
		}

		user.GenerateToken()
		_, err = handler.UserRepository.UpdateToken(user.ID, user.Token)
		if err != nil {
			http.Error(w, "Ошибка", http.StatusInternalServerError)
			return
		}

		verifyResponse := UserVerifyAuthResponse{
			Token: user.Token,
		}

		response.Json(w, verifyResponse, 200)
	}
}

func isCodeAllowed(code uint, allowed []uint) bool {
	return slices.Contains(allowed, code)
}
