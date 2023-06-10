package main

import (
	"apirest/db"
	"apirest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	// db.CreateTable(models.UserSchema, "users")

	// user := models.CreaUser("Sebatian Camacho", "123456", "Camacho@gmail.com")
	// user.Save()

	mux := mux.NewRouter()

	//EndPoint

	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))

	defer db.Close()
}
