package main

import (
	"backend/api/handlers"
	"backend/api/routes"
	"backend/internal/repository"
	"backend/internal/service"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Veritabanına bağlanırken hata oluştu:", err)
	}
	defer db.Close()

	// create earthquake repository, service and handler
	earthquakeRepo := repository.NewEarthquakeRepository(db)
	earthquakeService := service.NewEarthquakeService(earthquakeRepo)
	earthquakeHandler := handlers.NewEarthquakeHandler(earthquakeService)

	// router setup
	router := mux.NewRouter()
	earthquakeRoutes := routes.SetupEarthquakeRoutes(earthquakeHandler)
	router.PathPrefix("/").Handler(earthquakeRoutes)

	// cors middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // working address of the frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
