package handlers

import (
	"LoadingData/internal/auth"
	"LoadingData/internal/db"
	"LoadingData/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)


func GetPlacesHandler(store db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 10 
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			offset = 0 
		}

		places, total, err := store.GetPlaces(limit, offset)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := struct {
			Total  int            `json:"total"`
			Places []models.Place `json:"places"`
		}{
			Total:  total,
			Places: places,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}


func GetNearestPlacesHandler(store db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		if err != nil {
			http.Error(w, "Invalid 'lat' value", http.StatusBadRequest)
			return
		}
		lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
		if err != nil {
			http.Error(w, "Invalid 'lon' value", http.StatusBadRequest)
			return
		}

		places, err := store.GetNearestPlaces(lat, lon, 3) 
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		response := struct {
			Name   string         `json:"name"`
			Places []models.Place `json:"places"`
		}{
			Name:   "Recommendation",
			Places: places,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}


func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "admin" && password == "password" {
		token, err := auth.CreateToken(username)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			return
		}
		response := map[string]string{"token": token}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}


func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	_, err := auth.ValidateToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the protected endpoint"))
}
