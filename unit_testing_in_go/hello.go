/*
(c)Nokia 2024
Author: Maciej Norberciak
companion code for "Unit testing in Go" talk

Test Dive Conference, Krakow, 17.10.2024
*/

package main

func HelloName(language string, name string) string {

	greeting := ""

	switch language {
	case "english":
		greeting = "Hello"
	case "spanish":
		greeting = "Hola"
	case "german":
		greeting = "Hallo"
	case "polish":
		greeting = "Cześć"
	case "japanese":
		greeting = "今日は"
	default:
		greeting = ""
	}

	return greeting + " " + name

}
