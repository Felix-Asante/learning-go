package main

import (
	"fmt"
	"time"
)

func greet(message string, done chan bool) {
	fmt.Println(message)
	done <- true
}

func slowGreet(message string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(message)
	doneChan <- true
}

func main() {
	done := make(chan bool)
	go greet("Hello", done)
	go slowGreet("He...l..lo there", done)
	go greet("Bye", done)
	<-done
}
