package main

import "fmt"

type human struct {
	age  int
	name string
	sex  string
}

func Study(person human) {

	working, _, _ := StudyMessage(person.age, person.name, person.sex)

	fmt.Println(working)
}

func StudyMessage(age int, name string, sex string) (string, string, int) {

	return name + " is a hard working ", sex + "/n", age
}

func main() {
	var x int = 22
	var w = human{x, "Jonathan", "Male"}
	Study(w)
	fmt.Print(x)
	fmt.Print(" year old.")

}
