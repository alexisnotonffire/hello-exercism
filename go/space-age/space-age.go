// Package space provides a function for calculating your age on
// different planets (sorry Pluto!) in the solar system.
package space

// Declaration of available planets for conversion. The variable year
// is the number of seconds in an Earth year and the factor for each
// planet is from Planet years   to Earth years.
var (
	year          float64 = 31557600
	planetDetails         = map[string]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
)

// Age converts earth seconds to the equivalent number of years
// on the input planet.
func Age(earthAge float64, planet string) float64 {
	factor := planetDetails[planet]
	age := earthAge / year / factor
	return age
}
