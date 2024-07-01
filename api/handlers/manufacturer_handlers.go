package handlers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ch374n/vehicles-app/internal/models"
	"github.com/ch374n/vehicles-app/internal/repository"
	"github.com/ch374n/vehicles-app/logger"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

const (
	apiURL = "https://vpic.nhtsa.dot.gov/api/vehicles/getallmanufacturers?format=json"
)

type ManufacturerHandlers struct {
	repo        repository.ManufacturerRepo
	collection  *mongo.Collection
	redisClient *redis.Client
}

func NewManufacturerHandlers(r *repository.ManufacturerRepo, collection *mongo.Collection, redisClient *redis.Client) *ManufacturerHandlers {
	return &ManufacturerHandlers{
		repo:        *r,
		collection:  collection,
		redisClient: redisClient,
	}
}

func (h *ManufacturerHandlers) LoadManufacturers(w http.ResponseWriter, r *http.Request) {

	log := logger.Get()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(apiURL)

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

		mfrJSON, _ := json.Marshal(manufacturer)

		err := h.redisClient.Set(fmt.Sprintf("mfr:%d", manufacturer.MfrID), mfrJSON, 24*time.Hour).Err()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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

	var manufacturers []models.Manufacturer
	var cursor uint64 = 0

	for {
		var keys []string
		var err error

		keys, cursor, err = h.redisClient.Scan(cursor, "mfr:*", 50).Result()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for _, key := range keys {

			val, err := h.redisClient.Get(key).Result()

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			var mfr models.Manufacturer

			err = json.Unmarshal([]byte(val), &mfr)

			if err != nil {
				log.Printf("Failed to unmarshal manufacturer: %v", err)
				continue
			}

			manufacturers = append(manufacturers, mfr)
		}

		if cursor == 0 {
			break
		}
	}

	json.NewEncoder(w).Encode(manufacturers)

}

func (h *ManufacturerHandlers) GetManufacturer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val, err := h.redisClient.Get(fmt.Sprintf("mfr:%d", id)).Result()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var mfr models.Manufacturer

	err = json.Unmarshal([]byte(val), &mfr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if mfr.MfrName != "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mfr)
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

	mfrJSON, _ := json.Marshal(manufacturer)

	err = h.redisClient.Set(fmt.Sprintf("mfr:%d", manufacturer.MfrID), mfrJSON, 24*time.Hour).Err()

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

	val, err := h.redisClient.Get(fmt.Sprintf("mfr:%d", id)).Result()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var mfr models.Manufacturer

	err = json.Unmarshal([]byte(val), &mfr)

	if err != nil {
		log.Printf("Failed to unmarshal manufacturer: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mfrJSON, _ := json.Marshal(manufacturer)

	err = h.redisClient.Set(fmt.Sprintf("mfr:%d", manufacturer.MfrID), mfrJSON, 24*time.Hour).Err()

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

	_, err = h.redisClient.Del(fmt.Sprintf("mfr:%d", id)).Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
