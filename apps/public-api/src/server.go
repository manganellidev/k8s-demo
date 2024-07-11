package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Holidays(w http.ResponseWriter, r *http.Request) {
	countryParam := r.URL.Query().Get("country")
	if countryParam == "" {
		http.Error(w, "Country query parameter is required", http.StatusBadRequest)
		return
	}

	// Dummy holiday data for demonstration purposes
	holidays := map[string][]string{
		"US": {"Independence Day: July 4", "Thanksgiving: November 25"},
		"GB": {"Christmas: December 25", "Boxing Day: December 26"},
		"DE": {"Unity Day: October 3", "Christmas: December 25"},
		"JP": {"New Year's Day: January 1", "Golden Week: April 29 - May 5"},
		"BR": {"Carnival: February 28", "Independence Day: September 7"},
	}

	countryHolidays, exists := holidays[countryParam]
	if !exists {
		http.Error(w, "Invalid country code", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Holidays in %s:\n%s", countryParam, strings.Join(countryHolidays, "\n"))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/holidays", Holidays)
	http.HandleFunc("/healthz", Healthz)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
