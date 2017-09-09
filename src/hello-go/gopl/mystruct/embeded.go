package mystruct

import (
	"fmt"
)

type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

func init() {
	fmt.Println("// 测试结构体内嵌 --------------------")
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20 // 辐条数
	fmt.Println("测试下轮子:", w)

	refactor()

	anonymStructMembers()

	fmt.Println("// 测试结构体内嵌 ---- 完毕！ \n ")
}

func refactor() {
	type Point struct {
		X, Y int
	}

	type Circle struct {
		Center Point
		Radius int
	}

	type Wheel struct {
		Circle Circle
		Spokes int
	}
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20

	fmt.Println(w)
}

// 点
type Point struct {
	X, Y int
}

// 匿名成员
func anonymStructMembers() {
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}

	w := Wheel{
		Circle: Circle{}, Spokes: 10}

	fmt.Println(w)
	fmt.Println("等价访问：", w.Circle.X, w.X)

	// 不幸的是，结构体字面值并没有简短表示匿名成员的语法
	// w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	// w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}

	w.X = 42

	fmt.Printf("%#v\n", w)
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}
}
