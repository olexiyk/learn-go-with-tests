package main

const spanish = "sp"
const french = "fr"
const english = "en"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" {
		language = "World"
	}
	greeting := "Hello"
	switch language {
	case spanish:
		greeting = "Hola"
	case french:
		greeting = "Bonjour"
	}

	return greeting + ", " + name
}
