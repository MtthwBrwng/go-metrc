package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type SaleReceipt struct {
	ID                     int           `json:"Id"`
	ReceiptNumber          interface{}   `json:"ReceiptNumber"`
	SalesDateTime          string        `json:"SalesDateTime"`
	SalesCustomerType      string        `json:"SalesCustomerType"`
	PatientLicenseNumber   interface{}   `json:"PatientLicenseNumber"`
	CaregiverLicenseNumber interface{}   `json:"CaregiverLicenseNumber"`
	IdentificationMethod   interface{}   `json:"IdentificationMethod"`
	TotalPackages          int           `json:"TotalPackages"`
	TotalPrice             float64       `json:"TotalPrice"`
	Transactions           []interface{} `json:"Transactions"`
	IsFinal                bool          `json:"IsFinal"`
	ArchivedDate           interface{}   `json:"ArchivedDate"`
	RecordedDateTime       time.Time     `json:"RecordedDateTime"`
	RecordedByUserName     interface{}   `json:"RecordedByUserName"`
	LastModified           time.Time     `json:"LastModified"`
}

type SaleTransaction struct {
	SalesDate         string  `json:"SalesDate"`
	TotalTransactions int     `json:"TotalTransactions"`
	TotalPackages     int     `json:"TotalPackages"`
	TotalPrice        float64 `json:"TotalPrice"`
}

type SaleTransactionDetails struct {
	PackageID                               int         `json:"PackageId"`
	PackageLabel                            string      `json:"PackageLabel"`
	ProductName                             string      `json:"ProductName"`
	ProductCategoryName                     interface{} `json:"ProductCategoryName"`
	ItemStrainName                          interface{} `json:"ItemStrainName"`
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
	QuantitySold                            float64     `json:"QuantitySold"`
	UnitOfMeasureName                       string      `json:"UnitOfMeasureName"`
	UnitOfMeasureAbbreviation               string      `json:"UnitOfMeasureAbbreviation"`
	UnitThcPercent                          interface{} `json:"UnitThcPercent"`
	UnitThcContent                          interface{} `json:"UnitThcContent"`
	UnitWeight                              interface{} `json:"UnitWeight"`
	TotalPrice                              float64     `json:"TotalPrice"`
	SalesDeliveryState                      interface{} `json:"SalesDeliveryState"`
	ArchivedDate                            interface{} `json:"ArchivedDate"`
	RecordedDateTime                        time.Time   `json:"RecordedDateTime"`
	RecordedByUserName                      interface{} `json:"RecordedByUserName"`
	LastModified                            time.Time   `json:"LastModified"`
}

type SaleService struct {
	sling  *sling.Sling
	config *Config
}

func NewSaleService(sling *sling.Sling, config *Config) *SaleService {
	return &SaleService{sling: sling.Path("sales/"), config: config}
}

func (s *SaleService) CustomerTypes(userApiKey string) ([]string, *http.Response, error) {
	return nil, nil, nil
}

func (s *SaleService) Receipts(userApiKey string) ([]SaleReceipt, *http.Response, error) {
	return nil, nil, nil
}

func (s *SaleService) Receipt(userApiKey string) (*SaleReceipt, *http.Response, error) {
	return nil, nil, nil
}

func (s *SaleService) Transactions(userApiKey string) ([]SaleTransaction, *http.Response, error) {
	return nil, nil, nil
}

func (s *SaleService) TransactionDetails(userApiKey string) ([]SaleTransactionDetails, *http.Response, error) {
	return nil, nil, nil
}
