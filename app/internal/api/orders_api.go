package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kingofmidas/gfc-api/internal/model"
	"github.com/kingofmidas/gfc-api/internal/store"
)

// ResponseStatus ...
type responseStatus struct {
	Status string `json:"status"`
}

// Handler ...
type Handler struct {
	Store *store.Store
}

// CreateOrder ...
func (h *Handler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	var order model.Order
	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w)
	}

	err = h.Store.Create(&order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseStatus{"created"})
}

// UpdateOrderReady ...
func (h *Handler) UpdateOrderReady(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	err := h.Store.Update("ready", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}

	w.Header().Add("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStatus{"updated"})
}

// UpdateOrderComplete ...
func (h *Handler) UpdateOrderComplete(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])

	err := h.Store.Update("completed", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}

	w.Header().Add("Content-Type", "applications/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStatus{"updated"})
}

// GetOrdersAwait ...
func (h *Handler) GetOrdersAwait(w http.ResponseWriter, req *http.Request) {
	listOfOrders, err := h.Store.Get("not ready")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listOfOrders)
}

// GetOrdersReady ...
func (h *Handler) GetOrdersReady(w http.ResponseWriter, req *http.Request) {
	listOfOrders, err := h.Store.Get("ready")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listOfOrders)
}
