package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
)

type UnitOfMeasure struct {
	QuantityType string `json:"QuantityType"`
	Name         string `json:"Name"`
	Abbreviation string `json:"Abbreviation"`
}

type UnitsOfMeasureService struct {
	sling  *sling.Sling
	config *Config
}

func NewUnitsOfMeasureService(sling *sling.Sling, config *Config) *UnitsOfMeasureService {
	return &UnitsOfMeasureService{sling: sling.Path("unitsofmeasure/"), config: config}
}

func (s *UnitsOfMeasureService) List(userApiKey string) ([]UnitOfMeasure, *http.Response, error) {
	return nil, nil, nil
}
