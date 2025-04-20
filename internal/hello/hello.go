package main

import (
	"fmt"
)

const spanish = "spanish"
const russian = "russian"
const portuguese = "portuguese"
const spanishHelloPrefix = "Hola, "
const russianHelloPrefix = "Привет, "
const portugueseHelloPrefix = "Olá, "
const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return helloPrefix(language) + name + "!"
}

func helloPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case russian:
		prefix = russianHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
