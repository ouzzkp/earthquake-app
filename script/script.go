package script

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Earthquake struct {
	Id        int     `json:"id"`
	Latitude  string  `json:"Latitude"`
	Longitude string  `json:"Longitude"`
	Magnitude float64 `json:"Magnitude"`
	Time      string  `json:"Time"`
}

// generateCoord generates a random coordinate
func generateCoord() string {
	degrees := rand.Intn(180) - 90
	minutes := rand.Intn(60)
	seconds := rand.Float64() * 60

	return fmt.Sprintf("%d°%d'%4.2f''", degrees, minutes, seconds)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(5 * time.Second)

	for _ = range ticker.C {
		quake := Earthquake{
			Id:        rand.Intn(1000),
			Latitude:  generateCoord(),
			Longitude: generateCoord(),
			Magnitude: rand.Float64() * 10,
			Time:      time.Now().Format(time.RFC3339),
		}

		jsonData, err := json.Marshal(quake)
		if err != nil {
			fmt.Println("JSON marshalling error:", err)
			continue
		}

		resp, err := http.Post("http://localhost:3852/api/createEarthquake", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("POST request error:", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("Earthquake data sent: %+v\n", quake)
	}
}
