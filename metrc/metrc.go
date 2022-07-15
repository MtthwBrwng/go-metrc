package metrc

import (
	"github.com/dghubble/sling"
)

type Client struct {
	sling          *sling.Sling
	Employees      *EmployeeService
	Facilities     *FacilityService
	Harvests       *HarvestService
	Items          *ItemService
	LabTests       *LabTestService
	Locations      *LocationService
	Packages       *PackageService
	PlantBatches   *PlantBatchService
	Plants         *PlantService
	Sales          *SaleService
	Strains        *StrainService
	Transfers      *TransferService
	UnitsOfMeasure *UnitsOfMeasureService
}

type Config struct {
	Production bool
	APIKey     string
}

func NewClient(config *Config) *Client {
	metrcAPIUrl := "https://sandbox-api-ca.metrc.com/"
	if config.Production {
		metrcAPIUrl = "https://api-ca.metrc.com/"
	}

	base := sling.New().Base(metrcAPIUrl)
	return &Client{
		sling:          base,
		Employees:      NewEmployeeService(base.New(), config),
		Facilities:     NewFacilityService(base.New(), config),
		Harvests:       NewHarvestService(base.New(), config),
		Items:          NewItemService(base.New(), config),
		LabTests:       NewLabTestService(base.New(), config),
		Locations:      NewLocationService(base.New(), config),
		Packages:       NewPackageService(base.New(), config),
		PlantBatches:   NewPlantBatchService(base.New(), config),
		Plants:         NewPlantService(base.New(), config),
		Sales:          NewSaleService(base.New(), config),
		Strains:        NewStrainService(base.New(), config),
		Transfers:      NewTransferService(base.New(), config),
		UnitsOfMeasure: NewUnitsOfMeasureService(base.New(), config),
	}
}
