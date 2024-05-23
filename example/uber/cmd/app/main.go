package main

import (
	"example/uber/configs/database/postgres"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"uber/docs"

	httpSwagger "gitflex.diasoft.ru/mvp-go/golang-libraries/http-swagger.git"
)

// @title uber api
// @version 1.0
// @description uber api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /uber
func main() {
	gormdb := postgres.GetGorm()

	r := mux.NewRouter()
	r.PathPrefix("/uber/swagger/").Handler(httpSwagger.Handler()).Methods(http.MethodGet)

	initApiEndpoints(
		r,
	)

	docs.SwaggerInfo.BasePath = "/uber"

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	println(
		fmt.Sprintf("\n----------------------------------------------------------\n\t" +
			"Application \"uber\" is running! Access URLs:\n\t" +
			"Local: \t\thttp://localhost:8080/uber\n\t" +
			"Swagger UI: \thttp://localhost:8080/uber/swagger/index.html\n\t" +
			"\n----------------------------------------------------------\n",
		),
	)
	migrate(gormdb)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initApiEndpoints(
	r *mux.Router,
) {
}

func migrate(gormdb *gorm.DB) {
	if err := gormdb.Migrator().AutoMigrate(); err != nil {
		panic(err)
	}
}
