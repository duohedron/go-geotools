package geotools

import (
	"math"
)

// MaxLat is the  maximum latitude
const MaxLat float64 = 90

// MinLat is the  minimum latitude
const MinLat float64 = -90

// MaxLon is the  maximum longitude
const MaxLon float64 = 180

// MinLon is the  minimum longitude
const MinLon float64 = -180

const earthRadius float64 = 6378.1 // kilometers

func radToDeg(rad float64) float64 {
	return rad * (180 / math.Pi)
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

// Distance between two coordinates
func Distance(latitude1, longitude1, latitude2, longitude2 float64) float64 {
	// Convert degrees to radian
	radLat1 := latitude1 * (math.Pi / 180)
	radLon1 := longitude1 * (math.Pi / 180)
	radLat2 := latitude2 * (math.Pi / 180)
	radLon2 := longitude2 * (math.Pi / 180)
	return earthRadius * math.Acos(math.Sin(radLat1)*math.Sin(radLat2)+math.Cos(radLat1)*math.Cos(radLat2)*math.Cos(radLon1-radLon2))
}

// BoundingLocations of a circle represented py latitude longitude and distance
func BoundingLocations(latitude, longitude, distance float64) (float64, float64, float64, float64) {
	var maxLat, minLat, maxLon, minLon float64
	radLat := latitude * (math.Pi / 180)
	radLon := longitude * (math.Pi / 180)
	radDist := distance / earthRadius

	maxLat = radLat + radDist
	minLat = radLat - radDist

	if minLat > MinLat && maxLat < MaxLat {
		deltaLon := math.Asin(math.Sin(radDist) / math.Cos(radLat))
		minLon = radLon - deltaLon
		if minLon < MinLon {
			minLon += 2 * math.Pi
		}
		maxLon = radLon + deltaLon
		if maxLon > MaxLon {
			maxLon -= 2 * math.Pi
		}
	} else {
		maxLat = math.Min(maxLat, MaxLat)
		minLat = math.Max(minLat, MinLat)
		maxLon = MaxLon
		minLon = MinLon
	}

	return radToDeg(minLat), radToDeg(minLon), radToDeg(maxLat), radToDeg(maxLon)
}
