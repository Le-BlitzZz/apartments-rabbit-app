package search

import (
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/entity"
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/form"
)

func TransformedApartments(frm form.SearchTransformedApartments) (result entity.TransformedApartments, err error) {
	s := entity.Db().Where(&entity.TransformedApartment{})

	if frm.Split != "" {
		s = s.Where("split = ?", frm.Split)
	}

	if frm.Count > 0 {
		s = s.Limit(frm.Count)
	}

	if frm.MinSquareMeters > 0 {
		s = s.Where("square_meters >= ?", frm.MinSquareMeters)
	}

	if frm.MaxSquareMeters > 0 {
		s = s.Where("square_meters <= ?", frm.MaxSquareMeters)
	}

	if frm.MinRooms > 0 {
		s = s.Where("rooms >= ?", frm.MinRooms)
	}

	if frm.MaxRooms > 0 {
		s = s.Where("rooms <= ?", frm.MaxRooms)
	}

	if frm.MinFloor > 0 {
		s = s.Where("floor >= ?", frm.MinFloor)
	}

	if frm.MaxFloor > 0 {
		s = s.Where("floor <= ?", frm.MaxFloor)
	}

	if frm.MinFloorCount > 0 {
		s = s.Where("floor_count >= ?", frm.MinFloorCount)
	}

	if frm.MaxFloorCount > 0 {
		s = s.Where("floor_count <= ?", frm.MaxFloorCount)
	}

	if frm.MinCentreDistance > 0 {
		s = s.Where("centre_distance >= ?", frm.MinCentreDistance)
	}

	if frm.MaxCentreDistance > 0 {
		s = s.Where("centre_distance <= ?", frm.MaxCentreDistance)
	}

	if frm.MinPoiCount > 0 {
		s = s.Where("poi_count >= ?", frm.MinPoiCount)
	}

	if frm.MaxPoiCount > 0 {
		s = s.Where("poi_count <= ?", frm.MaxPoiCount)
	}

	if frm.MinAge > 0 {
		s = s.Where("age >= ?", frm.MinAge)
	}

	if frm.MaxAge > 0 {
		s = s.Where("age <= ?", frm.MaxAge)
	}

	if frm.HasParkingSpace != nil {
		s = s.Where("has_parking_space = ?", *frm.HasParkingSpace)
	}

	if frm.HasBalcony != nil {
		s = s.Where("has_balcony = ?", *frm.HasBalcony)
	}

	if frm.HasElevator != nil {
		s = s.Where("has_elevator = ?", *frm.HasElevator)
	}

	if frm.HasSecurity != nil {
		s = s.Where("has_security = ?", *frm.HasSecurity)
	}

	if frm.HasStorageRoom != nil {
		s = s.Where("has_storage_room = ?", *frm.HasStorageRoom)
	}

	if frm.HasSchool != nil {
		s = s.Where("has_school = ?", *frm.HasSchool)
	}

	if frm.HasClinic != nil {
		s = s.Where("has_clinic = ?", *frm.HasClinic)
	}

	if frm.HasPostoffice != nil {
		s = s.Where("has_postoffice = ?", *frm.HasPostoffice)
	}

	if frm.HasKindergarten != nil {
		s = s.Where("has_kindergarten = ?", *frm.HasKindergarten)
	}

	if frm.HasRestaurant != nil {
		s = s.Where("has_restaurant = ?", *frm.HasRestaurant)
	}

	if frm.HasCollege != nil {
		s = s.Where("has_college = ?", *frm.HasCollege)
	}

	if frm.HasPharmacy != nil {
		s = s.Where("has_pharmacy = ?", *frm.HasPharmacy)
	}

	if frm.TypeApartmentBuilding != nil {
		s = s.Where("type_apartment_building = ?", *frm.TypeApartmentBuilding)
	}

	if frm.TypeBlockOfFlats != nil {
		s = s.Where("type_block_of_flats = ?", *frm.TypeBlockOfFlats)
	}

	if frm.TypeTenement != nil {
		s = s.Where("type_tenement = ?", *frm.TypeTenement)
	}

	if frm.CityBialystok != nil {
		s = s.Where("city_bialystok = ?", *frm.CityBialystok)
	}

	if frm.CityBydgoszcz != nil {
		s = s.Where("city_bydgoszcz = ?", *frm.CityBydgoszcz)
	}

	if frm.CityCzestochowa != nil {
		s = s.Where("city_czestochowa = ?", *frm.CityCzestochowa)
	}

	if frm.CityGdansk != nil {
		s = s.Where("city_gdansk = ?", *frm.CityGdansk)
	}

	if frm.CityGdynia != nil {
		s = s.Where("city_gdynia = ?", *frm.CityGdynia)
	}

	if frm.CityKatowice != nil {
		s = s.Where("city_katowice = ?", *frm.CityKatowice)
	}

	if frm.CityKrakow != nil {
		s = s.Where("city_krakow = ?", *frm.CityKrakow)
	}

	if frm.CityLodz != nil {
		s = s.Where("city_lodz = ?", *frm.CityLodz)
	}

	if frm.CityLublin != nil {
		s = s.Where("city_lublin = ?", *frm.CityLublin)
	}

	if frm.CityPoznan != nil {
		s = s.Where("city_poznan = ?", *frm.CityPoznan)
	}

	if frm.CityRadom != nil {
		s = s.Where("city_radom = ?", *frm.CityRadom)
	}

	if frm.CityRzeszow != nil {
		s = s.Where("city_rzeszow = ?", *frm.CityRzeszow)
	}

	if frm.CitySzczecin != nil {
		s = s.Where("city_szczecin = ?", *frm.CitySzczecin)
	}

	if frm.CityWarszawa != nil {
		s = s.Where("city_warszawa = ?", *frm.CityWarszawa)
	}

	if frm.CityWroclaw != nil {
		s = s.Where("city_wroclaw = ?", *frm.CityWroclaw)
	}

	if frm.MinPrice > 0 {
		s = s.Where("price >= ?", frm.MinPrice)
	}

	if frm.MaxPrice > 0 {
		s = s.Where("price <= ?", frm.MaxPrice)
	}

	if err = s.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
