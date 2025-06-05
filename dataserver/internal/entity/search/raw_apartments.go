package search

import (
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/entity"
	"github.com/Le-BlitzZz/apartments-rabbit-app/dataserver/internal/form"
)

func RawApartments(frm form.SearchRawApartments) (result entity.RawApartments, err error) {
	s := entity.Db().Where(&entity.RawApartment{})

	if frm.Count > 0 {
		s = s.Limit(frm.Count)
	}

	if frm.City != "" {
		s = s.Where("city = ?", frm.City)
	}

	if frm.Type != "" {
		s = s.Where("type = ?", frm.Type)
	}

	if frm.Ownership != "" {
		s = s.Where("ownership = ?", frm.Ownership)
	}

	if frm.HasParkingSpace != "" {
		s = s.Where("has_parking_space = ?", frm.HasParkingSpace)
	}

	if frm.HasBalcony != "" {
		s = s.Where("has_balcony = ?", frm.HasBalcony)
	}

	if frm.HasElevator != "" {
		s = s.Where("has_elevator = ?", frm.HasElevator)
	}

	if frm.HasSecurity != "" {
		s = s.Where("has_security = ?", frm.HasSecurity)
	}

	if frm.HasStorageRoom != "" {
		s = s.Where("has_storage_room = ?", frm.HasStorageRoom)
	}

	if frm.Split != "" {
		s = s.Where("split = ?", frm.Split)
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

	if frm.MinBuildYear > 0 {
		s = s.Where("build_year >= ?", frm.MinBuildYear)
	}

	if frm.MaxBuildYear > 0 {
		s = s.Where("build_year <= ?", frm.MaxBuildYear)
	}

	if frm.MinPrice > 0 {
		s = s.Where("price >= ?", frm.MinPrice)
	}

	if frm.MaxPrice > 0 {
		s = s.Where("price <= ?", frm.MaxPrice)
	}

	switch frm.Order {
	case "price_asc":
		s = s.Order("price ASC")
	case "price_desc":
		s = s.Order("price DESC")
	case "city":
		s = s.Order("city ASC")
	case "type":
		s = s.Order("type ASC")
	case "ownership":
		s = s.Order("ownership ASC")
	case "build_year_asc":
		s = s.Order("build_year ASC")
	case "build_year_desc":
		s = s.Order("build_year DESC")
	}

	if err = s.Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
