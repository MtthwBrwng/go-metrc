package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Transfer struct {
	ID                                 int         `json:"Id"`
	ManifestNumber                     string      `json:"ManifestNumber"`
	ShipmentLicenseType                int         `json:"ShipmentLicenseType"`
	ShipperFacilityLicenseNumber       string      `json:"ShipperFacilityLicenseNumber"`
	ShipperFacilityName                string      `json:"ShipperFacilityName"`
	Name                               interface{} `json:"Name"`
	TransporterFacilityLicenseNumber   string      `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName            string      `json:"TransporterFacilityName"`
	DriverName                         string      `json:"DriverName"`
	DriverOccupationalLicenseNumber    string      `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber         string      `json:"DriverVehicleLicenseNumber"`
	VehicleMake                        string      `json:"VehicleMake"`
	VehicleModel                       string      `json:"VehicleModel"`
	VehicleLicensePlateNumber          string      `json:"VehicleLicensePlateNumber"`
	DeliveryCount                      int         `json:"DeliveryCount"`
	ReceivedDeliveryCount              int         `json:"ReceivedDeliveryCount"`
	PackageCount                       int         `json:"PackageCount"`
	ReceivedPackageCount               int         `json:"ReceivedPackageCount"`
	ContainsPlantPackage               bool        `json:"ContainsPlantPackage"`
	ContainsProductPackage             bool        `json:"ContainsProductPackage"`
	ContainsTradeSample                bool        `json:"ContainsTradeSample"`
	ContainsDonation                   bool        `json:"ContainsDonation"`
	ContainsTestingSample              bool        `json:"ContainsTestingSample"`
	ContainsProductRequiresRemediation bool        `json:"ContainsProductRequiresRemediation"`
	ContainsRemediatedProductPackage   bool        `json:"ContainsRemediatedProductPackage"`
	CreatedDateTime                    string      `json:"CreatedDateTime"`
	CreatedByUserName                  interface{} `json:"CreatedByUserName"`
	LastModified                       time.Time   `json:"LastModified"`
	DeliveryID                         int         `json:"DeliveryId"`
	RecipientFacilityLicenseNumber     string      `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName              string      `json:"RecipientFacilityName"`
	ShipmentTypeName                   string      `json:"ShipmentTypeName"`
	ShipmentTransactionType            string      `json:"ShipmentTransactionType"`
	EstimatedDepartureDateTime         string      `json:"EstimatedDepartureDateTime"`
	ActualDepartureDateTime            interface{} `json:"ActualDepartureDateTime"`
	EstimatedArrivalDateTime           string      `json:"EstimatedArrivalDateTime"`
	ActualArrivalDateTime              interface{} `json:"ActualArrivalDateTime"`
	DeliveryPackageCount               int         `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount       int         `json:"DeliveryReceivedPackageCount"`
	ReceivedDateTime                   string      `json:"ReceivedDateTime"`
	EstimatedReturnDepartureDateTime   interface{} `json:"EstimatedReturnDepartureDateTime"`
	ActualReturnDepartureDateTime      interface{} `json:"ActualReturnDepartureDateTime"`
	EstimatedReturnArrivalDateTime     interface{} `json:"EstimatedReturnArrivalDateTime"`
	ActualReturnArrivalDateTime        interface{} `json:"ActualReturnArrivalDateTime"`
}

