package main

import "fmt"
import "reflect"

func main() {
	var x [128]int

	fmt.Println(x)

	xp := new([128]int)
	fmt.Println(xp)

	s := make([]int, 50, 100)
	_ = s

	s2 := new([100]int)[0:50]
	fmt.Println(s2)
	fmt.Println(reflect.TypeOf(s2))

	b := [5]int{-1, -2, -3, -4, -5}
	s = b[1:4]
	fmt.Println(s)
}
