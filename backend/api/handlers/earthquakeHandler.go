package handlers

import (
	"backend/internal/domain"
	"backend/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type EarthquakeHandler struct {
	earthquakeService service.EarthquakeService
}

func NewEarthquakeHandler(earthquakeService service.EarthquakeService) *EarthquakeHandler {
	return &EarthquakeHandler{
		earthquakeService: earthquakeService,
	}
}

func (eh *EarthquakeHandler) GetAllEarthquakes(w http.ResponseWriter, r *http.Request) {
	earthquakes, err := eh.earthquakeService.GetAllEarthquakes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(earthquakes)
}

func (eh *EarthquakeHandler) GetEarthquakeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid earthquake ID", http.StatusBadRequest)
		return
	}

	earthquake, err := eh.earthquakeService.GetEarthquakeByID(id)
	if err != nil {
		http.Error(w, "Earthquake not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(earthquake)
}

func (eh *EarthquakeHandler) CreateEarthquake(w http.ResponseWriter, r *http.Request) {
	var earthquake domain.Earthquake
	if err := json.NewDecoder(r.Body).Decode(&earthquake); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := eh.earthquakeService.CreateEarthquake(earthquake)
	if err != nil {
		http.Error(w, "Error creating earthquake", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(earthquake)
}

func (eh *EarthquakeHandler) UpdateEarthquake(w http.ResponseWriter, r *http.Request) {
	var earthquake domain.Earthquake
	if err := json.NewDecoder(r.Body).Decode(&earthquake); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	earthquake.Id, _ = strconv.Atoi(params["id"])

	err := eh.earthquakeService.UpdateEarthquake(earthquake)
	if err != nil {
		http.Error(w, "Error updating earthquake", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (eh *EarthquakeHandler) DeleteEarthquake(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid earthquake ID", http.StatusBadRequest)
		return
	}

	err = eh.earthquakeService.DeleteEarthquake(id)
	if err != nil {
		http.Error(w, "Error deleting earthquake", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
