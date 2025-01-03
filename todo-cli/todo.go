package main

import (
	"errors"
	"strconv"
	"time"

	"os"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	completedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		completedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	err := todos.validateIndex(index)
	if err != nil {
		return err
	}

	t := *todos
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}
	isCompleted := (*todos)[index].Completed

	if !isCompleted {
		completedTime := time.Now()
		(*todos)[index].completedAt = &completedTime
		(*todos)[index].Completed = true
	}
	(*todos)[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = title

	return nil
}

func (todos *Todos) print() error {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for i, todo := range *todos {
		completed := "❌"
		completedAt := ""

		if todo.Completed {
			completed = "✅"
			if todo.completedAt != nil {
				completedAt = todo.completedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()

	return nil
}
