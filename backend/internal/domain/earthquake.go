package domain

import "time"

type Earthquake struct {
	Id        int       `json:"id"`
	Latitude  string    `json:"Latitude"`
	Longitude string    `json:"Longitude"`
	Magnitude float64   `json:"Magnitude"`
	Time      time.Time `json:"Time"`
}
