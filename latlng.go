package latlng

import (
	"fmt"
	"math"
)

const (
	earthRadius = 6.3781e6
)

type Position struct {
	Lat float64
	Lng float64
}

func DecToDMS(f float64) string {
	d := math.Floor(f)
	m := math.Floor((f - d) * 60)
	s := (f - d - m/60) * 3600

	return fmt.Sprintf("%.0f° %.0f' %.0f\"", d, m, s)
}

func DecToRad(f float64) float64 {
	return f * math.Pi / 180
}

func RadToDec(f float64) float64 {
	return f * 180 / math.Pi
}

func HaversineDistance(decPos1, decPos2 Position) float64 {
	ph1 := DecToRad(decPos1.Lat)
	ph2 := DecToRad(decPos2.Lat)

	deltaPh := DecToRad(decPos2.Lat - decPos1.Lat)
	deltaL := DecToRad(decPos2.Lng - decPos1.Lng)

	a := math.Sin(deltaPh/2)*math.Sin(deltaPh/2) + math.Cos(ph1)*math.Cos(ph2)*math.Sin(deltaL/2)*math.Sin(deltaL/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	fmt.Printf("%.2f\n", c)
	return earthRadius * c
}

func SphericalLawDistance(decPos1, decPos2 Position) float64 {
	ph1 := DecToRad(decPos1.Lat)
	ph2 := DecToRad(decPos2.Lat)
	deltaL := DecToRad(decPos2.Lng - decPos1.Lng)

	return math.Acos(math.Sin(ph1)*math.Sin(ph2)+math.Cos(ph1)*math.Cos(ph2)*math.Cos(deltaL)) * earthRadius

}

func Bearing(decPos1, decPos2 Position) float64 {
	ph1 := DecToRad(decPos1.Lat)
	ph2 := DecToRad(decPos2.Lat)
	l1 := DecToRad(decPos1.Lng)
	l2 := DecToRad(decPos2.Lng)

	y := math.Sin(l2-l1) * math.Cos(ph2)
	x := math.Cos(ph1)*math.Sin(ph2) - math.Sin(ph1)*math.Cos(ph2)*math.Cos(l2-l1)

	teta := math.Atan2(y, x)
	f := RadToDec(teta) + 360
	return f - (math.Floor(f/360) * 360)

}
