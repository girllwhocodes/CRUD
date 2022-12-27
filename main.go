package main

import (
	"edu_service_go/controllers"
	"edu_service_go/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect("host=localhost user=postgres password=Dream2408 dbname=CRUD port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	database.Migrate()
	router := mux.NewRouter()

	//Routing for Users
	//Create
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	//Read
	router.HandleFunc("/users/{userId}", controllers.GetUser).Methods("GET")
	//Read-All
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	//Update
	router.HandleFunc("/users/{userId}", controllers.UpdateUser).Methods("PUT")
	//Delete
	router.HandleFunc("/users/{userId}", controllers.DeleteUser).Methods("DELETE")

	//Routing for Educations
	//Create
	router.HandleFunc("/educations", controllers.CreateEdu).Methods("POST")
	//Read
	router.HandleFunc("/educations/{eduId}", controllers.GetEdu).Methods("GET")
	//Read-All
	router.HandleFunc("/educations", controllers.GetEducations).Methods("GET")
	//Update
	router.HandleFunc("/educations/{eduId}", controllers.UpdateEdu).Methods("PUT")
	//Delete
	router.HandleFunc("/educations/{eduId}", controllers.DeleteEdu).Methods("DELETE")

	//Routing for Documents
	//Create
	router.HandleFunc("/documents", controllers.CreateDocument).Methods("POST")
	//Read
	router.HandleFunc("/documents/{docId}", controllers.GetDocument).Methods("GET")
	//Read-All
	router.HandleFunc("/documents", controllers.GetDocuments).Methods("GET")
	//Update
	router.HandleFunc("/documents/{docId}", controllers.UpdateDocument).Methods("PUT")
	//Delete
	router.HandleFunc("/documents/{docId}", controllers.DeleteDocument).Methods("DELETE")

	log.Println("...Starting on the server :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
