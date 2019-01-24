package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	const SecondsInEarthYear = 31557600
	switch planet {
	case "Earth":
		return seconds / SecondsInEarthYear
	case "Mercury":
		return seconds / (SecondsInEarthYear * 0.2408467)
	case "Venus":
		return seconds / (SecondsInEarthYear * 0.61519726)
	case "Mars":
		return seconds / (SecondsInEarthYear * 1.8808158)
	case "Jupiter":
		return seconds / (SecondsInEarthYear * 11.862615)
	case "Saturn":
		return seconds / (SecondsInEarthYear * 29.447498)
	case "Uranus":
		return seconds / (SecondsInEarthYear * 84.016846)
	case "Neptune":
		return seconds / (SecondsInEarthYear * 164.79132)
	default:
		return 0.0
	}
}