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

	manufacturerRepo := repository.NewManufacturerRepo(cfg)

	mfHandler := handlers.NewManufacturerHandlers(manufacturerRepo)

	router := mux.NewRouter()

	mfRouter := router.PathPrefix("/api/v1/manufacturers").Subrouter()
	mfRouter.HandleFunc("", mfHandler.GetManufacturers).Methods(http.MethodGet)
	mfRouter.HandleFunc("", mfHandler.CreateManufacturer).Methods(http.MethodPost)
	mfRouter.HandleFunc("/{id}", mfHandler.GetManufacturer).Methods(http.MethodGet)
	mfRouter.HandleFunc("/{id}", mfHandler.UpdateManufacturer).Methods(http.MethodPut)
	mfRouter.HandleFunc("/{id}", mfHandler.DeleteManufacturer).Methods(http.MethodDelete)

	handlers.SwaggerUIHandler(*mfRouter)
	
	log.Println("Starting server on :8080")
	log.Fatal().Err(http.ListenAndServe(":8080", router))

}
