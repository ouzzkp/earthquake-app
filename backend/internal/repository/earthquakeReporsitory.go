package repository

import (
	"database/sql"
	"fmt"

	"github.com/ouzzkp/earthquake-app/backend/internal/domain"
)

// EarthquakeRepository defines the interface for earthquake data access
type EarthquakeRepository interface {
	CreateEarthquake(earthquake domain.Earthquake) error
	UpdateEarthquake(earthquake domain.Earthquake) error
	DeleteEarthquake(id int) error
	GetAllEarthquakes() ([]domain.Earthquake, error)
	GetEarthquakeByID(id int) (domain.Earthquake, error)
}

type earthquakeRepository struct {
	db *sql.DB
}

// NewEarthquakeRepository creates a new instance of earthquakeRepository
func NewEarthquakeRepository(db *sql.DB) EarthquakeRepository {
	return &earthquakeRepository{
		db: db,
	}
}

func (r *earthquakeRepository) CreateEarthquake(eq domain.Earthquake) error {
	query := `INSERT INTO earthquakes (Latitude, Longitude, Magnitude, Time) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, eq.Latitude, eq.Longitude, eq.Magnitude, eq.Time).Scan(&eq.Id)
	if err != nil {
		return fmt.Errorf("CreateEarthquake error: %w", err)
	}
	return nil
}

func (r *earthquakeRepository) UpdateEarthquake(eq domain.Earthquake) error {
	query := `UPDATE earthquakes SET Latitude = $1, Longitude = $2, Magnitude = $3, Time = $4 WHERE id = $5`
	_, err := r.db.Exec(query, eq.Latitude, eq.Longitude, eq.Magnitude, eq.Time, eq.Id)
	if err != nil {
		return fmt.Errorf("UpdateEarthquake error: %w", err)
	}
	return nil
}

func (r *earthquakeRepository) DeleteEarthquake(id int) error {
	query := `DELETE FROM earthquakes WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("DeleteEarthquake error: %w", err)
	}
	return nil
}

func (r *earthquakeRepository) GetAllEarthquakes() ([]domain.Earthquake, error) {
	query := `SELECT id, Latitude, Longitude, Magnitude, Time FROM earthquakes`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("GetAllEarthquakes error: %w", err)
	}
	defer rows.Close()

	var earthquakes []domain.Earthquake
	for rows.Next() {
		var eq domain.Earthquake
		if err := rows.Scan(&eq.Id, &eq.Latitude, &eq.Longitude, &eq.Magnitude, &eq.Time); err != nil {
			return nil, fmt.Errorf("Scan error: %w", err)
		}
		earthquakes = append(earthquakes, eq)
	}
	return earthquakes, nil
}

func (r *earthquakeRepository) GetEarthquakeByID(id int) (domain.Earthquake, error) {
	query := `SELECT id, Latitude, Longitude, Magnitude, Time FROM earthquakes WHERE id = $1`
	var eq domain.Earthquake
	err := r.db.QueryRow(query, id).Scan(&eq.Id, &eq.Latitude, &eq.Longitude, &eq.Magnitude, &eq.Time)
	if err != nil {
		return domain.Earthquake{}, fmt.Errorf("GetEarthquakeByID error: %w", err)
	}
	return eq, nil
}
