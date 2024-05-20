package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"petstore/configs/database/postgres"
	"petstore/docs"

	petRepoImport "petstore/internal/database/repositories/pet"
	petServiceImport "petstore/internal/service/pet"
	petControllerImport "petstore/internal/transport/http/in/pet"

	storeRepoImport "petstore/internal/database/repositories/store"
	storeServiceImport "petstore/internal/service/store"
	storeControllerImport "petstore/internal/transport/http/in/store"

	userRepoImport "petstore/internal/database/repositories/user"
	userServiceImport "petstore/internal/service/user"
	userControllerImport "petstore/internal/transport/http/in/user"

	"petstore/internal/models/address"

	"petstore/internal/models/category"
	"petstore/internal/models/customer"
	"petstore/internal/models/order"
	"petstore/internal/models/pet"
	"petstore/internal/models/tag"
	"petstore/internal/models/user"

	httpSwagger "gitflex.diasoft.ru/mvp-go/golang-libraries/http-swagger.git"
)

// @title petstore api
// @version 1.0
// @description petstore api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /petstore
func main() {
	gormdb := postgres.GetGorm()

	petRepo := petRepoImport.NewRepository(gormdb)
	petService := petServiceImport.NewService(petRepo)
	petController := petControllerImport.NewController(petService)

	storeRepo := storeRepoImport.NewRepository(gormdb)
	storeService := storeServiceImport.NewService(storeRepo)
	storeController := storeControllerImport.NewController(storeService)

	userRepo := userRepoImport.NewRepository(gormdb)
	userService := userServiceImport.NewService(userRepo)
	userController := userControllerImport.NewController(userService)

	r := mux.NewRouter()
	r.PathPrefix("/petstore/swagger/").Handler(httpSwagger.Handler()).Methods(http.MethodGet)

	initApiEndpoints(
		petController,
		storeController,
		userController,
		r,
	)

	docs.SwaggerInfo.BasePath = "/petstore"

	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	println(
		fmt.Sprintf("\n----------------------------------------------------------\n\t" +
			"Application \"petstore\" is running! Access URLs:\n\t" +
			"Local: \t\thttp://localhost:8080/petstore\n\t" +
			"Swagger UI: \thttp://localhost:8080/petstore/swagger/index.html\n\t" +
			"\n----------------------------------------------------------\n",
		),
	)
	migrate(gormdb)

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func initApiEndpoints(
	pet petControllerImport.Controller,
	store storeControllerImport.Controller,
	user userControllerImport.Controller,
	r *mux.Router,
) {
	r.HandleFunc("/petstore/pet", pet.UpdatePet).Methods("PUT")
	r.HandleFunc("/petstore/pet", pet.AddPet).Methods("POST")
	r.HandleFunc("/petstore/pet/findByStatus", pet.FindPetsByStatus).Methods("GET")
	r.HandleFunc("/petstore/pet/findByTags", pet.FindPetsByTags).Methods("GET")
	r.HandleFunc("/petstore/pet/{petId}", pet.GetPetById).Methods("GET")
	r.HandleFunc("/petstore/pet/{petId}", pet.UpdatePetWithForm).Methods("POST")
	r.HandleFunc("/petstore/pet/{petId}", pet.DeletePet).Methods("DELETE")
	r.HandleFunc("/petstore/pet/{petId}/uploadImage", pet.UploadFile).Methods("POST")
	r.HandleFunc("/petstore/store/inventory", store.GetInventory).Methods("GET")
	r.HandleFunc("/petstore/store/order", store.PlaceOrder).Methods("POST")
	r.HandleFunc("/petstore/store/order/{orderId}", store.GetOrderById).Methods("GET")
	r.HandleFunc("/petstore/store/order/{orderId}", store.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/petstore/user", user.CreateUser).Methods("POST")
	r.HandleFunc("/petstore/user/createWithList", user.CreateUsersWithListInput).Methods("POST")
	r.HandleFunc("/petstore/user/login", user.LoginUser).Methods("GET")
	r.HandleFunc("/petstore/user/logout", user.LogoutUser).Methods("GET")
	r.HandleFunc("/petstore/user/{username}", user.GetUserByName).Methods("GET")
	r.HandleFunc("/petstore/user/{username}", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/petstore/user/{username}", user.DeleteUser).Methods("DELETE")
}

func migrate(gormdb *gorm.DB) {
	if err := gormdb.Migrator().AutoMigrate(
		address.AddressEntity{},

		category.CategoryEntity{},

		customer.CustomerEntity{},

		order.OrderEntity{},

		pet.PetEntity{},

		tag.TagEntity{},

		user.UserEntity{},
	); err != nil {
		panic(err)
	}
}
