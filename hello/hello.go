package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and pring it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
