package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{
		FileName: fileName,
	}
}

func (t *Storage[T]) save(data T) error {
	fileData, error := json.MarshalIndent(data, "", "  ")
	if error != nil {
		return error
	}

	return os.WriteFile(t.FileName, fileData, 0644)

}

func (s *Storage[T]) load(data *T) error {
	fileData, error := os.ReadFile(s.FileName)
	if error != nil {
		return error
	}

	return json.Unmarshal(fileData, &data)
}
