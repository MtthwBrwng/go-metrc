package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Plant struct {
	ID                                int         `json:"Id"`
	Label                             string      `json:"Label"`
	State                             string      `json:"State"`
	GrowthPhase                       string      `json:"GrowthPhase"`
	PlantBatchID                      int         `json:"PlantBatchId"`
	PlantBatchName                    string      `json:"PlantBatchName"`
	PlantBatchTypeID                  int         `json:"PlantBatchTypeId"`
	PlantBatchTypeName                string      `json:"PlantBatchTypeName"`
	StrainID                          int         `json:"StrainId"`
	StrainName                        string      `json:"StrainName"`
	LocationID                        int         `json:"LocationId"`
	LocationName                      string      `json:"LocationName"`
	LocationTypeName                  interface{} `json:"LocationTypeName"`
	PatientLicenseNumber              interface{} `json:"PatientLicenseNumber"`
	HarvestID                         interface{} `json:"HarvestId"`
	HarvestedUnitOfWeightName         interface{} `json:"HarvestedUnitOfWeightName"`
	HarvestedUnitOfWeightAbbreviation interface{} `json:"HarvestedUnitOfWeightAbbreviation"`
	HarvestedWetWeight                interface{} `json:"HarvestedWetWeight"`
	HarvestCount                      int         `json:"HarvestCount"`
	IsOnHold                          bool        `json:"IsOnHold"`
	PlantedDate                       string      `json:"PlantedDate"`
	VegetativeDate                    string      `json:"VegetativeDate"`
	FloweringDate                     interface{} `json:"FloweringDate"`
	HarvestedDate                     interface{} `json:"HarvestedDate"`
	DestroyedDate                     interface{} `json:"DestroyedDate"`
	DestroyedNote                     interface{} `json:"DestroyedNote"`
	DestroyedByUserName               interface{} `json:"DestroyedByUserName"`
	LastModified                      time.Time   `json:"LastModified"`
}

type PlantAdditive struct {
	AdditiveTypeName      interface{} `json:"AdditiveTypeName"`
	ProductTradeName      string      `json:"ProductTradeName"`
	EpaRegistrationNumber interface{} `json:"EpaRegistrationNumber"`
	ProductSupplier       string      `json:"ProductSupplier"`
	ApplicationDevice     string      `json:"ApplicationDevice"`
	AmountUnitOfMeasure   string      `json:"AmountUnitOfMeasure"`
	TotalAmountApplied    float64     `json:"TotalAmountApplied"`
	PlantBatchID          interface{} `json:"PlantBatchId"`
	PlantBatchName        interface{} `json:"PlantBatchName"`
	PlantCount            int         `json:"PlantCount"`
}

type PlantWasteMethod struct {
	Name string `json:"Name"`
}

type PlantWasteReason struct {
	Name         string `json:"Name"`
	RequiresNote bool   `json:"RequiresNote"`
}

type PlantService struct {
	sling  *sling.Sling
	config *Config
}

func NewPlantService(sling *sling.Sling, config *Config) *PlantService {
	return &PlantService{sling: sling.Path("plants/"), config: config}
}

func (s *PlantService) GetByID(userApiKey string, id int) (*Plant, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantService) GetByLabel(userApiKey string, label string) (*Plant, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantService) List(userApiKey string) ([]Plant, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantService) Additives(userApiKey string) ([]PlantAdditive, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantService) GrowthPhases(userApiKey string) ([]string, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantService) AdditiveTypes(userApiKey string) ([]string, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantService) WasteMethods(userApiKey string) ([]PlantWasteMethod, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantService) WasteReasons(userApiKey string) ([]PlantWasteReason, *http.Response, error) {
	return nil, nil, nil
}
