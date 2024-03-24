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

func generateLatitude() string {
	return fmt.Sprintf("%d", rand.Intn(241)-120)
}

func generateLongitude() string {
	return fmt.Sprintf("%d", rand.Intn(241)-120)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

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
