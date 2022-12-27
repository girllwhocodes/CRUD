package controllers

import (
	"edu_service_go/database"
	"edu_service_go/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var document models.Document
	json.NewDecoder(r.Body).Decode(&document)
	database.Instance.Create(&document)

	json.NewEncoder(w).Encode(document)
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputID := params["docId"]
	if !checkIfDocumentExists(inputID) {
		json.NewEncoder(w).Encode("Document Not Found!")
		return
	}
	var document models.Document
	database.Instance.First(&document, inputID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(document)
}

func checkIfDocumentExists(docId string) bool {
	var document models.Document
	database.Instance.First(&document, docId)
	if document.ID == 0 {
		return false
	}
	return true
}

func GetDocuments(w http.ResponseWriter, r *http.Request) {
	var documents []models.Document
	database.Instance.Find(&documents)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(documents)
}

func UpdateDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["docId"]
	if !checkIfDocumentExists(inputID) {
		json.NewEncoder(w).Encode("Document Not Found!")
		return
	}
	var document models.Document
	database.Instance.First(&document, inputID)
	json.NewDecoder(r.Body).Decode(&document)
	database.Instance.Save(&document)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(document)
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputID := params["docId"]
	if !checkIfDocumentExists(inputID) {
		json.NewEncoder(w).Encode("Document Not Found!")
		return
	}
	var document models.Document
	database.Instance.Delete(&document, inputID)
	json.NewEncoder(w).Encode("Deleted successfully!")
}
