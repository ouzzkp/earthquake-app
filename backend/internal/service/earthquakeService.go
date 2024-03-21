package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type EarthquakeService interface {
	GetAllEarthquakes() ([]domain.Earthquake, error)
	GetEarthquakeByID(id int) (domain.Earthquake, error)
	CreateEarthquake(earthquake domain.Earthquake) error
	UpdateEarthquake(earthquake domain.Earthquake) error
	DeleteEarthquake(id int) error
}

type earthquakeService struct {
	repo repository.EarthquakeRepository
}

func NewEarthquakeService(repo repository.EarthquakeRepository) EarthquakeService {
	return &earthquakeService{repo: repo}
}

func (s *earthquakeService) CreateEarthquake(earthquake domain.Earthquake) error {
	return s.repo.CreateEarthquake(earthquake)
}

func (s *earthquakeService) UpdateEarthquake(earthquake domain.Earthquake) error {
	return s.repo.UpdateEarthquake(earthquake)
}

func (s *earthquakeService) DeleteEarthquake(id int) error {
	return s.repo.DeleteEarthquake(id)
}

func (s *earthquakeService) GetAllEarthquakes() ([]domain.Earthquake, error) {
	return s.repo.GetAllEarthquakes()
}

func (s *earthquakeService) GetEarthquakeByID(id int) (domain.Earthquake, error) {
	return s.repo.GetEarthquakeByID(id)
}
