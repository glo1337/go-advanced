package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type RandomHandler struct{}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/", handler.ThrowDice())
}

func (handler *RandomHandler) ThrowDice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, rand.Intn(7))
	}
}
