package metrc

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Item struct {
	ID                                  int         `json:"Id"`
	Name                                string      `json:"Name"`
	ProductCategoryName                 string      `json:"ProductCategoryName"`
	ProductCategoryType                 string      `json:"ProductCategoryType"`
	QuantityType                        string      `json:"QuantityType"`
	DefaultLabTestingState              string      `json:"DefaultLabTestingState"`
	UnitOfMeasureName                   string      `json:"UnitOfMeasureName"`
	ApprovalStatus                      string      `json:"ApprovalStatus"`
	ApprovalStatusDateTime              time.Time   `json:"ApprovalStatusDateTime"`
	StrainID                            int         `json:"StrainId"`
	StrainName                          string      `json:"StrainName"`
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
}

type ItemCategory struct {
	Name                         string `json:"Name"`
	ProductCategoryType          string `json:"ProductCategoryType"`
	QuantityType                 string `json:"QuantityType"`
	RequiresStrain               bool   `json:"RequiresStrain"`
	RequiresItemBrand            bool   `json:"RequiresItemBrand"`
	RequiresAdministrationMethod bool   `json:"RequiresAdministrationMethod"`
	RequiresUnitCbdPercent       bool   `json:"RequiresUnitCbdPercent"`
	RequiresUnitCbdContent       bool   `json:"RequiresUnitCbdContent"`
	RequiresUnitCbdContentDose   bool   `json:"RequiresUnitCbdContentDose"`
	RequiresUnitThcPercent       bool   `json:"RequiresUnitThcPercent"`
	RequiresUnitThcContent       bool   `json:"RequiresUnitThcContent"`
	RequiresUnitThcContentDose   bool   `json:"RequiresUnitThcContentDose"`
	RequiresUnitVolume           bool   `json:"RequiresUnitVolume"`
	RequiresUnitWeight           bool   `json:"RequiresUnitWeight"`
	RequiresServingSize          bool   `json:"RequiresServingSize"`
	RequiresSupplyDurationDays   bool   `json:"RequiresSupplyDurationDays"`
	RequiresNumberOfDoses        bool   `json:"RequiresNumberOfDoses"`
	RequiresPublicIngredients    bool   `json:"RequiresPublicIngredients"`
	RequiresDescription          bool   `json:"RequiresDescription"`
	RequiresProductPhotos        int    `json:"RequiresProductPhotos"`
	RequiresLabelPhotos          int    `json:"RequiresLabelPhotos"`
	RequiresPackagingPhotos      int    `json:"RequiresPackagingPhotos"`
	CanContainSeeds              bool   `json:"CanContainSeeds"`
	CanBeRemediated              bool   `json:"CanBeRemediated"`
	CanBeDestroyed               bool   `json:"CanBeDestroyed"`
}

type ItemBrand struct {
	Name   string `json:"Name"`
	Status string `json:"Status"`
}

type ItemService struct {
	sling  *sling.Sling
	config *Config
}

func NewItemService(sling *sling.Sling, config *Config) *ItemService {
	return &ItemService{sling: sling.Path("items/"), config: config}
}

type ItemGetByIDParams struct {
	LicenseNumber string `json:"licenseNumber"`
}

func (s *ItemService) GetByID(userApiKey string, params *ItemGetByIDParams) (*Item, *http.Response, error) {
	return nil, nil, nil
}

type ItemListParams struct {
	LicenseNumber string `json:"licenseNumber"`
}

func (s *ItemService) List(userApiKey string, params *ItemListParams) ([]Item, *http.Response, error) {
	return nil, nil, nil
}

type ItemCategoriesParams struct {
	LicenseNumber string `json:"licenseNumber"`
}

func (s *ItemService) Categories(userApiKey string, params *ItemCategoriesParams) ([]ItemCategory, *http.Response, error) {
	return nil, nil, nil
}

type ItemBrandParams struct {
	LicenseNumber string `json:"licenseNumber"`
}

func (s *ItemService) Brands(userApiKey string, params *ItemBrandParams) ([]ItemBrand, *http.Response, error) {
	return nil, nil, nil
}
