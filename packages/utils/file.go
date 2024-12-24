package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteValueToFile(filename string, value float64) {
	error := os.WriteFile(filename, []byte(fmt.Sprintf("%.2f", value)), 0644)

	if error != nil {
		fmt.Println("Error writing to file")
	}
}

func GetFloatFromFile(filename string) (float64, error) {
	data, error := os.ReadFile(filename)
	if error != nil {
		return 0, errors.New("failed to find file")
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}
	return value, nil
}
