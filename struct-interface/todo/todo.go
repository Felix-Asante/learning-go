package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"title"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: text,
	}, nil
}

func (t Todo) Show() {
	fmt.Printf("Todo text:%v\n", t.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	writeError := os.WriteFile(fileName, json, 0644)

	if writeError != nil {
		return writeError
	}
	return nil
}
