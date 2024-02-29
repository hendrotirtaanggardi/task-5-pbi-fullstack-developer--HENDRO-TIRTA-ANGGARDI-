package controllers

import (
	"encoding/json"
	"net/http"
	"pbi-hendrotirta-btpns/mod/app"
	"pbi-hendrotirta-btpns/mod/database"
	"pbi-hendrotirta-btpns/mod/helpers"

	"github.com/gorilla/mux"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user app.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user app.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var userFromDB app.User
	if err := database.DB.Where("email = ?", user.Email).First(&userFromDB).Error; err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	hashedPasswordFromDatabase := userFromDB.Password

	if !helpers.CheckPasswordHash(user.Password, hashedPasswordFromDatabase) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := helpers.GenerateToken(userFromDB.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	var user app.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	var updatedUser app.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Username = updatedUser.Username
	user.Email = updatedUser.Email

	if err := database.DB.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	var user app.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
