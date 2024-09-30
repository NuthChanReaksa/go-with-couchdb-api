package http

import (
	"encoding/json"
	"github.com/NuthChanReaksa/go-with-couchdb-api/internal/entities"
	"github.com/NuthChanReaksa/go-with-couchdb-api/internal/usecase"
	"net/http"
)

type ProductHandler struct {
	Usecase *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{Usecase: uc}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.Usecase.Create(r.Context(), &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) BulkCreateOrUpdate(w http.ResponseWriter, r *http.Request) {
	var products []*entities.Product
	if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results, err := h.Usecase.BulkCreateOrUpdate(r.Context(), products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// Similarly, implement other CRUD handlers (Get, Update, Delete)
