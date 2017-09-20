package main

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"hello-go/Features/file/current_path/widgets"
)

// @see https://www.crifan.com/go_language_get_current_file_filename/
func main() {
	_, filename, line, _ := runtime.Caller(0)
	fmt.Println(filename)
	fmt.Println(line)

	fmt.Printf(" the base is : %s  \n", path.Base(filename))
	fmt.Printf(" the dir is : %s  \n", path.Dir(filename))
	fmt.Printf(" the ext is : %s  \n", path.Ext(filename))

	var currentDir string
	currentDir, err := os.Getwd()
	fmt.Println("currentDir=%s", currentDir)
	if err != nil {
		fmt.Println("get current directory error=%s\n", err)
	}
	// todo  暂时好像不能实现！
	widget := widgets.HelloWidget{}
	fmt.Println("widget path :", widget.ViewPath())
}
