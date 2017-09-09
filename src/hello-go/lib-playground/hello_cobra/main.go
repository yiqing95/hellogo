package main

import (
	"fmt"
	"os"

	"hello-go/lib-playground/hello_cobra/cmd"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
