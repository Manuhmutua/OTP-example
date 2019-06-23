package main

import (
	"fmt"
	"github.com/Manuhmutua/movies-backend-apis/app"
	"github.com/Manuhmutua/movies-backend-apis/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/auth", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/v1/user/resetOTP", controllers.Reset).Methods("POST")
	router.HandleFunc("/api/v1/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/v1/getMovies", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "9999" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
