package main

import (
	"fmt"
)

const Title = "Person Details"

var Country = "USA"

func main() {
	fname, lname := "qing", "yi"

	age := 35
	fmt.Println(Title)

	fmt.Println("first name is ", fname, " and the last name is", lname)

	fmt.Println("Age : ", age)

	fmt.Println("Country: ", Country)
}
