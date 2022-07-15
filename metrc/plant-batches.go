package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type PlantBatch struct {
	ID                    int           `json:"Id"`
	Name                  string        `json:"Name"`
	PlantBatchTypeID      int           `json:"PlantBatchTypeId"`
	PlantBatchTypeName    string        `json:"PlantBatchTypeName"`
	LocationID            interface{}   `json:"LocationId"`
	LocationName          interface{}   `json:"LocationName"`
	LocationTypeName      interface{}   `json:"LocationTypeName"`
	StrainID              int           `json:"StrainId"`
	StrainName            string        `json:"StrainName"`
	PatientLicenseNumber  interface{}   `json:"PatientLicenseNumber"`
	UntrackedCount        int           `json:"UntrackedCount"`
	TrackedCount          int           `json:"TrackedCount"`
	PackagedCount         int           `json:"PackagedCount"`
	HarvestedCount        int           `json:"HarvestedCount"`
	DestroyedCount        int           `json:"DestroyedCount"`
	SourcePackageID       interface{}   `json:"SourcePackageId"`
	SourcePackageLabel    interface{}   `json:"SourcePackageLabel"`
	SourcePlantID         interface{}   `json:"SourcePlantId"`
	SourcePlantLabel      interface{}   `json:"SourcePlantLabel"`
	SourcePlantBatchIds   []interface{} `json:"SourcePlantBatchIds"`
	SourcePlantBatchNames interface{}   `json:"SourcePlantBatchNames"`
	MultiPlantBatch       bool          `json:"MultiPlantBatch"`
	SourcePlantBatchID    int           `json:"SourcePlantBatchId"`
	SourcePlantBatchName  interface{}   `json:"SourcePlantBatchName"`
	PlantedDate           string        `json:"PlantedDate"`
	LastModified          time.Time     `json:"LastModified"`
}

type PlantBatchService struct {
	sling  *sling.Sling
	config *Config
}

func NewPlantBatchService(sling *sling.Sling, config *Config) *PlantBatchService {
	return &PlantBatchService{sling: sling.Path("plantbatches/"), config: config}
}

func (s *PlantBatchService) GetByID(userApiKey string, id int, licenseNumber string) (*PlantBatch, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantBatchService) List(userApiKey string, status string) ([]PlantBatch, *http.Response, error) {
	return nil, nil, nil
}

func (s *PlantBatchService) Types(userApiKey string) ([]string, *http.Response, error) {
	return nil, nil, nil
}
