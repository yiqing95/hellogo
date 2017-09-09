package array

import (
	"fmt"
	"hello-go/gopl"
	"crypto/sha256"
)

func init() {
	fmt.Println("hi am from array package ... ")
	// runIt()
	gopl.AddRunFunc(runIt)

}

var a [3]int // array of 3 integers

func runIt() {
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2], q)

	func() {
		q := [...]int{1, 2, 3}
		fmt.Printf("%T\n", q)
	}()

	func() {
		type Currency int
		const (
			USD Currency = iota // 美元
			EUR                 // 欧元
			GBP                 // 英镑
			RMB                 // 人民币
		)
		symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

		fmt.Println(RMB, symbol[RMB]) // "3 ￥"
	}()

	func() {
		r := [...]int{99: -1}
		fmt.Printf("the length of array is %d \n", len(r))
		fmt.Printf("the cap of array is %d \n", cap(r))
		fmt.Printf("the Type of array is %T \n",r)
	}()

	// 数组比较
	campareIt := func(){
		a := [2]int{1,2}
		b := [...]int{1, 2}
		c := [2]int{1,3}
		fmt.Println( a== b, b == c ,c == a)
		/*
		d := [3]int{1, 2}
		fmt.Println(a == d)
		*/
	}
	campareIt()

	func(){
		c1 := sha256.Sum256([]byte("x"))
		c2 := sha256.Sum256([]byte("X"))

		fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1==c2 , c1)
	}()
}
