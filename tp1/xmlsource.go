package tp1

import (
	"encoding/xml"
	"os"
	"time"
)

type xmlDataset struct {
	XMLName  struct{}     `xml:"weather_dataset"`
	Version  string       `xml:"version,attr"`
	Source   string       `xml:"source,attr"`
	Stations []xmlStation `xml:"station"`
}

type xmlStation struct {
	ID      string    `xml:"id,attr"`
	Country string    `xml:"country,attr"`
	Name    string    `xml:"name"`
	Coords  xmlCoords `xml:"coordinates"`
	HW      xmlHardware `xml:"hardware"`
	ObsList xmlObsList `xml:"observations"`
}

type xmlCoords struct {
	Lat      float64 `xml:"lat,attr"`
	Lon      float64 `xml:"lon,attr"`
	Altitude int16   `xml:"altitude,attr"`
}

type xmlHardware struct {
	Vendor string `xml:"vendor,attr"`
	Model  string `xml:"model,attr"`
	Since  string `xml:"since,attr"`
}

type xmlObsList struct {
	Count        int               `xml:"count,attr"`
	Observations []xmlObservation  `xml:"observation"`
}

type xmlObservation struct {
	At         string         `xml:"at,attr"`
	Sky        string         `xml:"sky,attr"`
	Measures   []xmlMeasure   `xml:"measure"`
	Wind       xmlWind        `xml:"wind"`
	AirQuality xmlAirQuality  `xml:"air_quality"`
	Note       string         `xml:"note"`
}

type xmlMeasure struct {
	Type  string  `xml:"type,attr"`
	Unit  string  `xml:"unit,attr"`
	Value float64 `xml:",chardata"`
}

type xmlWind struct {
	Speed     float64 `xml:"speed,attr"`
	Direction uint16  `xml:"direction,attr"`
}

type xmlAirQuality struct {
	Pollutants []xmlPollutant `xml:"pollutant"`
}

type xmlPollutant struct {
	Name  string  `xml:"name,attr"`
	Value float64 `xml:",chardata"`
}

func LoadFromXML(path string) ([]Station, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var ds xmlDataset
	if err := xml.Unmarshal(data, &ds); err != nil {
		return nil, err
	}

	stations := make([]Station, 0, len(ds.Stations))
	for _, xs := range ds.Stations {
		installedOn, _ := time.Parse("2006-01-02", xs.HW.Since)

		obs := make([]Observation, 0, len(xs.ObsList.Observations))
		for _, xo := range xs.ObsList.Observations {
			ts, err := time.Parse(time.RFC3339, xo.At)
			if err != nil {
				continue
			}

			var temperature, humidity, pressure, precipitation float64
			for _, m := range xo.Measures {
				switch m.Type {
				case "temperature":
					temperature = m.Value
				case "humidity":
					humidity = m.Value
				case "pressure":
					pressure = m.Value
				case "precipitation":
					precipitation = m.Value
				}
			}

			var pm25, pm10, no2 float64
			for _, p := range xo.AirQuality.Pollutants {
				switch p.Name {
				case "PM2.5":
					pm25 = p.Value
				case "PM10":
					pm10 = p.Value
				case "NO2":
					no2 = p.Value
				}
			}

			var notes *string
			if xo.Note != "" {
				notes = &xo.Note
			}

			obs = append(obs, Observation{
				Timestamp:     ts,
				Temperature:   temperature,
				Humidity:      int8(humidity),
				Pressure:      pressure,
				WindSpeed:     xo.Wind.Speed,
				WindDirection: xo.Wind.Direction,
				Precipitation: precipitation,
				AirQuality: AirQuality{
					PM25: pm25,
					PM10: pm10,
					NO2:  no2,
				},
				Conditions: xo.Sky,
				Notes:      notes,
			})
		}

		stations = append(stations, Station{
			ID:           xs.ID,
			Name:         xs.Name,
			Country:      xs.Country,
			Coordinates:  Coordinates{Latitude: xs.Coords.Lat, Longitude: xs.Coords.Lon},
			Altitude:     xs.Coords.Altitude,
			DeviceType:   xs.HW.Model,
			Manufacturer: xs.HW.Vendor,
			InstalledOn:  installedOn,
			Observations: obs,
		})
	}

	return stations, nil
}
