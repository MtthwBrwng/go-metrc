package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Facility struct {
	HireDate              string        `json:"HireDate"`
	IsOwner               bool          `json:"IsOwner"`
	IsManager             bool          `json:"IsManager"`
	Occupations           []interface{} `json:"Occupations"`
	Name                  string        `json:"Name"`
	Alias                 string        `json:"Alias"`
	DisplayName           string        `json:"DisplayName"`
	CredentialedDate      string        `json:"CredentialedDate"`
	SupportActivationDate interface{}   `json:"SupportActivationDate"`
	SupportExpirationDate interface{}   `json:"SupportExpirationDate"`
	SupportLastPaidDate   interface{}   `json:"SupportLastPaidDate"`
	FacilityType          interface{}   `json:"FacilityType"`
	License               struct {
		Number      string `json:"Number"`
		StartDate   string `json:"StartDate"`
		EndDate     string `json:"EndDate"`
		LicenseType string `json:"LicenseType"`
	} `json:"License"`
}

type FacilityService struct {
	sling  *sling.Sling
	config *Config
}

func NewFacilityService(sling *sling.Sling, config *Config) *FacilityService {
	return &FacilityService{sling: sling.Path("facilities/"), config: config}
}

func (s *FacilityService) List(userApiKey string) (*[]Facility, *http.Response, error) {
	list := new([]Facility)
	resp, err := s.sling.New().SetBasicAuth(s.config.APIKey, userApiKey).Get("v1/").Receive(list, nil)
	return list, resp, err
}
