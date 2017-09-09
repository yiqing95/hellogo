package util

import (
	"fmt"
	"os"
)

// CheckError 检查错误
func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