type TransferDelivery struct {
	ID                               int         `json:"Id"`
	RecipientFacilityLicenseNumber   string      `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName            string      `json:"RecipientFacilityName"`
	ShipmentTypeName                 string      `json:"ShipmentTypeName"`
	ShipmentTransactionType          string      `json:"ShipmentTransactionType"`
	EstimatedDepartureDateTime       string      `json:"EstimatedDepartureDateTime"`
	ActualDepartureDateTime          interface{} `json:"ActualDepartureDateTime"`
	EstimatedArrivalDateTime         string      `json:"EstimatedArrivalDateTime"`
	ActualArrivalDateTime            interface{} `json:"ActualArrivalDateTime"`
	GrossWeight                      interface{} `json:"GrossWeight"`
	GrossUnitOfWeightID              interface{} `json:"GrossUnitOfWeightId"`
	GrossUnitOfWeightName            interface{} `json:"GrossUnitOfWeightName"`
	PlannedRoute                     string      `json:"PlannedRoute"`
	DeliveryPackageCount             int         `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount     int         `json:"DeliveryReceivedPackageCount"`
	ReceivedDateTime                 string      `json:"ReceivedDateTime"`
	EstimatedReturnDepartureDateTime interface{} `json:"EstimatedReturnDepartureDateTime"`
	ActualReturnDepartureDateTime    interface{} `json:"ActualReturnDepartureDateTime"`
	EstimatedReturnArrivalDateTime   interface{} `json:"EstimatedReturnArrivalDateTime"`
	ActualReturnArrivalDateTime      interface{} `json:"ActualReturnArrivalDateTime"`
	RejectedPackagesReturned         bool        `json:"RejectedPackagesReturned"`
}

type TransferDeliveryTransporter struct {
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName          string `json:"TransporterFacilityName"`
	TransporterDirection             int    `json:"TransporterDirection"`
}

