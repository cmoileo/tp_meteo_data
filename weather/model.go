package weather

import "time"

type Station struct {
	ID           string
	Name         string
	Country      string
	Coordinates  Coordinates
	Altitude     int16
	DeviceType   string
	Manufacturer string
	InstalledOn  time.Time
	Observations []Observation
}

type Observation struct {
	Timestamp     time.Time
	Temperature   float64
	Humidity      int8
	Pressure      float64
	WindSpeed     float64
	WindDirection uint16
	Precipitation float64
	AirQuality    AirQuality
	Conditions    string
	Notes         *string
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type AirQuality struct {
	PM25 float64
	PM10 float64
	NO2  float64
}
