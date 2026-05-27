package main

import (
	"fmt"
	"tp_meteo_data/weather"
)

func main() {
	data, err := weather.LoadFromJSON("datas/meteo_datas.json")
	if err != nil {
		return
	}

	fmt.Println(data)
}
