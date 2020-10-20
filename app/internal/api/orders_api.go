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
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Handler ...
type Handler struct {
	Store *store.Store
}

// CreateOrder ...
func (h *Handler) CreateOrder(w http.ResponseWriter, req *http.Request) {
	var order model.Order
	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	err = h.Store.Create(&order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseStatus{
		Status: "created",
	})
}

// UpdateOrderReady ...
func (h *Handler) UpdateOrderReady(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	err = h.Store.Update("ready", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStatus{
		Status: "updated",
	})
}

// UpdateOrderComplete ...
func (h *Handler) UpdateOrderComplete(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	err = h.Store.Update("completed", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseStatus{
		Status: "updated",
	})
}

// GetOrdersAwait ...
func (h *Handler) GetOrdersAwait(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	listOfOrders, err := h.Store.Get("not ready")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listOfOrders)
}

// GetOrdersReady ...
func (h *Handler) GetOrdersReady(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	listOfOrders, err := h.Store.Get("ready")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseStatus{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listOfOrders)
}
