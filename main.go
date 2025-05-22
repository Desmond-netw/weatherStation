package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("--- Weather Station ---")

	station := NewWeatherStation()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		switch line {
		case "exit":
			fmt.Println("Exiting...")
			return
		case "get":
			station.Print()
		case "clear":
			station.Clear()
		default:
			ProcessUpdateLine(station, line)
		}
	}
}
