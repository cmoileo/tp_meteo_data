package weather

func FilterByCountry(stations []Station, iso string) []Station {
	var result []Station
	for _, s := range stations {
		if s.Country == iso {
			result = append(result, s)
		}
	}
	return result
}

func AvgTemperature(s Station) float64 {
	if len(s.Observations) == 0 {
		return 0
	}
	var sum float64
	for _, o := range s.Observations {
		sum += o.Temperature
	}
	return sum / float64(len(s.Observations))
}

func MaxWindGust(stations []Station) (Station, float64) {
	var maxStation Station
	var maxSpeed float64
	for _, s := range stations {
		for _, o := range s.Observations {
			if o.WindSpeed > maxSpeed {
				maxSpeed = o.WindSpeed
				maxStation = s
			}
		}
	}
	return maxStation, maxSpeed
}

func CountByCountry(stations []Station) map[string]int {
	counts := make(map[string]int)
	for _, s := range stations {
		counts[s.Country]++
	}
	return counts
}
