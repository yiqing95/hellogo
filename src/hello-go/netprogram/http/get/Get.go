package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	url := os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	// 这段检测Content-Type 的代码有bug ！有的http响应就没有编码提示 可能默认应该就是utf-8 但函式中没检测
	/*
		fmt.Println("the header is ", response.Header)
		contentTypes := response.Header["Content-Type"]
		if !acceptableCharset(contentTypes) {
			fmt.Println("Cannot handle", contentTypes)
			os.Exit(4)
		}
	*/

	fmt.Println("The response body is")
	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
	os.Exit(0)
}

func acceptableCharset(contentTypes []string) bool {
	fmt.Println("the content-type is :", contentTypes)
	// each type is like [text/html; charset=utf-8]
	// we want the UTF-8 only
	for _, cType := range contentTypes {
		if strings.Index(cType, "utf-8") != -1 {
			return true
		}
	}
	return false
}
