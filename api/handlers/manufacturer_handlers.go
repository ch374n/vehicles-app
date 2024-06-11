// api/handlers/manufacturer_handlers.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/ch374n/vehicles-app/internal/models"
	"github.com/ch374n/vehicles-app/internal/repository"

	"github.com/gorilla/mux"
)

type ManufacturerHandlers struct {
	repo repository.ManufacturerRepo
}

func NewManufacturerHandlers(r *repository.ManufacturerRepo) *ManufacturerHandlers {
	return &ManufacturerHandlers{repo: *r}
}

func (h *ManufacturerHandlers) GetManufacturers(w http.ResponseWriter, r *http.Request) {
	manufacturers, err := h.repo.GetAllManufacturers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manufacturers)
}

func (h *ManufacturerHandlers) GetManufacturer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	manufacturer, err := h.repo.GetManufacturer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manufacturer)
}

func (h *ManufacturerHandlers) CreateManufacturer(w http.ResponseWriter, r *http.Request) {
	var manufacturer models.Manufacturer
	err := json.NewDecoder(r.Body).Decode(&manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.CreateManufacturer(manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ManufacturerHandlers) UpdateManufacturer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var manufacturer models.Manufacturer
	err = json.NewDecoder(r.Body).Decode(&manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.UpdateManufacturer(id, manufacturer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ManufacturerHandlers) DeleteManufacturer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteManufacturer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}