type TransferDeliveryTransporterDetails struct {
	DriverName                      string      `json:"DriverName"`
	DriverOccupationalLicenseNumber string      `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber      string      `json:"DriverVehicleLicenseNumber"`
	VehicleMake                     string      `json:"VehicleMake"`
	VehicleModel                    string      `json:"VehicleModel"`
	VehicleLicensePlateNumber       string      `json:"VehicleLicensePlateNumber"`
	ActualDriverStartDateTime       interface{} `json:"ActualDriverStartDateTime"`
}

type TransferDeliveryPackage struct {
	PackageID                               int         `json:"PackageId"`
	PackageLabel                            string      `json:"PackageLabel"`
	PackageType                             string      `json:"PackageType"`
	SourceHarvestNames                      interface{} `json:"SourceHarvestNames"`
	SourcePackageLabels                     interface{} `json:"SourcePackageLabels"`
	ItemID                                  int         `json:"ItemId"`
	ProductName                             string      `json:"ProductName"`
	ItemName                                string      `json:"ItemName"`
	ProductCategoryName                     string      `json:"ProductCategoryName"`
	ItemCategoryName                        string      `json:"ItemCategoryName"`
	ItemStrainName                          interface{} `json:"ItemStrainName"`
	ItemBrandName                           interface{} `json:"ItemBrandName"`
	ItemUnitCbdPercent                      interface{} `json:"ItemUnitCbdPercent"`
	ItemUnitCbdContent                      interface{} `json:"ItemUnitCbdContent"`
	ItemUnitCbdContentUnitOfMeasureName     interface{} `json:"ItemUnitCbdContentUnitOfMeasureName"`
	ItemUnitCbdContentDose                  interface{} `json:"ItemUnitCbdContentDose"`
	ItemUnitCbdContentDoseUnitOfMeasureName interface{} `json:"ItemUnitCbdContentDoseUnitOfMeasureName"`
	ItemUnitThcPercent                      interface{} `json:"ItemUnitThcPercent"`
	ItemUnitThcContent                      interface{} `json:"ItemUnitThcContent"`
	ItemUnitThcContentUnitOfMeasureName     interface{} `json:"ItemUnitThcContentUnitOfMeasureName"`
	ItemUnitThcContentDose                  interface{} `json:"ItemUnitThcContentDose"`
	ItemUnitThcContentDoseUnitOfMeasureName interface{} `json:"ItemUnitThcContentDoseUnitOfMeasureName"`
	ItemUnitVolume                          interface{} `json:"ItemUnitVolume"`
	ItemUnitVolumeUnitOfMeasureName         interface{} `json:"ItemUnitVolumeUnitOfMeasureName"`
	ItemUnitWeight                          interface{} `json:"ItemUnitWeight"`
	ItemUnitWeightUnitOfMeasureName         interface{} `json:"ItemUnitWeightUnitOfMeasureName"`
	ItemServingSize                         interface{} `json:"ItemServingSize"`
	ItemSupplyDurationDays                  interface{} `json:"ItemSupplyDurationDays"`
	ItemUnitQuantity                        interface{} `json:"ItemUnitQuantity"`
	ItemUnitQuantityUnitOfMeasureName       interface{} `json:"ItemUnitQuantityUnitOfMeasureName"`
	LabTestingState                         string      `json:"LabTestingState"`
	ProductionBatchNumber                   interface{} `json:"ProductionBatchNumber"`
	IsTradeSample                           bool        `json:"IsTradeSample"`
	IsTradeSamplePersistent                 bool        `json:"IsTradeSamplePersistent"`
	SourcePackageIsTradeSample              bool        `json:"SourcePackageIsTradeSample"`
	IsDonation                              bool        `json:"IsDonation"`
	SourcePackageIsDonation                 bool        `json:"SourcePackageIsDonation"`
	IsTestingSample                         bool        `json:"IsTestingSample"`
	ProductRequiresRemediation              bool        `json:"ProductRequiresRemediation"`
	ContainsRemediatedProduct               bool        `json:"ContainsRemediatedProduct"`
	RemediationDate                         interface{} `json:"RemediationDate"`
	ShipmentPackageState                    string      `json:"ShipmentPackageState"`
	ShippedQuantity                         float64     `json:"ShippedQuantity"`
	ShippedUnitOfMeasureName                string      `json:"ShippedUnitOfMeasureName"`
	GrossUnitOfWeightName                   interface{} `json:"GrossUnitOfWeightName"`
	ReceivedQuantity                        float64     `json:"ReceivedQuantity"`
	ReceivedUnitOfMeasureName               string      `json:"ReceivedUnitOfMeasureName"`
}

type TransferDeliveryWholesalePackage struct {
	PackageID              int         `json:"PackageId"`
	PackageLabel           string      `json:"PackageLabel"`
	ShipperWholesalePrice  interface{} `json:"ShipperWholesalePrice"`
	ReceiverWholesalePrice interface{} `json:"ReceiverWholesalePrice"`
}

type TransferDeliveryRequiredLabTestPackage struct {
	PackageID        int    `json:"PackageId"`
	LabTestBatchID   int    `json:"LabTestBatchId"`
	LabTestBatchName string `json:"LabTestBatchName"`
}

type TransferType struct {
	Name                           string `json:"Name"`
	ForLicensedShipments           bool   `json:"ForLicensedShipments"`
	ForExternalIncomingShipments   bool   `json:"ForExternalIncomingShipments"`
	ForExternalOutgoingShipments   bool   `json:"ForExternalOutgoingShipments"`
	RequiresDestinationGrossWeight bool   `json:"RequiresDestinationGrossWeight"`
	RequiresPackagesGrossWeight    bool   `json:"RequiresPackagesGrossWeight"`
}

type TransferService struct {
	sling  *sling.Sling
	config *Config
}

func NewTransferService(sling *sling.Sling, config *Config) *TransferService {
	return &TransferService{sling: sling.Path("transfers/"), config: config}
}

func (s *TransferService) List(userApiKey string) ([]Transfer, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) Deliveries(userApiKey string) ([]TransferDelivery, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) DeliveryTransporters(userApiKey string) ([]TransferDeliveryTransporter, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) DeliveryTransporterDetails(userApiKey string) ([]TransferDeliveryTransporterDetails, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) DeliveryPackages(userApiKey string) ([]TransferDeliveryPackage, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) DeliveryWholesalePackages(userApiKey string) ([]TransferDeliveryWholesalePackage, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) DeliveryRequiredLabTestPackage(userApiKey string) ([]TransferDeliveryRequiredLabTestPackage, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) DeliveryPackageStates(userApiKey string) ([]string, *http.Response, error) {
	return nil, nil, nil
}

func (s *TransferService) Types(userApiKey string) ([]TransferType, *http.Response, error) {
	return nil, nil, nil
}
