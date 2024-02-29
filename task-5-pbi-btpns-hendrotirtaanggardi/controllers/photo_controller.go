package controllers

import (
	"encoding/json"
	"net/http"
	"pbi-hendrotirta-btpns/mod/database"
	"pbi-hendrotirta-btpns/mod/models"

	"github.com/gorilla/mux"
)

func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	var photo models.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&photo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(photo)
}

func GetPhotos(w http.ResponseWriter, r *http.Request) {
	var photos []models.Photo

	if err := database.DB.Find(&photos).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(photos)
}

func UpdatePhotos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoID := params["photoId"]
	var photo models.Photo
	if err := database.DB.First(&photo, photoID).Error; err != nil {
		http.Error(w, "Photo not found", http.StatusNotFound)
		return
	}

	var updatedPhoto models.Photo
	err := json.NewDecoder(r.Body).Decode(&updatedPhoto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	photo.Title = updatedPhoto.Title
	photo.Caption = updatedPhoto.Caption
	photo.PhotoUrl = updatedPhoto.PhotoUrl
	photo.UserID = updatedPhoto.UserID

	if err := database.DB.Save(&photo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(photo)
}

func DeletePhotos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	photoID := params["photoId"]

	var photo models.Photo
	if err := database.DB.First(&photo, photoID).Error; err != nil {
		http.Error(w, "Photo not found", http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&photo).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
