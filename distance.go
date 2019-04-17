package nhsfinder

// Lovingly copied and pasted and edited from
// https://gist.github.com/cdipaolo/d3f8db3848278b49db68

import "math"

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

// Distance returns distance between 2 postcodes
func Distance(p1, p2 LatLng) float64 {
	return distance(p1.Lat, p1.Lng, p2.Lat, p2.Lng)
}
