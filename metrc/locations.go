package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Location struct {
	ID               int    `json:"Id"`
	Name             string `json:"Name"`
	LocationTypeID   int    `json:"LocationTypeId"`
	LocationTypeName string `json:"LocationTypeName"`
	ForPlantBatches  bool   `json:"ForPlantBatches"`
	ForPlants        bool   `json:"ForPlants"`
	ForHarvests      bool   `json:"ForHarvests"`
	ForPackages      bool   `json:"ForPackages"`
}

type LocationType struct {
	ID              int    `json:"Id"`
	Name            string `json:"Name"`
	ForPlantBatches bool   `json:"ForPlantBatches"`
	ForPlants       bool   `json:"ForPlants"`
	ForHarvests     bool   `json:"ForHarvests"`
	ForPackages     bool   `json:"ForPackages"`
}

type LocationService struct {
	sling  *sling.Sling
	config *Config
}

func NewLocationService(sling *sling.Sling, config *Config) *LocationService {
	return &LocationService{sling: sling.Path("locations/"), config: config}
}

func (s *LocationService) GetByID(userApiKey string, id int) (*Location, *http.Response, error) {
	return nil, nil, nil
}

func (s *LocationService) List(userApiKey string) ([]Location, *http.Response, error) {
	return nil, nil, nil
}

func (s *LocationService) Types(userApiKey string) ([]LocationType, *http.Response, error) {
	return nil, nil, nil
}
