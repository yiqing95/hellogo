package gopl

import "fmt"

type (
	RunFunc func()

)
var runFuncList []RunFunc

func AddRunFunc(runFunc RunFunc){
   runFuncList = append(runFuncList,runFunc)
}

func Run()  {
	fmt.Println("hi I am running ")
	// 运行钩子方法集合
	for idx , runFunc := range runFuncList{
		fmt.Print(idx)
		runFunc()
	}
}