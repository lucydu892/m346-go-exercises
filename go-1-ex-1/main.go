package main

import "fmt"

func main() {
	firstName := "Lucy"
	lastName := "Duss"
	dayOfBirth := 05
	monthOfBirth := 01
	yearOfBirth := 2007
	numberOfSiblings := 0
	heightInMeters := 1.75
	zodiacSign := "Capricorn"

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
