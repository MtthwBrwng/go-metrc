package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
)

type Employee struct {
	FullName string      `json:"FullName"`
	License  interface{} `json:"License"`
}

type EmployeeService struct {
	sling  *sling.Sling
	config *Config
}

func NewEmployeeService(sling *sling.Sling, config *Config) *EmployeeService {
	return &EmployeeService{sling: sling.Path("employees/"), config: config}
}

type EmployeeListParams struct {
	LicenseNumber string `json:"licenseNumber"`
}

func (s *EmployeeService) List(userApiKey string, params *EmployeeListParams) (*[]Employee, *http.Response, error) {
	list := new([]Employee)
	resp, err := s.sling.New().SetBasicAuth(s.config.APIKey, userApiKey).Get("v1/").QueryStruct(params).Receive(list, nil)
	return list, resp, err
}
