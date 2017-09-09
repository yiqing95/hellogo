package main

import (
	"fmt"
	"hello-go/lib-playground/gopl/ch2/conf"
)

const SEX_MAN = 1
const SEX_FEMALE = 0

func main()  {

	fmt.Println("hi")
	fmt.Println(SEX_FEMALE)

	fmt.Println(conf.DB_NAME)
}
