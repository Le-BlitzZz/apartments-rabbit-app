package entity

type Apartments interface {
	Create() error
}

func TypeName(a Apartments) string {
	switch a.(type) {
	case RawApartments:
		return "Raw Apartments"
	case TransformedApartments:
		return "Transformed Apartments"
	default:
		return "Unknown Apartments"
	}
}
