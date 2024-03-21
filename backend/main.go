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
	// PostgreSQL'e bağlan
	connStr := os.Getenv("DATABASE_URL") // DATABASE_URL ortam değişkeninden alınır
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Veritabanına bağlanırken hata oluştu:", err)
	}
	defer db.Close()

	// EarthquakeRepository, Service ve Handlers'ın oluşturulması
	earthquakeRepo := repository.NewEarthquakeRepository(db)
	earthquakeService := service.NewEarthquakeService(earthquakeRepo)
	earthquakeHandler := handlers.NewEarthquakeHandler(earthquakeService)

	// Routes'ların ayarlanması
	router := mux.NewRouter()
	earthquakeRoutes := routes.SetupEarthquakeRoutes(earthquakeHandler) // Earthquake route'larını al
	router.PathPrefix("/").Handler(earthquakeRoutes)                    // Ana router'a earthquake route'larını ekle

	// CORS middleware'inin eklenmesi
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http:localhost:3000"}, // working address of the frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)

	// HTTP sunucusunun başlatılması
	log.Println("Server port 8080 üzerinde başlatılıyor...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
