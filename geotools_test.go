package geotools

import (
	"math"
	"testing"
)

func TestRadToDeg(t *testing.T) {
	var testData = []struct {
		deg, rad float64
	}{
		{0, 0},
		{10, 10 * (math.Pi / 180)},
		{45, math.Pi / 4},
		{90, math.Pi / 2},
		{180, math.Pi},
		{270, math.Pi * 1.5},
		{360, math.Pi * 2},
	}
	for _, want := range testData {
		if deg := radToDeg(want.rad); deg != want.deg {
			t.Errorf("%v => got: %v, want: %v", want.rad, deg, want.deg)
		}
	}
}

func TestDegToRad(t *testing.T) {
	var testData = []struct {
		deg, rad float64
	}{
		{0, 0},
		{10, 10 * (math.Pi / 180)},
		{45, math.Pi / 4},
		{90, math.Pi / 2},
		{180, math.Pi},
		{270, math.Pi * 1.5},
		{360, math.Pi * 2},
	}
	for _, want := range testData {
		if rad := degToRad(want.deg); rad != want.rad {
			t.Errorf("%v => got: %v, want: %v", want.deg, rad, want.rad)
		}
	}
}
func TestDistance01(t *testing.T) {
	var testData = []struct {
		lat1, lon1, lat2, lon2, dist float64
	}{
		{0, 0, 47, 19, 5548.817065746805},
		{52.654, 12.456, -60.6548, -32.456, 13229.661393713557},
	}
	for _, want := range testData {
		got := Distance(want.lat1, want.lon1, want.lat2, want.lon2)

		if got != want.dist {
			t.Errorf("Distance got: %v want: %v", got, want.dist)
		}
	}
}

func TestBoundingLocations01(t *testing.T) {
	var testData = []struct {
		lat, lon, dist, minLat, minLon, maxLat, maxLon float64
	}{
		{0, 0, 10, -0.08983204953368921, -0.08983204953368921, 0.08983204953368921, 0.08983204953368921},
		{0, 0, 100, -0.8983204953368921, -0.8983204953368921, 0.8983204953368921, 0.8983204953368921},
		{0, 0, 1000, -8.983204953368922, -8.98320495336892, 8.983204953368922, 8.98320495336892},
		{47, 19, 10, 46.91016795046631, 18.868281073506612, 47.08983204953369, 19.131718926493388},
		{47, 19, 100, 46.10167950466311, 17.682749283000515, 47.8983204953369, 20.317250716999485},
		{47, 19, 125.4, 45.87350609884754, 17.348123020935592, 48.126493901152465, 20.65187697906441},
		{-91, -181, 100, -91.89832049533689, -117.06067736082304, -90.10167950466311, -244.939322639177},
	}

	for _, want := range testData {
		gotMinLat, gotMinLon, gotMaxLat, gotMaxLon := BoundingLocations(want.lat, want.lon, want.dist)

		if gotMinLat != want.minLat {
			t.Errorf("Lat: %v, Lon: %v, Dist: %v => MinLat got: %v want: %v", want.lat, want.lon, want.dist, gotMinLat, want.minLat)
		}
		if gotMinLon != want.minLon {
			t.Errorf("Lat: %v, Lon: %v, Dist: %v => MinLon got: %v want: %v", want.lat, want.lon, want.dist, gotMinLon, want.minLon)
		}
		if gotMaxLat != want.maxLat {
			t.Errorf("Lat: %v, Lon: %v, Dist: %v => MaxLat got: %v want: %v", want.lat, want.lon, want.dist, gotMaxLat, want.maxLat)
		}
		if gotMaxLon != want.maxLon {
			t.Errorf("Lat: %v, Lon: %v, Dist: %v => MaxLon got: %v want: %v", want.lat, want.lon, want.dist, gotMaxLon, want.maxLon)
		}

	}
}
