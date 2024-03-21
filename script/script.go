package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Earthquake struct {
	Id        int     `json:"id"`
	Latitude  string  `json:"Latitude"`
	Longitude string  `json:"Longitude"`
	Magnitude float64 `json:"Magnitude"`
	Time      string  `json:"Time"`
}

// generateCoord generates a random coordinate in degrees, minutes, and seconds
func generateLatitude() string {
	return fmt.Sprintf("%d", rand.Intn(241)-120) // -90 ile +90 arasında rastgele bir değer üretir
}

// generateLongitude generates a random longitude between -120 and +120
func generateLongitude() string {
	return fmt.Sprintf("%d", rand.Intn(241)-120) // -120 ile +120 arasında rastgele bir değer üretir
}

func main() {
	// Random seed initialization
	rand.Seed(time.Now().UnixNano())
	// Ticker for periodic execution
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Get the backend URL from the environment variable
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		log.Fatal("BACKEND_URL environment variable is not set.")
	}

	for range ticker.C {
		quake := Earthquake{
			Latitude:  generateLatitude(),
			Longitude: generateLongitude(),
			Magnitude: rand.Float64() * 10,
			Time:      time.Now().Format(time.RFC3339),
		}

		jsonData, err := json.Marshal(quake)
		if err != nil {
			log.Println("Error marshalling JSON:", err)
			continue
		}

		// Use the BACKEND_URL environment variable to construct the request URL
		requestURL := fmt.Sprintf("%s/earthquakes", backendURL)
		resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Println("Error sending POST request:", err)
			continue
		}

		defer resp.Body.Close()

		log.Printf("Earthquake data sent: %+v\n", quake)
	}
}
