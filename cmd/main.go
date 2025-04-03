package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	filePath  = "../data/measurements.txt"
	fileLines = 1000000
)

type OutputData struct {
	City  string
	Min   float64
	Mean  float64
	Max   float64
	Count int
}

func main() {
	file, err := openFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cityMap := make(map[string]OutputData, fileLines)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		tempData := strings.Split(line, ";")
		city := tempData[0]

		temperatureStr := tempData[1]
		if temperatureStr == "" {
			panic(nil)
		}

		if len(temperatureStr) < 3 {
			continue
		}

		temperature, parseErr := strconv.ParseFloat(tempData[1], 64)
		if parseErr != nil {
			panic(parseErr)
		}

		value, ok := cityMap[city]
		if !ok {
			cityMap[city] = OutputData{
				Min:   temperature,
				Mean:  temperature,
				Max:   temperature,
				Count: 1,
			}
			continue
		}
		value.Count++

		if temperature < value.Min {
			value.Min = temperature
		}

		if temperature > value.Max {
			value.Max = temperature
		}

		value.Mean = value.Mean + temperature

		cityMap[city] = value
	}

	var sData []OutputData
	for city, value := range cityMap {
		tV := OutputData{
			City:  city,
			Min:   value.Min,
			Mean:  value.Mean,
			Max:   value.Max,
			Count: value.Count,
		}

		sData = append(sData, tV)
	}

	// sort sData by city name
	sort.Slice(sData, func(i, j int) bool {
		return sData[i].City < sData[j].City
	})

	finalOutput := formatOutput(sData)

	fmt.Println(finalOutput)
}

func formatOutput(data []OutputData) string {
	finalOutput := "{"
	for _, value := range data {
		tempMean := value.Mean / float64(value.Count)
		mean := float64(int(tempMean*10)) / 10.0

		s := fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", value.City, value.Min, mean, value.Max)
		finalOutput += s
	}

	// remove last two chars
	finalOutput = finalOutput[:len(finalOutput)-2] + "}"
	return finalOutput
}

func openFile() (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
