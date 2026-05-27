package weather

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
