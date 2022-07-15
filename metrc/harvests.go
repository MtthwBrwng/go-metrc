package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Harvest struct {
	ID                     int           `json:"Id"`
	Name                   string        `json:"Name"`
	HarvestType            string        `json:"HarvestType"`
	SourceStrainCount      int           `json:"SourceStrainCount"`
	SourceStrainNames      interface{}   `json:"SourceStrainNames"`
	Strains                []interface{} `json:"Strains"`
	DryingLocationID       int           `json:"DryingLocationId"`
	DryingLocationName     string        `json:"DryingLocationName"`
	DryingLocationTypeName interface{}   `json:"DryingLocationTypeName"`
	PatientLicenseNumber   interface{}   `json:"PatientLicenseNumber"`
	CurrentWeight          float64       `json:"CurrentWeight"`
	TotalWasteWeight       float64       `json:"TotalWasteWeight"`
	PlantCount             int           `json:"PlantCount"`
	TotalWetWeight         float64       `json:"TotalWetWeight"`
	TotalRestoredWeight    float64       `json:"TotalRestoredWeight"`
	PackageCount           int           `json:"PackageCount"`
	TotalPackagedWeight    float64       `json:"TotalPackagedWeight"`
	UnitOfWeightName       string        `json:"UnitOfWeightName"`
	LabTestingState        interface{}   `json:"LabTestingState"`
	LabTestingStateDate    interface{}   `json:"LabTestingStateDate"`
	IsOnHold               bool          `json:"IsOnHold"`
	HarvestStartDate       string        `json:"HarvestStartDate"`
	FinishedDate           interface{}   `json:"FinishedDate"`
	ArchivedDate           interface{}   `json:"ArchivedDate"`
	LastModified           time.Time     `json:"LastModified"`
}

type HarvestWasteType struct {
	Name string `json:"Name"`
}

type HarvestService struct {
	sling  *sling.Sling
	config *Config
}

func NewHarvestService(sling *sling.Sling, config *Config) *HarvestService {
	return &HarvestService{sling: sling.Path("harvests/"), config: config}
}

type HarvestGetByIDParams struct {
	LicenseNumber string `json:"licenseNumber"`
}

func (s *HarvestService) GetByID(userApiKey string, params *HarvestGetByIDParams) (*Harvest, *http.Response, error) {
	return nil, nil, nil
}

type HarvestListParams struct {
	LicenseNumber     string `json:"licenseNumber"`
	LastModifiedStart string `json:"lastModifiedStart"`
	LastModifiedEnd   string `json:"lastModifiedEnd"`
}

func (s *HarvestService) List(userApiKey string, params *HarvestListParams) (*[]Harvest, *http.Response, error) {
	return nil, nil, nil
}

func (s *Harvest) WasteTypes(userApiKey string) (*[]HarvestWasteType, *http.Response, error) {
	return nil, nil, nil
}
