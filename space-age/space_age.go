// Package space contains functions for computing space ages
package space

// Planet holds the planet name as a string
type Planet string

// Age computes the age in years on various planets given
// arguments s for age in seconds and p, planet name as string
func Age(s float64, p Planet) float64 {
	factors := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return s / (31557600 * factors[p])
}
