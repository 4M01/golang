package main

import "fmt"

func main() {
	fmt.Println("Calling function - " + Hello("Vijay", "fr"))

}

// func Hello(name, language string) string {
// 	if language == "es" {
// 		return "Hola, " + name
// 	}

// 	if language == "fr" {
// 		return "Bonjour, " + name
// 	}

// 	return "Hello, " + name

// }

func Hello(name, language string) string {
	return fmt.Sprintf(
		"%s, %s",
		greeting(language),
		name,
	)
}

var greetings = map[string]string{
	"es": "Hola",
	"fr": "Bonjour",
}

func greeting(language string) string {
	greeting, exist := greetings[language]

	if exist {
		return greeting
	}

	return "Hello"
}
