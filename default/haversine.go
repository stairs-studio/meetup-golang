package main

import (
	"math"
)

const (
	earthRaidusKm = 6378.16 // radius of the earth in kilometers.
)

// Coord represents a geographic coordinate.
type Coord struct {
	lat float64
	lng float64
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func radiansToDegrees(d float64) float64 {
	return d * 180 / math.Pi
}

//Coordinate - new Object Coord
func Coordinate(lat, lng float64) *Coord {
	return &Coord{
		lat: lat,
		lng: lng,
	}
}

//Distance - distance between two points
func Distance(p, q *Coord) (km float64) {
	lat1 := degreesToRadians(p.lat)
	lon1 := degreesToRadians(p.lng)
	lat2 := degreesToRadians(q.lat)
	lon2 := degreesToRadians(q.lng)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	km = c * earthRaidusKm

	return km
}
