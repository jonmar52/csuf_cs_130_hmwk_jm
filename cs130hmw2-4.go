package main

import "fmt"

type Full_name struct {
	first_name string
	last_name  string
}

func main() {
	var user_first_name string
	var user_last_name string

	user_first_name = "Jonathan"
	user_last_name = "Martin"

	var first *string = &user_first_name
	var last *string = &user_last_name

	var x = Full_name{*first, *last}

	fmt.Println(x.first_name, x.last_name)

}
