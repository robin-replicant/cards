package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an englisg greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating an spanish greeting
	return "Hola"
}

/*
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}


func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/
