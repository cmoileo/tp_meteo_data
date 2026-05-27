package weather

import (
	"encoding/json"
	"os"
	"time"
)

type jsonRoot struct {
	Metadata jsonMetadata  `json:"metadata"`
	Stations []jsonStation `json:"stations"`
}

type jsonMetadata struct {
	Version      string `json:"version"`
	Source       string `json:"source"`
	GeneratedAt  string `json:"generated_at"`
	StationCount int    `json:"station_count"`
	License      string `json:"license"`
}

type jsonStation struct {
	ID       string       `json:"id"`
	Name     string       `json:"name"`
	Country  string       `json:"country"`
	Altitude int16        `json:"altitude_m"`
	Location jsonLocation `json:"location"`
	Device   jsonDevice   `json:"device"`
	Obs      []jsonObs    `json:"observations"`
}

type jsonLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type jsonDevice struct {
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	InstalledOn  string `json:"installed_on"`
}

type jsonObs struct {
	Timestamp     string         `json:"timestamp"`
	Temp          float64        `json:"temperature_celsius"`
	Humidity      int8           `json:"humidity_percent"`
	Pressure      float64        `json:"pressure_hpa"`
	Wind          jsonWind       `json:"wind"`
	Precipitation float64        `json:"precipitation_mm"`
	AirQuality    jsonAirQuality `json:"air_quality"`
	Conditions    string         `json:"conditions"`
	Notes         *string        `json:"notes,omitempty"`
}

type jsonWind struct {
	Speed     float64 `json:"speed_kmh"`
	Direction uint16  `json:"direction_deg"`
}

type jsonAirQuality struct {
	PM25 float64 `json:"pm25"`
	PM10 float64 `json:"pm10"`
	NO2  float64 `json:"no2"`
}

var countryMapping = map[string]string{
	"France":    "FR",
	"Espagne":   "ES",
	"Allemagne": "DE",
	"Italie":    "IT",
	"Portugal":  "PT",
	"Belgique":  "BE",
	"Pays-Bas":  "NL",
	"Autriche":  "AT",
	"Suisse":    "CH",
	"Tchéquie":  "CZ",
	"Danemark":  "DK",
	"Suède":     "SE",
	"Norvège":   "NO",
	"Pologne":   "PL",
}

func LoadFromJSON(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var root jsonRoot
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	stations := make([]Station, 0, len(root.Stations))
	for _, js := range root.Stations {
		installedOn, _ := time.Parse("2006-01-02", js.Device.InstalledOn)

		country := js.Country
		if c, ok := countryMapping[country]; ok {
			country = c
		}

		obs := make([]Observation, 0, len(js.Obs))
		for _, jo := range js.Obs {
			ts, err := time.Parse(time.RFC3339, jo.Timestamp)
			if err != nil {
				continue
			}
			obs = append(obs, Observation{
				Timestamp:     ts,
				Temperature:   jo.Temp,
				Humidity:      jo.Humidity,
				Pressure:      jo.Pressure,
				WindSpeed:     jo.Wind.Speed,
				WindDirection: jo.Wind.Direction,
				Precipitation: jo.Precipitation,
				AirQuality: AirQuality{
					PM25: jo.AirQuality.PM25,
					PM10: jo.AirQuality.PM10,
					NO2:  jo.AirQuality.NO2,
				},
				Conditions: jo.Conditions,
				Notes:      jo.Notes,
			})
		}

		stations = append(stations, Station{
			ID:           js.ID,
			Name:         js.Name,
			Country:      country,
			Coordinates:  Coordinates{Latitude: js.Location.Latitude, Longitude: js.Location.Longitude},
			Altitude:     js.Altitude,
			DeviceType:   js.Device.Type,
			Manufacturer: js.Device.Manufacturer,
			InstalledOn:  installedOn,
			Observations: obs,
		})
	}

	return stations, nil
}
