package producer

type Apartment struct {
	UUID                 string  `csv:"id" json:"UUID"`
	City                 string  `csv:"city" json:"City"`
	Type                 string  `csv:"type" json:"Type"`
	SquareMeters         float64 `csv:"squareMeters" json:"SquareMeters"`
	Rooms                float64 `csv:"rooms" json:"Rooms"`
	Floor                float64 `csv:"floor" json:"Floor"`
	FloorCount           float64 `csv:"floorCount" json:"FloorCount"`
	BuildYear            float64 `csv:"buildYear" json:"BuildYear"`
	Latitude             float64 `csv:"latitude" json:"Latitude"`
	Longitude            float64 `csv:"longitude" json:"Longitude"`
	CentreDistance       float64 `csv:"centreDistance" json:"CentreDistance"`
	PoiCount             float64 `csv:"poiCount" json:"PoiCount"`
	SchoolDistance       float64 `csv:"schoolDistance" json:"SchoolDistance"`
	ClinicDistance       float64 `csv:"clinicDistance" json:"ClinicDistance"`
	PostOfficeDistance   float64 `csv:"postOfficeDistance" json:"PostOfficeDistance"`
	KindergartenDistance float64 `csv:"kindergartenDistance" json:"KindergartenDistance"`
	RestaurantDistance   float64 `csv:"restaurantDistance" json:"RestaurantDistance"`
	CollegeDistance      float64 `csv:"collegeDistance" json:"CollegeDistance"`
	PharmacyDistance     float64 `csv:"pharmacyDistance" json:"PharmacyDistance"`
	Ownership            string  `csv:"ownership" json:"Ownership"`
	BuildingMaterial     string  `csv:"buildingMaterial" json:"BuildingMaterial"`
	Condition            string  `csv:"condition" json:"Condition"`
	HasParkingSpace      string  `csv:"hasParkingSpace" json:"HasParkingSpace"`
	HasBalcony           string  `csv:"hasBalcony" json:"HasBalcony"`
	HasElevator          string  `csv:"hasElevator" json:"HasElevator"`
	HasSecurity          string  `csv:"hasSecurity" json:"HasSecurity"`
	HasStorageRoom       string  `csv:"hasStorageRoom" json:"HasStorageRoom"`
	Price                int64   `csv:"price" json:"Price"`
}
