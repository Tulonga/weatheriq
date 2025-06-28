package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type WeatherRes struct {
	City			string	`json:"city"`
	Temperature		float64	`json:"temperature"`
	Condition		string	`json:"condition"`
}

func fetchWeather(city string) WeatherRes {
	return WeatherRes{
		City:			city,
		Temperature:	30,
		Condition:		"Partial Sunny",
	}
}

func weatherHandles(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "Missing 'city' City needed", http.StatusBadRequest)
		return
	}

	weather := fetchWeather(city)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(weather)
}

func main() {

	http.HandleFunc("/weather", weatherHandles)
	
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
