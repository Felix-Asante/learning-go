package main

func main() {
	todos := Todos{}
	storage := Storage[Todos]{"todos.json"}
	storage.load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.save(todos)
}
