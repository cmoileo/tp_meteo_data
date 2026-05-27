package main

import (
	"fmt"
	"log"
	"tp_meteo_data/weather"
)

func main() {
	jsonStations, err := weather.LoadFromJSON("datas/meteo_datas.json")
	if err != nil {
		log.Fatalf("JSON: %v", err)
	}
	xmlStations, err := weather.LoadFromXML("datas/meteo_datas.xml")
	if err != nil {
		log.Fatalf("XML: %v", err)
	}

	jsonObs := 0
	for _, s := range jsonStations {
		jsonObs += len(s.Observations)
	}
	xmlObs := 0
	for _, s := range xmlStations {
		xmlObs += len(s.Observations)
	}

	fmt.Printf("JSON : %d stations, %d observations\n", len(jsonStations), jsonObs)
	fmt.Printf("XML  : %d stations, %d observations\n", len(xmlStations), xmlObs)

	if len(jsonStations) == len(xmlStations) && jsonObs == xmlObs {
		fmt.Println("Cohérence : OK")
	} else {
		fmt.Println("Cohérence : KO")
	}

	station, gust := weather.MaxWindGust(jsonStations)
	fmt.Printf("Station la plus ventée : %s (%.1f km/h)\n", station.ID, gust)

	bordeaux := weather.FilterByCountry(jsonStations, "FR")
	var bordeauxStation weather.Station
	for _, s := range bordeaux {
		if s.ID == "FR-BOR-001" {
			bordeauxStation = s
			break
		}
	}
	avg := weather.AvgTemperature(bordeauxStation)
	fmt.Printf("Temp. moyenne Bordeaux Mérignac : %.1f °C\n", avg)

	byCountry := weather.CountByCountry(jsonStations)
	fmt.Printf("Stations par pays : %v\n", byCountry)
}
