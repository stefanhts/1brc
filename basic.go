package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const file = "weather_stations.csv"

type WeatherStation struct {
	Name      string
	MinTemp   float64
	MaxTemp   float64
	SumTemp   float64
	CountTemp int
}

var weather_stations map[string]WeatherStation

func main() {
	start := time.Now()
	weather_stations = make(map[string]WeatherStation)
	file, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	/* 	sc / anner.Scan() // skip the first line
	/* 	scanner.Scan() // skip the second line */
	for scanner.Scan() {
		line := scanner.Text()
		name, temp := processLine(line)
		if station, ok := weather_stations[name]; !ok {
			station := WeatherStation{
				Name:      name,
				MinTemp:   temp,
				MaxTemp:   temp,
				SumTemp:   temp,
				CountTemp: 1,
			}
			weather_stations[name] = station
		} else {
			station.MinTemp = min(temp, station.MinTemp)
			station.MaxTemp = max(temp, station.MaxTemp)
			station.SumTemp += temp
			station.CountTemp++

			weather_stations[name] = station
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
	printResults()
}

func processLine(line string) (string, float64) {
	splits := strings.Split(line, ";")
	name := splits[0]
	temp, _ := strconv.ParseFloat(splits[1], 64)
	temp = round(temp)

	return name, temp
}

func printResults() {
	keys := make([]string, 0, len(weather_stations))

	// Iterate over the map and append each key to the slice
	for k := range weather_stations {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	fmt.Printf("{")
	for i, k := range keys {
		v := weather_stations[k]
		fmt.Printf("%s=%.1f/%.1f/%.1f", k, v.MinTemp, round((v.SumTemp / float64(v.CountTemp))), v.MaxTemp)
		if i < len(keys)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}")
}

func round(f float64) float64 {
	return math.Round(f*10) / 10
}
