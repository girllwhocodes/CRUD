package controllers

import (
	"edu_service_go/database"
	"edu_service_go/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateEdu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var edu models.Education
	json.NewDecoder(r.Body).Decode(&edu)
	database.Instance.Create(&edu)

	json.NewEncoder(w).Encode(edu)
}

func GetEdu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputID := params["eduId"]
	if !checkIfEduExists(inputID) {
		json.NewEncoder(w).Encode("Education Not Found!")
		return
	}
	var edu models.Education
	database.Instance.First(&edu, inputID)
	json.NewEncoder(w).Encode(edu)
}

func checkIfEduExists(eduId string) bool {
	var edu models.Education
	database.Instance.First(&edu, eduId)
	if edu.ID == 0 {
		return false
	}
	return true
}

func GetEducations(w http.ResponseWriter, r *http.Request) {
	var educations []models.Education
	database.Instance.Find(&educations)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(educations)
}

func UpdateEdu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["eduId"]
	if !checkIfEduExists(inputID) {
		json.NewEncoder(w).Encode("Education Not Found!")
		return
	}
	var edu models.Education
	database.Instance.First(&edu, inputID)
	json.NewDecoder(r.Body).Decode(&edu)
	database.Instance.Save(&edu)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(edu)
}

func DeleteEdu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["eduId"]
	if !checkIfEduExists(inputID) {
		json.NewEncoder(w).Encode("Education Not Found!")
		return
	}
	var edu models.Education
	database.Instance.Delete(&edu, inputID)
	json.NewEncoder(w).Encode("Deleted successfully!")
}
