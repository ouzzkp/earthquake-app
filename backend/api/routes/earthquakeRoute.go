package routes

import (
	"backend/api/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes fonksiyonu, Earthquake ile ilgili route'ları tanımlar ve router'ı döndürür
func SetupEarthquakeRoutes(earthquakeHandler *handlers.EarthquakeHandler) *mux.Router {
	r := mux.NewRouter()

	// Earthquake ile ilgili HTTP endpoint'lerinin tanımlanması
	r.HandleFunc("/earthquakes", earthquakeHandler.GetAllEarthquakes).Methods("GET")
	r.HandleFunc("/earthquakes", earthquakeHandler.CreateEarthquake).Methods("POST")
	r.HandleFunc("/earthquakes/{id}", earthquakeHandler.GetEarthquakeByID).Methods("GET")
	r.HandleFunc("/earthquakes/{id}", earthquakeHandler.UpdateEarthquake).Methods("PUT")
	r.HandleFunc("/earthquakes/{id}", earthquakeHandler.DeleteEarthquake).Methods("DELETE")

	return r
}
