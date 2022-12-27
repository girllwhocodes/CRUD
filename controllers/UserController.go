package controllers

import (
	"edu_service_go/database"
	"edu_service_go/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Create(&user)

	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputID := params["userId"]
	if !checkIfUserExists(inputID) {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var user models.User
	database.Instance.First(&user, inputID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func checkIfUserExists(userId string) bool {
	var user models.User
	database.Instance.First(&user, userId)
	if user.ID == 0 {
		return false
	}
	return true
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.Instance.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["userId"]
	if !checkIfUserExists(inputID) {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var user models.User
	database.Instance.First(&user, inputID)
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["userId"]
	if !checkIfUserExists(inputID) {
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var user models.User
	database.Instance.Delete(&user, inputID)
	json.NewEncoder(w).Encode("Deleted successfully!")
}
