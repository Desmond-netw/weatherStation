package main

import (
	"strconv"
	"strings"
)

// Handles parsing a data update line like "11,15.5"
func ProcessUpdateLine(ws *WeatherStation, line string) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return
	}
	value := strings.TrimSpace(parts[1])
	ws.Update(id, value)
}
