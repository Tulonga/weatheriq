package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

type WeatherRes struct {
	City			string	`json:"city"`
	Temperature		float64	`json:"temperature"`
	Condition		string	`json:"condition"`
}

type WeatherAPIRes struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`

	Current struct {
		TempC	float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

func fetchWeather(city string, apiKey string) (WeatherRes, error)	{
		url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

		resp, err := http.Get(url)
		if err != nil {
			return WeatherRes{}, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			return WeatherRes{}, err
		}

		var apiResp WeatherAPIRes
		err = json.Unmarshal(body, &apiResp)
		if err != nil{
			return WeatherRes{}, err
		}

		result := WeatherRes{
			City: 				apiResp.Location.Name,
			Temperature:		apiResp.Current.TempC,
			Condition:			apiResp.Current.Condition.Text,
		}
		return result, nil
}

func weatherHandles(apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		

	city := r.URL.Query().Get("city")

	if city == "" {
		http.Error(w, "Missing 'city' City needed", http.StatusBadRequest)
		return
	}

	weather, err := fetchWeather(city, apiKey)
	if err != nil{
		http.Error(w, "Failed to get weather data (lol)", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(weather)
	}
}

func main() {

	apiKey := os.Getenv("WEATHERAPI_KEY")
	if apiKey == ""{
		log.Fatal("Please set the WEATHERAPI_KEY environment variable")
	}

	http.HandleFunc("/weather", weatherHandles(apiKey))
	
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
