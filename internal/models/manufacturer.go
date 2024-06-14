package models

type VehicleType struct {
	IsPrimary bool   `json:"IsPrimary" bson:"isPrimary"`
	Name      string `json:"Name" bson:"name"`
}

type Manufacturer struct {
	Country       string        `json:"Country" bson:"country"`
	MfrCommonName string        `json:"Mfr_CommonName" bson:"mfrCommonName"`
	MfrID         int           `json:"Mfr_ID" bson:"mfrID"`
	MfrName       string        `json:"Mfr_Name" bson:"mfrName"`
	VehicleTypes  []VehicleType `json:"VehicleTypes" bson:"vehicleTypes"`
}
