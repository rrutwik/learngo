// This is the main package for the simple Go program which prints Hello
// world to the terminal.
package main

import "fmt"

var helloWorldString = "Hello, "

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Hello returns a friendly greeting.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloWorldString + name
}

var name = "Rutwik"

// main is the entry point for the Go program.
func main() {
	fmt.Println(Hello(name))
}
