package models

type VehicleType struct {
	IsPrimary bool   `json:"IsPrimary" bson:"isprimary"`
	Name      string `json:"Name" bson:"name"`
}

type Manufacturer struct {
	Country       string        `json:"Country" bson:"country"`
	MfrCommonName string        `json:"Mfr_CommonName" bson:"mfrcommonname"`
	MfrID         int           `json:"Mfr_ID" bson:"mfrid"`
	MfrName       string        `json:"Mfr_Name" bson:"mfrname"`
	VehicleTypes  []VehicleType `json:"VehicleTypes" bson:"vehicletypes"`
}
