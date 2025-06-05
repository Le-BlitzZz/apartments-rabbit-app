package form

type SearchRawApartments struct {
	City            string `form:"city"`
	Type            string `form:"type"`
	Ownership       string `form:"ownership"`
	HasParkingSpace string `form:"has_parking_space"`
	HasBalcony      string `form:"has_balcony"`
	HasElevator     string `form:"has_elevator"`
	HasSecurity     string `form:"has_security"`
	HasStorageRoom  string `form:"has_storage_room"`
	Split           string `form:"split"`

	MinSquareMeters float64 `form:"min_square_meters"`
	MaxSquareMeters float64 `form:"max_square_meters"`
	MinRooms        float64 `form:"min_rooms"`
	MaxRooms        float64 `form:"max_rooms"`
	MinFloor        float64 `form:"min_floor"`
	MaxFloor        float64 `form:"max_floor"`
	MinFloorCount   float64 `form:"min_floor_count"`
	MaxFloorCount   float64 `form:"max_floor_count"`
	MinBuildYear    int     `form:"min_build_year"`
	MaxBuildYear    int     `form:"max_build_year"`
	// MinLatitude             float64 `form:"min_latitude"`
	// MaxLatitude             float64 `form:"max_latitude"`
	// MinLongitude            float64 `form:"min_longitude"`
	// MaxLongitude            float64 `form:"max_longitude"`
	// MinCentreDistance       float64 `form:"min_centre_distance"`
	// MaxCentreDistance       float64 `form:"max_centre_distance"`
	// MinPoiCount             float64 `form:"min_poi_count"`
	// MaxPoiCount             float64 `form:"max_poi_count"`
	// MinSchoolDistance       float64 `form:"min_school_distance"`
	// MaxSchoolDistance       float64 `form:"max_school_distance"`
	// MinClinicDistance       float64 `form:"min_clinic_distance"`
	// MaxClinicDistance       float64 `form:"max_clinic_distance"`
	// MinPostOfficeDistance   float64 `form:"min_post_office_distance"`
	// MaxPostOfficeDistance   float64 `form:"max_post_office_distance"`
	// MinKindergartenDistance float64 `form:"min_kindergarten_distance"`
	// MaxKindergartenDistance float64 `form:"max_kindergarten_distance"`
	// MinRestaurantDistance   float64 `form:"min_restaurant_distance"`
	// MaxRestaurantDistance   float64 `form:"max_restaurant_distance"`
	// MinCollegeDistance      float64 `form:"min_college_distance"`
	// MaxCollegeDistance      float64 `form:"max_college_distance"`
	// MinPharmacyDistance     float64 `form:"min_pharmacy_distance"`
	// MaxPharmacyDistance     float64 `form:"max_pharmacy_distance"`
	MinPrice int64 `form:"min_price"`
	MaxPrice int64 `form:"max_price"`

	Count int    `form:"count"`
	Order string `form:"order"`
}
