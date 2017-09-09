package slice

import (
	"hello-go/gopl"
	"fmt"
)

func init()  {
	gopl.AddRunFunc(runIt)
}

func runIt()  {
	s := "hello my string"
	s2 := s[5:]
	fmt.Println(s2)

	b := []byte("hi am a byte array")
	b2 := b[4:]
	fmt.Printf("%T",b2)
	fmt.Println(b2)
	fmt.Println(string(b2))

	myI := [...]int{1,3,4,5,6,7,8}
	// fmt.Println(reverse(myI))
	reverse(myI[:])
	fmt.Println(myI)
	// 左旋2
	func(){
		s := []int{0,1,2,3,4,5,6,7,8}
		// 先反转前2个元素
		reverse(s[:2])
		// 再反转剩余的
		reverse(s[2:])
		// 再反转整体
		reverse(s)
		fmt.Println(s)

	}()

	// make 创建
	func(){
		s := make([]int, 0)
		s2 := make([]int, 2 , 3)
		fmt.Println(len(s2), cap(s2))
		fmt.Println(len(s), cap(s))

	}()

	// append 函数
	func(){
		var runes []rune
		for _, r := range "hello 世界"{
			runes = append(runes,r)
		}
		fmt.Printf("%q\n", runes)
	}()

	// append内部原理
	func(){
		var x, y []int
		for i :=0 ; i<0 ; i++{
			y = appendInt(x, i)
			fmt.Printf("%d cap=%d \t %v \n", i ,cap(y), y)
			x = y
		}
	}()

	// 追加多个元素
	func(){
		var x []int
		x = append(x, 1)
		x = append(x, 2, 3)
		x = append(x, 4, 5, 6)
		x = append(x, x...) // append the slice x
		fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
	}()
}

func reverse(s []int)  {
	for i, j := 0 ,len(s)-1; i <j ; i,j=i+1, j-1{
		s[i] ,s[j] = s[j], s[i]
	}
}


func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}
