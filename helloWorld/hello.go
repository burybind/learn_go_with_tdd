package main

import "fmt"

const (
	englishHello = "Hello"
	spanishHello = "Hola"
	frenchHello  = "Bonjour"
)

type languages int

const (
	english = iota
	spanish
	french
)

func main() {
	fmt.Println(Hello("Brendan", english))
}

func Hello(name string, language languages) string {
	if name == "" {
		name = "World"
	}

	var prefix string
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}

	return fmt.Sprintf("%s, %s", prefix, name)
}
