package main

import (
	"fmt"
	"hello-go/go-stdlib/sync/math"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("init func called!")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	fmt.Println(strings.Repeat("--", 60))
	math.Run()
}

func foo() {
	for i := 0; i <= 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 45; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(53 * time.Millisecond))
	}
	wg.Done()
}
