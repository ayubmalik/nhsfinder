package nhsfinder

// LatLng is point with latitude and longitude
type LatLng struct {
	Lat float64
	Lng float64
}

// Postcode is a single UK postcode with latlng
type Postcode struct {
	Value  string
	LatLng LatLng
}

// Address is a UK address
type Address struct {
	Line1    string
	Line2    string
	Line3    string
	Line4    string
	Line5    string
	Postcode Postcode
}

// UpdatePostcode updates postcode
func (a *Address) UpdatePostcode(postcode Postcode) {
	a.Postcode = postcode
}

// Pharmacy in the UK
type Pharmacy struct {
	ID      string
	Name    string
	Address *Address
	Phone   string
}
