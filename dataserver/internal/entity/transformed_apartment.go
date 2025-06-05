package entity

import (
	"gorm.io/gorm"
)

type TransformedApartments []TransformedApartment

func (a TransformedApartments) Create() error {
	return db.CreateInBatches(a, 1000).Error
}

type TransformedApartment struct {
	gorm.Model

	UUID string `json:"UUID"`

	SquareMeters   float64 `json:"num__SquareMeters"`
	Rooms          float64 `json:"num__Rooms"`
	Floor          float64 `json:"num__Floor"`
	FloorCount     float64 `json:"num__FloorCount"`
	CentreDistance float64 `json:"num__CentreDistance"`
	PoiCount       float64 `json:"num__PoiCount"`
	Age            float64 `json:"num__Age"`

	HasParkingSpace float64 `json:"bool__HasParkingSpace"`
	HasBalcony      float64 `json:"bool__HasBalcony"`
	HasElevator     float64 `json:"bool__HasElevator"`
	HasSecurity     float64 `json:"bool__HasSecurity"`
	HasStorageRoom  float64 `json:"bool__HasStorageRoom"`
	HasSchool       float64 `json:"bool__HasSchool"`
	HasClinic       float64 `json:"bool__HasClinic"`
	HasPostOffice   float64 `json:"bool__HasPostOffice"`
	HasKindergarten float64 `json:"bool__HasKindergarten"`
	HasRestaurant   float64 `json:"bool__HasRestaurant"`
	HasCollege      float64 `json:"bool__HasCollege"`
	HasPharmacy     float64 `json:"bool__HasPharmacy"`

	CityBialystok   float64 `json:"cat__City_bialystok"`
	CityBydgoszcz   float64 `json:"cat__City_bydgoszcz"`
	CityCzestochowa float64 `json:"cat__City_czestochowa"`
	CityGdansk      float64 `json:"cat__City_gdansk"`
	CityGdynia      float64 `json:"cat__City_gdynia"`
	CityKatowice    float64 `json:"cat__City_katowice"`
	CityKrakow      float64 `json:"cat__City_krakow"`
	CityLodz        float64 `json:"cat__City_lodz"`
	CityLublin      float64 `json:"cat__City_lublin"`
	CityPoznan      float64 `json:"cat__City_poznan"`
	CityRadom       float64 `json:"cat__City_radom"`
	CityRzeszow     float64 `json:"cat__City_rzeszow"`
	CitySzczecin    float64 `json:"cat__City_szczecin"`
	CityWarszawa    float64 `json:"cat__City_warszawa"`
	CityWroclaw     float64 `json:"cat__City_wroclaw"`

	TypeApartmentBuilding float64 `json:"cat__Type_apartmentBuilding"`
	TypeBlockOfFlats      float64 `json:"cat__Type_blockOfFlats"`
	TypeTenement          float64 `json:"cat__Type_tenement"`

	Price int64 `json:"Price"`
}

func (*TransformedApartment) TableName() string {
	return "transformed_apartments"
}

func FindTransformedApartment(uuid string) *TransformedApartment {
	if uuid == "" {
		return nil
	}

	result := &TransformedApartment{}

	if err := db.Where("uuid = ?", uuid).First(result).Error; err != nil {
		return nil
	}

	return result
}

func TransformedApartmentsEmpty() bool {
	var count int64
	db.Model(&TransformedApartment{}).Count(&count)
	return count == 0
}
