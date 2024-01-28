package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const hindi = "Hindi"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const hindiHelloPrefix = "Namaste, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Nihar", ""))
}
