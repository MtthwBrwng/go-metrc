package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Strain struct {
	ID               int         `json:"Id"`
	Name             string      `json:"Name"`
	TestingStatus    string      `json:"TestingStatus"`
	ThcLevel         interface{} `json:"ThcLevel"`
	CbdLevel         interface{} `json:"CbdLevel"`
	IndicaPercentage float64     `json:"IndicaPercentage"`
	SativaPercentage float64     `json:"SativaPercentage"`
	IsUsed           bool        `json:"IsUsed"`
	Genetics         string      `json:"Genetics"`
}

type StrainService struct {
	sling  *sling.Sling
	config *Config
}

func NewStrainService(sling *sling.Sling, config *Config) *StrainService {
	return &StrainService{sling: sling.Path("strains/"), config: config}
}

func (s *StrainService) GetByID(userApiKey string) (*Strain, *http.Response, error) {
	return nil, nil, nil
}

func (s *StrainService) List(userApiKey string) ([]Strain, *http.Response, error) {
	return nil, nil, nil
}