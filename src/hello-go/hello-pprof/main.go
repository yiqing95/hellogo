package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:13001", nil)
	}()

	fmt.Println("运行啦！")
}
