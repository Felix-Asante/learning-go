package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, error := os.Open(path)

	if error != nil {
		fmt.Println("Failed to open file", error)
		return nil, error
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	error = scanner.Err()

	if error != nil {
		file.Close()
		fmt.Println("Error reading file")
		return nil, error
	}

	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, error := os.Create(path)

	if error != nil {
		file.Close()
		return errors.New("failed to create file")
	}

	error = json.NewEncoder(file).Encode(data)

	if error != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil
}
