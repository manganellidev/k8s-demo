package main

import (
	"fmt"
	"net/http"
	"time"
)

var countryRegionsTimeZones = map[string]map[string]string{
	"US": {
		"New_York": "America/New_York",
	},
	"GB": {
		"London": "Europe/London",
	},
	"DE": {
		"Berlin": "Europe/Berlin",
	},
	"JP": {
		"Tokyo": "Asia/Tokyo",
	},
	"BR": {
		"Sao_Paulo": "America/Sao_Paulo",
	},
}

func Time(w http.ResponseWriter, r *http.Request) {
	countryParam := r.URL.Query().Get("country")
	regionParam := r.URL.Query().Get("region")
	if countryParam == "" || regionParam == "" {
		http.Error(w, "Country and region query parameters are required", http.StatusBadRequest)
		return
	}

	regions, exists := countryRegionsTimeZones[countryParam]
	if !exists {
		http.Error(w, "Invalid country code", http.StatusBadRequest)
		return
	}

	timeZone, exists := regions[regionParam]
	if !exists {
		http.Error(w, "Invalid region for the specified country", http.StatusBadRequest)
		return
	}

	location, err := time.LoadLocation(timeZone)
	if err != nil {
		http.Error(w, "Failed to load time zone", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().In(location)

	fmt.Fprintf(w, "Hello, it is %s in %s, %s.", currentTime.Format("Mon Jan 2 15:04:05 MST 2006"), regionParam, countryParam)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/time", Time)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8080", nil)
}
