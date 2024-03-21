package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Earthquake struct {
	Id        int     `json:"id"`
	Latitude  string  `json:"Latitude"`
	Longitude string  `json:"Longitude"`
	Magnitude float64 `json:"Magnitude"`
	Time      string  `json:"Time"`
}

func main() {
	// connect to the postgres db
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// create a new table if it doesn't exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS earthquakes (id SERIAL PRIMARY KEY, Latitude TEXT, Longitude TEXT, Magnitude FLOAT, Time TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// create a new router
	router := mux.NewRouter()
	router.HandleFunc("/earthquakes", getEarthquakes(db)).Methods("GET")
	router.HandleFunc("/earthquakes/{id}", getEarthquake(db)).Methods("GET")
	router.HandleFunc("/earthquakes", createEarthquake(db)).Methods("POST")
	router.HandleFunc("/earthquakes/{id}", updateEarthquake(db)).Methods("PUT")
	router.HandleFunc("/earthquakes/{id}", deleteEarthquake(db)).Methods("DELETE")

	// start the server
	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))

}

func jsonContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

// getEarthquakes returns a handler function that returns a list of earthquakes
func getEarthquakes(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM earthquakes")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		earthquakes := []Earthquake{}
		for rows.Next() {
			var eq Earthquake
			if err := rows.Scan(&eq.Id, &eq.Latitude, &eq.Longitude, &eq.Magnitude); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			earthquakes = append(earthquakes, eq)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(earthquakes)
	}
}

// getEarthquake returns a handler function that returns a single earthquake
func getEarthquake(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var eq Earthquake
		err := db.QueryRow("SELECT * FROM earthquakes WHERE id = $1", id).Scan(&eq.Id, &eq.Latitude, &eq.Longitude, &eq.Magnitude)
		if err != nil {
			//TODO: handle error
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(eq)
	}
}

// createEarthquake returns a handler function that creates a new earthquake
func createEarthquake(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var eq Earthquake
		if err := json.NewDecoder(r.Body).Decode(&eq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := db.QueryRow("INSERT INTO earthquakes(Latitude, Longitude, Magnitude) VALUES($1, $2, $3) RETURNING id", eq.Latitude, eq.Longitude, eq.Magnitude).Scan(&eq.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(eq)
	}
}

// updateEarthquake returns a handler function that updates an existing earthquake
func updateEarthquake(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var eq Earthquake
		if err := json.NewDecoder(r.Body).Decode(&eq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE earthquakes SET Latitude = $1, Longitude = $2, Magnitude = $3 WHERE id = $4", eq.Latitude, eq.Longitude, eq.Magnitude, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// deleteEarthquake returns a handler function that deletes an existing earthquake
func deleteEarthquake(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("DELETE FROM earthquakes WHERE id = $1", id)
		if err != nil {
			//TODO: handle error
			w.WriteHeader(http.StatusNotFound)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
