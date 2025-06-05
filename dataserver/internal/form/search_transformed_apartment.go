package form

type SearchTransformedApartments struct {
	Split string `form:"split"`
	Count int    `form:"count"`
	Order string `form:"order"`

	MinSquareMeters   float64 `form:"min_squareMeters"`
	MaxSquareMeters   float64 `form:"max_squareMeters"`
	MinRooms          float64 `form:"min_rooms"`
	MaxRooms          float64 `form:"max_rooms"`
	MinFloor          float64 `form:"min_floor"`
	MaxFloor          float64 `form:"max_floor"`
	MinFloorCount     float64 `form:"min_floorCount"`
	MaxFloorCount     float64 `form:"max_floorCount"`
	MinCentreDistance float64 `form:"min_centreDistance"`
	MaxCentreDistance float64 `form:"max_centreDistance"`
	MinPoiCount       float64 `form:"min_poiCount"`
	MaxPoiCount       float64 `form:"max_poiCount"`
	MinAge            float64 `form:"min_age"`
	MaxAge            float64 `form:"max_age"`

	HasParkingSpace *bool `form:"has_ParkingSpace"`
	HasBalcony      *bool `form:"has_Balcony"`
	HasElevator     *bool `form:"has_Elevator"`
	HasSecurity     *bool `form:"has_Security"`
	HasStorageRoom  *bool `form:"has_StorageRoom"`
	HasSchool       *bool `form:"has_School"`
	HasClinic       *bool `form:"has_Clinic"`
	HasPostoffice   *bool `form:"has_Postoffice"`
	HasKindergarten *bool `form:"has_Kindergarten"`
	HasRestaurant   *bool `form:"has_Restaurant"`
	HasCollege      *bool `form:"has_College"`
	HasPharmacy     *bool `form:"has_Pharmacy"`

	CityBialystok   *bool `form:"city_bialystok"`
	CityBydgoszcz   *bool `form:"city_bydgoszcz"`
	CityCzestochowa *bool `form:"city_czestochowa"`
	CityGdansk      *bool `form:"city_gdansk"`
	CityGdynia      *bool `form:"city_gdynia"`
	CityKatowice    *bool `form:"city_katowice"`
	CityKrakow      *bool `form:"city_krakow"`
	CityLodz        *bool `form:"city_lodz"`
	CityLublin      *bool `form:"city_lublin"`
	CityPoznan      *bool `form:"city_poznan"`
	CityRadom       *bool `form:"city_radom"`
	CityRzeszow     *bool `form:"city_rzeszow"`
	CitySzczecin    *bool `form:"city_szczecin"`
	CityWarszawa    *bool `form:"city_warszawa"`
	CityWroclaw     *bool `form:"city_wroclaw"`

	TypeApartmentBuilding *bool `form:"type_apartmentBuilding"`
	TypeBlockOfFlats      *bool `form:"type_blockOfFlats"`
	TypeTenement          *bool `form:"type_tenement"`

	MinPrice int64 `form:"min_price"`
	MaxPrice int64 `form:"max_price"`
}
