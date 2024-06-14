package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ch374n/vehicles-app/internal/models"
	"github.com/ch374n/vehicles-app/internal/repository"
	"github.com/ch374n/vehicles-app/logger"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

type ManufacturerHandlers struct {
	repo       repository.ManufacturerRepo
	collection *mongo.Collection
}

func NewManufacturerHandlers(r *repository.ManufacturerRepo, collection *mongo.Collection) *ManufacturerHandlers {
	return &ManufacturerHandlers{
		repo:       *r,
		collection: collection,
	}
}

func (h *ManufacturerHandlers) LoadManufacturers(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=json"

	log := logger.Get()

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var manufacturers struct {
		Results []models.Manufacturer `json:"Results"`
	}

	err = json.Unmarshal(body, &manufacturers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, manufacturer := range manufacturers.Results {
		err = h.repo.CreateManufacturer(r.Context(), h.collection, manufacturer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Manufacturers loaded successfully")
}

func (h *ManufacturerHandlers) GetManufacturers(w http.ResponseWriter, r *http.Request) {
	manufacturers, err := h.repo.GetAllManufacturers(r.Context(), h.collection)
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

	manufacturer, err := h.repo.GetManufacturer(r.Context(), h.collection, id)
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

	err = h.repo.CreateManufacturer(r.Context(), h.collection, manufacturer)
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

	err = h.repo.UpdateManufacturer(r.Context(), h.collection, id, manufacturer)
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

	err = h.repo.DeleteManufacturer(r.Context(), h.collection, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
