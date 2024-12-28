package main

import (
	"fmt"

	"example.com/struct-interface/note"
	"example.com/struct-interface/todo"
)

type saver interface {
	Save() error
}

func saveData(s saver) {
	err := s.Save()

	if err != nil {
		fmt.Println(err)
		return
	}
}

type displayer interface {
	Show()
	saver
}

func outputData(data displayer) {
	data.Show()
	saveData(data)
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	note, err := note.New(title, content)
	todo, todoError := todo.New(todoText)

	if err != nil || todoError != nil {
		fmt.Println(err)
		return
	}

	// saveData(note)
	// saveData(todo)
	outputData(note)
	outputData(todo)
}
