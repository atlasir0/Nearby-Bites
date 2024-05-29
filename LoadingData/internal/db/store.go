package db

import "LoadingData/internal/models"

type Store interface {
	GetPlaces(limit, offset int) ([]models.Place, int, error)
	GetNearestPlaces(lat, lon float64, limit int) ([]models.Place, error)
}
