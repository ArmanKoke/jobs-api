package main

import (
	"jobs-api/controllers"
	health "jobs-api/utils/checker"
	"jobs-api/utils/databases"
	"jobs-api/utils/env"
	"jobs-api/utils/logger"
	utilSwagger "jobs-api/utils/swagger"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload" //autoload env
	"go.elastic.co/apm/module/apmgorilla"
	"go.uber.org/zap"
)

//go:generate swagger generate spec
func main() {
	//instances
	logInstance := logger.NewLogEnv()
	log := logInstance.ZapLogger()
	defer log.Sync()

	var swaggerInstance utilSwagger.Swagger
	swaggerInstance.Log = logInstance

	dbInstance, err := databases.NewDbEnv()
	if err != nil {
		log.Fatal("Unable to init databases", zap.Error(err))
	}

	//health
	health := health.Create()

	router := mux.NewRouter()
	router.Path("/check").HandlerFunc(health.Get).Methods(http.MethodGet, http.MethodOptions)
	router.Path("/swagger.json").HandlerFunc(swaggerInstance.GetSwagger).Methods(http.MethodGet, http.MethodOptions, http.MethodHead)

	reviewControllerInstance := controllers.NewReviewController(logInstance, dbInstance)
	router.HandleFunc("/reviews", reviewControllerInstance.Show).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/review/create", reviewControllerInstance.Create).Methods(http.MethodPost, http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))

	apmgorilla.Instrument(router)

	log.Info("Ready to serve", zap.String(env.IPAddress, env.AppIPAddress()), zap.String(env.Port, env.AppPort()))
	err = http.ListenAndServe(env.AppIPAddress()+":"+env.AppPort(), router)
	if err != nil {
		log.Fatal("Unable to bind to port", zap.Error(err))
	}
}
