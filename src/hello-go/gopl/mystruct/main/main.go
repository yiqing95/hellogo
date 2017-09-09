package main

import (
	"fmt"
	_ "hello-go/gopl/mystruct"
)


type Point struct{ X, Y int }

func main(){

p1 := Point{}
fmt.Println("default member value :" , p1)


// 包内部使用 不推荐api客户方用这种形式 一旦结构体成员定义顺序变更 下面的就不能用了
// 隐式使用未导出成员也是编译不通过的 /*package xPkg*/ type MyT struct { a, b int } // 另一个包中使用： var _ := xPkg.MyT{1,2} 
p := Point{1,2}
fmt.Println(p.X, p.Y)


p3 := Point{Y:4}

fmt.Println(p3.X, p3.Y)

// 结构体作为参数 和函数返回值
fmt.Println(Scala(Point{1,2}, 5))

//  错误用法： fmt.Println(Bonus(Emplyoee{Name:"yiqing", 1000},12))
fmt.Println(Bonus(&Emplyoee{Name:"yiqing", Salary: 1000},12))

	// 指针
	pp := &Point{1,2}

	// 等价下面
	pp2 := new(Point)
	*pp2 = Point{1,2}

	fmt.Println(pp, pp2)

	// ## 比较
	// ----------------------------------------  +| 
	// 可比较性  如果结构体所有成员可比较 那么结构体可比较
	func(){
		p := Point{1,2}
		q := Point{2,1}
		fmt.Println(p.X== q.X && p.Y == q.Y)
		fmt.Println(p == q)
		fmt.Println(p == p)
		// 可比较结构体做为map的key
		type address struct{
			hostname string
			port int
		}

		hits := make(map[address]int)
		hits[address{"golang.org", 443}] ++
		fmt.Println(hits)
	}()
	// ----------------------------------------  +/| 
}

func Scala(p Point, factor int) Point{
	return Point{p.X * factor, p.Y * factor}
}

type Emplyoee struct {
	Name string
	 Salary int
	}
// 指针形式
func Bonus(e *Emplyoee, percent int) int {
	return e.Salary * percent / 100
}
// 修改结构体变量
func AwardAnnualRaise(e *Emplyoee){
	e.Salary = e.Salary * 105/100
}