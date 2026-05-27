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

	fmt.Printf("JSON: %d stations, %d observations\n", len(jsonStations), jsonObs)
	fmt.Printf("XML:  %d stations, %d observations\n", len(xmlStations), xmlObs)

	if len(jsonStations) == len(xmlStations) && jsonObs == xmlObs {
		fmt.Println("Coherence: OK")
	} else {
		fmt.Println("Coherence: KO")
	}
}
