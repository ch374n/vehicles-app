package models

type VehicleType struct {
	IsPrimary bool   `json:"IsPrimary"`
	Name      string `json:"Name"`
}

type Manufacturer struct {
	Country        string        `json:"Country"`
	MfrCommonName  string        `json:"Mfr_CommonName"`
	MfrID          int           `json:"Mfr_ID"`
	MfrName        string        `json:"Mfr_Name"`
	VehicleTypes   []VehicleType `json:"VehicleTypes"`
}