package main

import (
	"fmt"

	"example.com/struct-project/note"
)

func main() {
	title, content := getNoteData()
	note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	note.Show()
	err = note.Save()

	if err != nil {
		fmt.Println("failed to save note")
		return
	}

	fmt.Println("note saved")
}
