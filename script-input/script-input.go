package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
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

func main() {
	latPtr := flag.String("lat", "0", "Latitude value")
	lonPtr := flag.String("lon", "0", "Longitude value")
	magPtr := flag.Float64("mag", 0.0, "Magnitude value")

	flag.Parse()

	quake := Earthquake{
		Latitude:  *latPtr,
		Longitude: *lonPtr,
		Magnitude: *magPtr,
		Time:      time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(quake)
	if err != nil {
		log.Fatal("JSON marshalling error: ", err)
	}

	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		log.Fatal("BACKEND_URL environment variable is not set")
	}

	requestURL := fmt.Sprintf("%s/earthquakes", backendURL)
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error while POST process", err)
	}
	defer resp.Body.Close()

	// Cevap loglanıyor.
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal("Error while reading response:", err)
	}
	log.Println("Response Status:", resp.Status)
	log.Printf("Sended eartquake data: %+v\n", quake)
}
