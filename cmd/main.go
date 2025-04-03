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
	fileLines = 1000000000
)

type TempTemp struct {
	Min   float64
	Mean  float64
	Max   float64
	Count int
}

type testOutput struct {
	City string
	Min  float64
	Mean float64
	Max  float64
}

func main() {
	// 1. load the input data from the data file
	file, err := openFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tempMap := make(map[string][]float64, fileLines)
	outputData := make(map[string]TempTemp, fileLines)

	scanner := bufio.NewScanner(file)

	lineCounter := 0
	for scanner.Scan() {
		lineCounter++
		line := scanner.Text()

		// slip the line into two parts by ;
		temp := strings.Split(line, ";")
		city := temp[0]

		temperatureStr := temp[1]
		if temperatureStr == "" {
			panic(nil)
		}

		if len(temperatureStr) < 3 {
			continue
		}

		temperature, parseErr := strconv.ParseFloat(temp[1], 64)
		if parseErr != nil {
			panic(parseErr)
		}

		// TODO: we might be able to calculate the min, max and mean temperature here
		tempMap[city] = append(tempMap[city], temperature)

		value, ok := outputData[city]
		if !ok {
			outputData[city] = TempTemp{
				Min:   temperature,
				Mean:  temperature,
				Max:   temperature,
				Count: 1,
			}
			continue
		}

		if temperature < value.Min {
			value.Min = temperature
		}

		if temperature > value.Max {
			value.Max = temperature
		}

		value.Count++

		value.Mean = (value.Mean + temperature) / float64(value.Count)

		outputData[city] = value
	}

	// sort outputData by city name
	tData := []testOutput{}
	for city, value := range outputData {
		tV := testOutput{
			City: city,
			Min:  value.Min,
			Mean: value.Mean,
			Max:  value.Max,
		}

		tData = append(tData, tV)
	}

	// sort tData by city name
	sort.Slice(tData, func(i, j int) bool {
		return tData[i].City < tData[j].City
	})

	outputCounter := 0
	finalOutput := "{"
	for _, value := range tData {
		outputCounter++
		// round up mean temperature to 1 decimal
		mean := float64(int(value.Mean*10)) / 10.0

		v := fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", value.City, value.Min, mean, value.Max)
		finalOutput += v
	}

	// remove last two chars
	finalOutput = finalOutput[:len(finalOutput)-2] + "}"

	fmt.Println(finalOutput)
}

func openFile() (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}
