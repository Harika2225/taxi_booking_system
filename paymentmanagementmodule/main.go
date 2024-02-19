package main

import (
	"net/http"

	auth "com.example.paymentmanagement/auth"
	app "com.example.paymentmanagement/config"
	"com.example.paymentmanagement/controllers"
	eureka "com.example.paymentmanagement/eurekaregistry"
	"com.example.paymentmanagement/handler"
	"com.example.paymentmanagement/migrate"
	"github.com/asim/go-micro/v3"
	mhttp "github.com/go-micro/plugins/v3/server/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/micro/micro/v3/service/logger"
)

var configurations eureka.RegistrationVariables

func main() {
	defer cleanup()
	app.Setconfig()
	migrate.MigrateAndCreateDatabase()
	auth.SetClient()
	handler.InitializeDb()
	service_registry_url := app.GetVal("GO_MICRO_SERVICE_REGISTRY_URL")
	InstanceId := "paymentmanagementmodule:" + uuid.New().String()
	configurations = eureka.RegistrationVariables{ServiceRegistryURL: service_registry_url, InstanceId: InstanceId}
	port := app.GetVal("GO_MICRO_SERVICE_PORT")
	srv := micro.NewService(
		micro.Server(mhttp.NewServer()),
	)
	opts1 := []micro.Option{
		micro.Name("paymentmanagementmodule"),
		micro.Version("latest"),
		micro.Address(":" + port),
	}
	srv.Init(opts1...)
	r := mux.NewRouter().StrictSlash(true)
	r.Use(corsMiddleware)
	registerRoutes(r)
	var handlers http.Handler = r

	go eureka.ManageDiscovery(configurations)

	if err := micro.RegisterHandler(srv.Server(), handlers); err != nil {
		logger.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func cleanup() {
	eureka.Cleanup(configurations)
}

func registerRoutes(router *mux.Router) {
	registerControllerRoutes(controllers.PaymentController{}, router)
	registerControllerRoutes(controllers.CommunicationController{}, router)
	registerControllerRoutes(controllers.ManagementController{}, router)
}

func registerControllerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept,Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
