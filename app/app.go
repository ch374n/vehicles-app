package app

import (
	"net/http"

	"github.com/ch374n/vehicles-app/api/handlers"
	config "github.com/ch374n/vehicles-app/configs"
	"github.com/ch374n/vehicles-app/internal/repository"
	"github.com/ch374n/vehicles-app/logger"
	"github.com/ch374n/vehicles-app/pkg/database"
	"github.com/gorilla/mux"
)

func Run() {

	log := logger.Get()
	cfg, err := config.Load(log)

	log.Info().Msgf("%v", cfg)

	if err != nil {
		log.Fatal().Err(err).Msg("Unable to load the application configuration")
	}

	log.Info().Msgf("vehicles service listening on %s", cfg.Port)

	err = database.ConnectMongoDB(cfg.MongoUri)

	if err != nil {
		log.Fatal().Msgf("Failed to connect to MongoDB: %v", err)
	}
	defer database.DisconnectMongoDB()

	manufacturerRepo := repository.NewManufacturerRepo(database.MongoClient.Database(cfg.DBName).Collection(cfg.CollectionName))

	mfHandler := handlers.NewManufacturerHandlers(&manufacturerRepo)

	router := mux.NewRouter()

	handlers.SwaggerUIHandler(router)

	mfRouter := router.PathPrefix("/api/v1/manufacturers").Subrouter()
	mfRouter.HandleFunc("", mfHandler.GetManufacturers).Methods(http.MethodGet)
	mfRouter.HandleFunc("", mfHandler.CreateManufacturer).Methods(http.MethodPost)
	mfRouter.HandleFunc("/load", mfHandler.LoadManufacturers).Methods(http.MethodGet)
	mfRouter.HandleFunc("/{id}", mfHandler.GetManufacturer).Methods(http.MethodGet)
	mfRouter.HandleFunc("/{id}", mfHandler.UpdateManufacturer).Methods(http.MethodPut)
	mfRouter.HandleFunc("/{id}", mfHandler.DeleteManufacturer).Methods(http.MethodDelete)

	log.Println("Starting server on :8081")
	log.Fatal().Err(http.ListenAndServe(":8081", router))

}
