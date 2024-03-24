package routes

import (
	"backend/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupEarthquakeRoutes(earthquakeHandler *handlers.EarthquakeHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/earthquakes", earthquakeHandler.GetAllEarthquakes).Methods("GET")
	r.HandleFunc("/earthquakes", earthquakeHandler.CreateEarthquake).Methods("POST")
	r.HandleFunc("/earthquakes/{id}", earthquakeHandler.GetEarthquakeByID).Methods("GET")
	r.HandleFunc("/earthquakes/{id}", earthquakeHandler.UpdateEarthquake).Methods("PUT")
	r.HandleFunc("/earthquakes/{id}", earthquakeHandler.DeleteEarthquake).Methods("DELETE")
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		earthquakeHandler.WebSocketHandler(w, r)
	})
	return r
}
