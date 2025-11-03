package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  int
}
type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       string
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Lucy",
			LastName: "Duss",
		},
		BirthDate: BirthDate{
			Day: 05,
			Month: 01,
			Year: 2007,
		},
		NumberOfSiblings: 0,
		ZodiacSign:       "Capricorn",
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
