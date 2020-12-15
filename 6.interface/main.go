package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func (englishBot) getGreeting() string {
	return "Hello World"
}
func (spanishBot) getGreeting() string {
	return "Hola World"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
