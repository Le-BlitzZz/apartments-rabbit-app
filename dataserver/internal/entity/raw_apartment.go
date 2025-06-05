package entity

import (
	"gorm.io/gorm"
)

type RawApartments []RawApartment

func (a RawApartments) Create() error {
	return db.CreateInBatches(a, 1000).Error
}

type RawApartment struct {
	gorm.Model

	UUID                 string  `json:"UUID"`
	City                 string  `json:"City"`
	Type                 string  `json:"Type"`
	SquareMeters         float64 `json:"SquareMeters"`
	Rooms                float64 `json:"Rooms"`
	Floor                float64 `json:"Floor"`
	FloorCount           float64 `json:"FloorCount"`
	BuildYear            float64 `json:"BuildYear"`
	Latitude             float64 `json:"Latitude"`
	Longitude            float64 `json:"Longitude"`
	CentreDistance       float64 `json:"CentreDistance"`
	PoiCount             float64 `json:"PoiCount"`
	SchoolDistance       float64 `json:"SchoolDistance"`
	ClinicDistance       float64 `json:"ClinicDistance"`
	PostOfficeDistance   float64 `json:"PostOfficeDistance"`
	KindergartenDistance float64 `json:"KindergartenDistance"`
	RestaurantDistance   float64 `json:"RestaurantDistance"`
	CollegeDistance      float64 `json:"CollegeDistance"`
	PharmacyDistance     float64 `json:"PharmacyDistance"`
	Ownership            string  `json:"Ownership"`
	BuildingMaterial     string  `json:"BuildingMaterial"`
	Condition            string  `json:"Condition"`
	HasParkingSpace      string  `json:"HasParkingSpace"`
	HasBalcony           string  `json:"HasBalcony"`
	HasElevator          string  `json:"HasElevator"`
	HasSecurity          string  `json:"HasSecurity"`
	HasStorageRoom       string  `json:"HasStorageRoom"`
	Price                int64   `json:"Price"`
}

func (*RawApartment) TableName() string {
	return "raw_apartments"
}

func FindRawApartment(uuid string) *RawApartment {
	if uuid == "" {
		return nil
	}

	result := &RawApartment{}

	if err := db.Where("uuid = ?", uuid).First(result).Error; err != nil {
		return nil
	}

	return result
}

func RawApartmentsEmpty() bool {
	var count int64
	db.Model(&RawApartment{}).Count(&count)
	return count == 0
}
