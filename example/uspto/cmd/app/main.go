package main

import (
	"example/uspto/configs/database/postgres"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"uspto/docs"

	metadataRepoImport "example/uspto/internal/database/repositories/metadata"
	metadataServiceImport "example/uspto/internal/service/metadata"
	metadataControllerImport "example/uspto/internal/transport/http/in/metadata"

	searchRepoImport "example/uspto/internal/database/repositories/search"
	searchServiceImport "example/uspto/internal/service/search"
	searchControllerImport "example/uspto/internal/transport/http/in/search"

	httpSwagger "gitflex.diasoft.ru/mvp-go/golang-libraries/http-swagger.git"
)

// @title uspto api
// @version 1.0
// @description uspto api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /uspto
func main() {
	gormdb := postgres.GetGorm()

	metadataRepo := metadataRepoImport.NewRepository(gormdb)
	metadataService := metadataServiceImport.NewService(metadataRepo)
	metadataController := metadataControllerImport.NewController(metadataService)

	searchRepo := searchRepoImport.NewRepository(gormdb)
	searchService := searchServiceImport.NewService(searchRepo)
	searchController := searchControllerImport.NewController(searchService)

	r := mux.NewRouter()
	r.PathPrefix("/uspto/swagger/").Handler(httpSwagger.Handler()).Methods(http.MethodGet)

	initApiEndpoints(
		metadataController,
		searchController,
		r,
	)

	docs.SwaggerInfo.BasePath = "/uspto"

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	println(
		fmt.Sprintf("\n----------------------------------------------------------\n\t" +
			"Application \"uspto\" is running! Access URLs:\n\t" +
			"Local: \t\thttp://localhost:8080/uspto\n\t" +
			"Swagger UI: \thttp://localhost:8080/uspto/swagger/index.html\n\t" +
			"\n----------------------------------------------------------\n",
		),
	)
	migrate(gormdb)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initApiEndpoints(
	metadata metadataControllerImport.Controller,
	search searchControllerImport.Controller,
	r *mux.Router,
) {
	r.HandleFunc("/uspto/", metadata.ListDataSets).Methods("GET")
	r.HandleFunc("/uspto/{dataset}/{version}/fields", metadata.ListSearchableFields).Methods("GET")
	r.HandleFunc("/uspto/{dataset}/{version}/records", search.PerformSearch).Methods("POST")
}

func migrate(gormdb *gorm.DB) {
	if err := gormdb.Migrator().AutoMigrate(); err != nil {
		panic(err)
	}
}
