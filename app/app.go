package app

import (
	"net/http"

	"github.com/ch374n/vehicles-app/api/handlers"
	config "github.com/ch374n/vehicles-app/configs"
	"github.com/ch374n/vehicles-app/internal/repository"
	"github.com/ch374n/vehicles-app/logger"
	"github.com/ch374n/vehicles-app/middleware"
	"github.com/ch374n/vehicles-app/pkg/database"
	"github.com/ch374n/vehicles-app/scheduler"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	manufacturerRepo := repository.NewManufacturerRepo()
	collection := database.MongoClient.Database(cfg.DBName).Collection(cfg.CollectionName)
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisConnection,
	})

	mfHandler := handlers.NewManufacturerHandlers(
		&manufacturerRepo,
		collection,
		redisClient,
	)

	router := mux.NewRouter()

	handlers.SwaggerUIHandler(router)

	metricsMiddleware := middleware.NewMetricsMiddleware()

	router.Use(metricsMiddleware.Metrics)
	router.Handle("/metrics", promhttp.Handler())

	mfRouter := router.PathPrefix("/api/v1/manufacturers").Subrouter()
	mfRouter.HandleFunc("", mfHandler.GetManufacturers).Methods(http.MethodGet)
	mfRouter.HandleFunc("", mfHandler.CreateManufacturer).Methods(http.MethodPost)
	mfRouter.HandleFunc("/load", mfHandler.LoadManufacturers).Methods(http.MethodGet)
	mfRouter.HandleFunc("/{id}", mfHandler.GetManufacturer).Methods(http.MethodGet)
	mfRouter.HandleFunc("/{id}", mfHandler.UpdateManufacturer).Methods(http.MethodPut)
	mfRouter.HandleFunc("/{id}", mfHandler.DeleteManufacturer).Methods(http.MethodDelete)

	go func() {
		scheduler := scheduler.NewScheduler(collection, redisClient)

		scheduler.SyncToMongodb()
	}()

	log.Println("Starting server on :8081")
	log.Fatal().Err(http.ListenAndServe(":8081", router))

}
