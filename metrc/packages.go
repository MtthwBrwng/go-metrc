package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Package struct {
	ID                                int         `json:"Id"`
	Label                             string      `json:"Label"`
	PackageType                       string      `json:"PackageType"`
	SourceHarvestCount                int         `json:"SourceHarvestCount"`
	SourcePackageCount                int         `json:"SourcePackageCount"`
	SourceProcessingJobCount          int         `json:"SourceProcessingJobCount"`
	SourceHarvestNames                interface{} `json:"SourceHarvestNames"`
	LocationID                        interface{} `json:"LocationId"`
	LocationName                      interface{} `json:"LocationName"`
	LocationTypeName                  interface{} `json:"LocationTypeName"`
	Quantity                          float64     `json:"Quantity"`
	UnitOfMeasureName                 string      `json:"UnitOfMeasureName"`
	UnitOfMeasureAbbreviation         string      `json:"UnitOfMeasureAbbreviation"`
	PatientLicenseNumber              interface{} `json:"PatientLicenseNumber"`
	ItemFromFacilityLicenseNumber     interface{} `json:"ItemFromFacilityLicenseNumber"`
	ItemFromFacilityName              interface{} `json:"ItemFromFacilityName"`
	Note                              interface{} `json:"Note"`
	PackagedDate                      string      `json:"PackagedDate"`
	InitialLabTestingState            string      `json:"InitialLabTestingState"`
	LabTestingState                   string      `json:"LabTestingState"`
	LabTestingStateDate               string      `json:"LabTestingStateDate"`
	IsProductionBatch                 bool        `json:"IsProductionBatch"`
	ProductionBatchNumber             interface{} `json:"ProductionBatchNumber"`
	SourceProductionBatchNumbers      interface{} `json:"SourceProductionBatchNumbers"`
	IsTradeSample                     bool        `json:"IsTradeSample"`
	IsTradeSamplePersistent           bool        `json:"IsTradeSamplePersistent"`
	SourcePackageIsTradeSample        bool        `json:"SourcePackageIsTradeSample"`
	IsDonation                        bool        `json:"IsDonation"`
	IsDonationPersistent              bool        `json:"IsDonationPersistent"`
	SourcePackageIsDonation           bool        `json:"SourcePackageIsDonation"`
	IsTestingSample                   bool        `json:"IsTestingSample"`
	IsProcessValidationTestingSample  bool        `json:"IsProcessValidationTestingSample"`
	ProductRequiresRemediation        bool        `json:"ProductRequiresRemediation"`
	ContainsRemediatedProduct         bool        `json:"ContainsRemediatedProduct"`
	RemediationDate                   interface{} `json:"RemediationDate"`
	ReceivedDateTime                  interface{} `json:"ReceivedDateTime"`
	ReceivedFromManifestNumber        interface{} `json:"ReceivedFromManifestNumber"`
	ReceivedFromFacilityLicenseNumber interface{} `json:"ReceivedFromFacilityLicenseNumber"`
	ReceivedFromFacilityName          interface{} `json:"ReceivedFromFacilityName"`
	IsOnHold                          bool        `json:"IsOnHold"`
	ArchivedDate                      interface{} `json:"ArchivedDate"`
	FinishedDate                      interface{} `json:"FinishedDate"`
	LastModified                      time.Time   `json:"LastModified"`
	Item                              struct {
		ID                                  int         `json:"Id"`
		Name                                string      `json:"Name"`
		ProductCategoryName                 string      `json:"ProductCategoryName"`
		ProductCategoryType                 int         `json:"ProductCategoryType"`
		QuantityType                        int         `json:"QuantityType"`
		DefaultLabTestingState              int         `json:"DefaultLabTestingState"`
		UnitOfMeasureName                   interface{} `json:"UnitOfMeasureName"`
		ApprovalStatus                      int         `json:"ApprovalStatus"`
		ApprovalStatusDateTime              time.Time   `json:"ApprovalStatusDateTime"`
		StrainID                            interface{} `json:"StrainId"`
		StrainName                          interface{} `json:"StrainName"`
		AdministrationMethod                interface{} `json:"AdministrationMethod"`
		UnitCbdPercent                      interface{} `json:"UnitCbdPercent"`
		UnitCbdContent                      interface{} `json:"UnitCbdContent"`
		UnitCbdContentUnitOfMeasureName     interface{} `json:"UnitCbdContentUnitOfMeasureName"`
		UnitCbdContentDose                  interface{} `json:"UnitCbdContentDose"`
		UnitCbdContentDoseUnitOfMeasureName interface{} `json:"UnitCbdContentDoseUnitOfMeasureName"`
		UnitThcPercent                      interface{} `json:"UnitThcPercent"`
		UnitThcContent                      interface{} `json:"UnitThcContent"`
		UnitThcContentUnitOfMeasureName     interface{} `json:"UnitThcContentUnitOfMeasureName"`
		UnitThcContentDose                  interface{} `json:"UnitThcContentDose"`
		UnitThcContentDoseUnitOfMeasureName interface{} `json:"UnitThcContentDoseUnitOfMeasureName"`
		UnitVolume                          interface{} `json:"UnitVolume"`
		UnitVolumeUnitOfMeasureName         interface{} `json:"UnitVolumeUnitOfMeasureName"`
		UnitWeight                          interface{} `json:"UnitWeight"`
		UnitWeightUnitOfMeasureName         interface{} `json:"UnitWeightUnitOfMeasureName"`
		ServingSize                         interface{} `json:"ServingSize"`
		SupplyDurationDays                  interface{} `json:"SupplyDurationDays"`
		NumberOfDoses                       interface{} `json:"NumberOfDoses"`
		UnitQuantity                        interface{} `json:"UnitQuantity"`
		UnitQuantityUnitOfMeasureName       interface{} `json:"UnitQuantityUnitOfMeasureName"`
		PublicIngredients                   interface{} `json:"PublicIngredients"`
		Description                         interface{} `json:"Description"`
		IsUsed                              bool        `json:"IsUsed"`
	} `json:"Item"`
}

type PackageAdjustReason struct {
	Name         string `json:"Name"`
	RequiresNote bool   `json:"RequiresNote"`
}

type PackageService struct {
	sling  *sling.Sling
	config *Config
}

func NewPackageService(sling *sling.Sling, config *Config) *PackageService {
	return &PackageService{sling: sling.Path("packages/"), config: config}
}

func (s *PackageService) GetByID(userApiKey string, id int) (*Package, *http.Response, error) {
	return nil, nil, nil
}

func (s *PackageService) GetByLabel(userApiKey string, label string) (*Package, *http.Response, error) {
	return nil, nil, nil
}

func (s *PackageService) List(userApiKey string) ([]Package, *http.Response, error) {
	return nil, nil, nil
}

func (s *PackageService) Types(userApiKey string) ([]string, *http.Response, error) {
	return nil, nil, nil
}

func (s *PackageService) AdjustReasons(userApiKey string) ([]PackageAdjustReason, *http.Response, error) {
	return nil, nil, nil
}
