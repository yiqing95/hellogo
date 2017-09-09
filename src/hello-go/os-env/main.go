package main

import (
	"fmt"
	"os" //我们要用到os包中的env
	"strings"
)

func main() {
	//os.Getenv检索环境变量并返回值，如果变量是不存在的，这将是空的。
	HOME := os.Getenv("HOME")
	fmt.Println(HOME)
	fmt.Printf(os.Getenv("JAVA_HOME"))

	var path string = os.Getenv("Path")
	fmt.Println(path)

	if strings.Contains(path, "Dwimperl") {
		fmt.Println("包含呀")
	} else {
		fmt.Println("不包含!")
	}
}
