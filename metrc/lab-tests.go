package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type LabTestType struct {
	ID                   int         `json:"Id"`
	Name                 string      `json:"Name"`
	RequiresTestResult   bool        `json:"RequiresTestResult"`
	InformationalOnly    bool        `json:"InformationalOnly"`
	AlwaysPasses         bool        `json:"AlwaysPasses"`
	LabTestResultMode    int         `json:"LabTestResultMode"`
	LabTestResultMinimum interface{} `json:"LabTestResultMinimum"`
	LabTestResultMaximum interface{} `json:"LabTestResultMaximum"`
	DependencyMode       int         `json:"DependencyMode"`
}

type LabTestResult struct {
	PackageID                int         `json:"PackageId"`
	LabTestResultID          int         `json:"LabTestResultId"`
	LabFacilityLicenseNumber string      `json:"LabFacilityLicenseNumber"`
	LabFacilityName          string      `json:"LabFacilityName"`
	SourcePackageLabel       string      `json:"SourcePackageLabel"`
	ProductName              string      `json:"ProductName"`
	ProductCategoryName      string      `json:"ProductCategoryName"`
	TestPerformedDate        string      `json:"TestPerformedDate"`
	OverallPassed            bool        `json:"OverallPassed"`
	RevokedDate              interface{} `json:"RevokedDate"`
	ResultReleased           bool        `json:"ResultReleased"`
	ResultReleaseDateTime    time.Time   `json:"ResultReleaseDateTime"`
	TestTypeName             string      `json:"TestTypeName"`
	TestPassed               bool        `json:"TestPassed"`
	TestResultLevel          float64     `json:"TestResultLevel"`
	TestComment              string      `json:"TestComment"`
	TestInformationalOnly    bool        `json:"TestInformationalOnly"`
	LabTestDetailRevokedDate interface{} `json:"LabTestDetailRevokedDate"`
}

type LabTestService struct {
	sling  *sling.Sling
	config *Config
}

func NewLabTestService(sling *sling.Sling, config *Config) *LabTestService {
	return &LabTestService{sling: sling.Path("labtests/"), config: config}
}

func (s *LabTestService) States(userApiKey string) ([]string, *http.Response, error) {
	return nil, nil, nil
}

func (s *LabTestService) Types(userApiKey string) ([]LabTestType, *http.Response, error) {
	return nil, nil, nil
}

type LabTestResultParams struct {
	PackageID     string `json:"packageId"`
	LicenseNumber string `json:"licenseNumber"`
}

func (s *LabTestService) Results(userApiKey string, params *LabTestResultParams) ([]LabTestResult, *http.Response, error) {
	return nil, nil, nil
}
