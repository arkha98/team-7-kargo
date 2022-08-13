package truck

type AddTruck struct {
	LicenseNumber  string `json:"license_number"`
	LicenseType    string `json:"license_type"`
	TruckType      string `json:"truck_type"`
	ProductionYear int    `json:"production_year"`
	STNK           string `json:"stnk"`
	KIR            string `json:"kir"`
}
