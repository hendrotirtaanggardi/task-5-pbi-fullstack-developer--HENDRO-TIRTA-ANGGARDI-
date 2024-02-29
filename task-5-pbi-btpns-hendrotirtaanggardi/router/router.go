package router

import (
	"pbi-hendrotirta-btpns/mod/controllers"
	"pbi-hendrotirta-btpns/mod/middleware"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthMiddleware)
	router.HandleFunc("/users/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/photos/create", controllers.CreatePhoto).Methods("POST")
	router.HandleFunc("/photos/all", controllers.GetPhotos).Methods("GET")
	router.HandleFunc("/photos/{photoId}", controllers.UpdatePhotos).Methods("PUT")
	router.HandleFunc("/photos/{photoId}", controllers.DeletePhotos).Methods("DELETE")

	return router
}
