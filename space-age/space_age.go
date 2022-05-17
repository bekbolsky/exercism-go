package space

type Planet string

const (
	earthSeconds   float64 = 31557600
	mercurySeconds float64 = earthSeconds * 0.2408467
	venusSeconds   float64 = earthSeconds * 0.61519726
	marsSeconds    float64 = earthSeconds * 1.8808158
	jupiterSeconds float64 = earthSeconds * 11.862615
	saturnSeconds  float64 = earthSeconds * 29.447498
	uranusSeconds  float64 = earthSeconds * 84.016846
	neptuneSeconds float64 = earthSeconds * 164.79132
)

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return seconds / earthSeconds
	case "Mercury":
		return seconds / mercurySeconds
	case "Venus":
		return seconds / venusSeconds
	case "Mars":
		return seconds / marsSeconds
	case "Jupiter":
		return seconds / jupiterSeconds
	case "Saturn":
		return seconds / saturnSeconds
	case "Uranus":
		return seconds / uranusSeconds
	case "Neptune":
		return seconds / neptuneSeconds
	}
	return 0
}
