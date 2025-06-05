package entity

var Entities = map[string]any{
	(&RawApartment{}).TableName():         &RawApartment{},
	(&TransformedApartment{}).TableName(): &TransformedApartment{},
}
