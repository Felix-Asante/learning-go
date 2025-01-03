package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Edit   string
	Delete int
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify the new title. id:new_title")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	// parse the command line into the defined flags
	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todo *Todos) {
	// fmt.Println(flag.Args()) // returns the unFlagged arguments
	switch {
	case cf.List:
		todo.print()
	case cf.Add != "":
		todo.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Invalid edit format. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index")
			os.Exit(1)
		}
		todo.edit(index, parts[1])
	case cf.Delete != -1:
		todo.delete(cf.Delete)
	case cf.Toggle != -1:
		todo.toggle(cf.Toggle)

	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}
