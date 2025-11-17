package main

import "fmt"

type Student struct {
	firstname string
	lastname  string
}
type Class struct {
	Students []Student
}
type Module struct {
	Number  int
	Classes []Class
}

func main() {
	// TODO: declare a type for Student (with first and last name)

	class1 := Class{
		Students: []Student{
			{firstname: "Jessica", lastname: "Meier"},
			{firstname: "Leo", lastname: "Holzer"},
			{firstname: "Carol", lastname: "Pfister"},
		},
	}
	class2 := Class{
		Students: []Student{
			{firstname: "Peter", lastname: "Christen"},
			{firstname: "Sarah", lastname: "Duss"},
			{firstname: "Noel", lastname: "Hermann"},
		},
	}
	// TODO: declare a type for Class (consisting of multiple students)

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int]Module{
		346: {
			Number:  346,
			Classes: []Class{class1},
		},
		104: {
			Number:  104,
			Classes: []Class{class1, class2},
		},
		117: {
			Number:  117,
			Classes: []Class{class2},
		},
	}
	// TODO: output everything using fmt.Println()
	for moduleNumber, module := range modules {
		fmt.Printf("Modul %d:\n", moduleNumber)
		for i, class := range module.Classes {
			fmt.Printf("  Klasse %d:\n", i+1)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.firstname, student.lastname)
			}
		}
		fmt.Println()
	}
}
