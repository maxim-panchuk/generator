package main

import (
	"example/link/configs/database/postgres"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"link/docs"
	"net/http"

	repoRepoImport "example/link/internal/database/repositories/repo"
	repoServiceImport "example/link/internal/service/repo"
	repoControllerImport "example/link/internal/transport/http/in/repo"

	userRepoImport "example/link/internal/database/repositories/user"
	userServiceImport "example/link/internal/service/user"
	userControllerImport "example/link/internal/transport/http/in/user"

	httpSwagger "gitflex.diasoft.ru/mvp-go/golang-libraries/http-swagger.git"
)

// @title link api
// @version 1.0
// @description link api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /link
func main() {
	gormdb := postgres.GetGorm()

	repoRepo := repoRepoImport.NewRepository(gormdb)
	repoService := repoServiceImport.NewService(repoRepo)
	repoController := repoControllerImport.NewController(repoService)

	userRepo := userRepoImport.NewRepository(gormdb)
	userService := userServiceImport.NewService(userRepo)
	userController := userControllerImport.NewController(userService)

	r := mux.NewRouter()
	r.PathPrefix("/link/swagger/").Handler(httpSwagger.Handler()).Methods(http.MethodGet)

	initApiEndpoints(
		repoController,
		userController,
		r,
	)

	docs.SwaggerInfo.BasePath = "/link"

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	println(
		fmt.Sprintf("\n----------------------------------------------------------\n\t" +
			"Application \"link\" is running! Access URLs:\n\t" +
			"Local: \t\thttp://localhost:8080/link\n\t" +
			"Swagger UI: \thttp://localhost:8080/link/swagger/index.html\n\t" +
			"\n----------------------------------------------------------\n",
		),
	)
	migrate(gormdb)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initApiEndpoints(
	repo repoControllerImport.Controller,
	user userControllerImport.Controller,
	r *mux.Router,
) {
	r.HandleFunc("/link/2.0/repositories/{username}", repo.GetRepositoriesByOwner).Methods("GET")
	r.HandleFunc("/link/2.0/repositories/{username}/{slug}", repo.GetRepository).Methods("GET")
	r.HandleFunc("/link/2.0/repositories/{username}/{slug}/pullrequests", repo.GetPullRequestsByRepository).Methods("GET")
	r.HandleFunc("/link/2.0/repositories/{username}/{slug}/pullrequests/{pid}", repo.GetPullRequestsById).Methods("GET")
	r.HandleFunc("/link/2.0/repositories/{username}/{slug}/pullrequests/{pid}/merge", repo.MergePullRequest).Methods("POST")
	r.HandleFunc("/link/2.0/users/{username}", user.GetUserByName).Methods("GET")
}

func migrate(gormdb *gorm.DB) {
	if err := gormdb.Migrator().AutoMigrate(); err != nil {
		panic(err)
	}
}
