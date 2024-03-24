package handlers

import (
	"backend/internal/domain"
	"backend/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type EarthquakeHandler struct {
	earthquakeService service.EarthquakeService
}

var clients = make(map[*websocket.Conn]bool) // WebSocket clients
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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

	// if earthquake magnitude is greater than 5, send it to clients (anomaly detection)
	if earthquake.Magnitude > 5 {
		jsonData, err := json.Marshal(earthquake)
		if err != nil {
			log.Printf("Error marshalling earthquake data: %v", err)
		} else {
			BroadcastToClients(jsonData)
		}
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

func (eh *EarthquakeHandler) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Could not upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error: %v", err)
			delete(clients, conn)
			break
		}
	}
}

func BroadcastToClients(message []byte) {
	for conn := range clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Websocket error: %s", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}
