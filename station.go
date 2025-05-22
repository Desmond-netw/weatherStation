package main

import (
	"fmt"
)

type WeatherStation struct {
	data map[int]*string
}

var idToKey = map[int]string{
	1:  "airTemp",
	2:  "airPressure",
	7:  "precipitation",
	11: "windSpeed",
	12: "windDirection",
	13: "humidity",
	14: "dewPoint",
	15: "soilMoisture",
	22: "cloudCover",
}

var validIDs = []int{1, 2, 7, 11, 12, 13, 14, 15, 22}

func NewWeatherStation() *WeatherStation {
	ws := &WeatherStation{data: make(map[int]*string)}
	ws.Clear()
	return ws
}

func (ws *WeatherStation) Update(id int, value string) {
	if _, ok := idToKey[id]; !ok {
		return
	}
	if value == "NULL" {
		ws.data[id] = nil
	} else {
		val := value
		ws.data[id] = &val
	}
}

func (ws *WeatherStation) Clear() {
	for _, id := range validIDs {
		ws.data[id] = nil
	}
}

func (ws *WeatherStation) Print() {
	for _, id := range validIDs {
		val := "NULL"
		if ws.data[id] != nil {
			val = *ws.data[id]
		}
		fmt.Printf("%s:%s\n", idToKey[id], val)
	}
}
