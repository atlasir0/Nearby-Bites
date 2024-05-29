package db

import (
	"LoadingData/internal/models"
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8"
)


type ElasticsearchStore struct {
	client *elasticsearch.Client
}

func NewElasticsearchStore() (*ElasticsearchStore, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating Elasticsearch client: %s", err)
	}

	return &ElasticsearchStore{
		client: client,
	}, nil
}


func (s *ElasticsearchStore) GetPlaces(limit, offset int) ([]models.Place, int, error) {
	return nil, 0, errors.New("not implemented")
}

func (s *ElasticsearchStore) GetNearestPlaces(lat, lon float64, limit int) ([]models.Place, error) {
	return nil, errors.New("not implemented")
}


func (s *ElasticsearchStore) IndexPlace(place models.Place) error {
	var buf bytes.Buffer
	id := strconv.Itoa(place.ID)
	if _, err := buf.WriteString(fmt.Sprintf(`{"id": %s, "name": "%s", "location": {"lat": %f, "lon": %f}}`, id, place.Name, place.Location.Lat, place.Location.Lon)); err != nil {
		return err
	}

	res, err := s.client.Index("places", &buf)
	if err != nil {
		return fmt.Errorf("error indexing document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}

	return nil
}
