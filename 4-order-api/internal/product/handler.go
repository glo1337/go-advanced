package product

import (
	"net/http"
	"order-api/request"
	"order-api/response"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductRepository *ProductRepository
}

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("POST /product", handler.Create())
	router.Handle("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}

		product := NewProduct(body.Name)

		createdProduct, err := handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, createdProduct, 201)
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		product, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		response.Json(w, product, 201)
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = handler.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = handler.ProductRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Json(w, nil, 200)
	}
}
