package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"time"
	"unicode/utf8"

	"hello-go/gopl"
	_ "hello-go/gopl/array"
	_ "hello-go/gopl/slice"
	_ "hello-go/gopl/mymap"
	_ "hello-go/gopl/mystruct"

)


func main() {
	fmt.Println("hello go")

	playBit()

	playFloat()

	playBoolean()

	friendlyEcho(playString)

	playConst()

      gopl.Run()
}

type noParamFunc func()

func friendlyEcho(f noParamFunc) {
	fmt.Println(" <<  begin -----------------------  ")

	f()

	fmt.Println(" ----------------------- end >> ")
}

func playBit() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			// membership test
			fmt.Println(i) // "1", "5"
		}
	}

	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

func playFloat() {
	fmt.Println("浮点数最小值", math.MaxFloat64)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
}

func playBoolean() {

	var itob = func(intVar int) bool {
		return intVar != 0
		/*
			    if intVar != 0 {
				    return true
			    }
				return false */
	}
	var btoi = func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	fmt.Print(itob(3), itob(0), itob(-1))
	fmt.Println(btoi(true))

	if true {
		fmt.Println("hi i am a true condition!")
	}
}

func playString() {
	s := "hello, world"
	fmt.Println(len(s))

	fmt.Println(s[0], s[7])

	fmt.Println(s[0:5])

	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	fmt.Println("真的可以响耶！ \a \a \a")
	time.Sleep(600000)
	fmt.Println("真的可以响耶！ \a \a \a")

	const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`

	fmt.Println(GoUsage)

	func() {
		s := "hello , 世界"
		fmt.Println(len(s))
		fmt.Println(utf8.RuneCountInString(s))

		for i := 0; i < len(s); {
			r, size := utf8.DecodeRuneInString(s[i:])
			fmt.Printf("%d\t%c\n", i, r)
			i += size
		}

		for i, r := range "hello, 世界" {
			fmt.Printf("%d\t%c\n", i, r)
		}

		func(s string) {
			n := 0
			for range s {
				n++
			}
			fmt.Printf("字符串总数 %d", n)
		}(s)

	}()

	func() {
		// "program" in Japanese katakana
		s := "プログラム"
		fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
		r := []rune(s)
		fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

		fmt.Println(string(r))

		fmt.Println(string(65))     // "A", not "65"
		fmt.Println(string(0x4eac)) // "京"

		// 无效码点
		fmt.Println(string(1234567)) // "(?)"
	}()

	func() {
		type funcType func(string) string
		var commaFunc funcType
		commaFunc = func(s string) string {
			n := len(s)
			if n <= 3 {
				return s
			}
			return commaFunc(s[:n-3]) + "," + s[n-3:]
		}

		fmt.Println(commaFunc("hi i am here"))

	}()

	// 字符串和字节数组的相互转换
	func() {
		s := "abc"
		b := []byte(s)
		s2 := string(b)

		fmt.Println(s, s2)

	}()

	// 测试 bytes包下的函数
	func() {
		type tmpFunc func([]int) string
		var intsToString tmpFunc
		intsToString = func(values []int) string {
			var buf bytes.Buffer
			buf.WriteByte('[')
			for i, v := range values {
				if i > 0 {
					buf.WriteString(", ")
				}
				fmt.Fprintf(&buf, "%d", v)
			}
			buf.WriteByte(']')
			return buf.String()
		}
		fmt.Println(intsToString([]int{1, 2, 3}))
	}()

	// 转换
	func() {
		x := 123
		y := fmt.Sprintf("%d", x)       // 转字符串
		fmt.Println(y, strconv.Itoa(x)) // int 转为ascII
		fmt.Println(strconv.FormatInt(int64(x), 2))

		s := fmt.Sprintf("x=%b", x) // "x=1111011"
		fmt.Println(s)

		x, err := strconv.Atoi("123")
		y2, err := strconv.ParseInt("123", 10, 64)

		fmt.Println(x, y2, err)
	}()
}

func playConst() {
	// 常量都是基础类型
	const pi = 3.141592 //
	fmt.Println(pi)

	func() {
		// 打印外层的pi值
		fmt.Println(pi)
		// 常量的scope  可以覆盖哦
		const (
			e  = 2.71828182845904523536028747135266249775724709369995957496696763
			pi = 3.14159265358979323846264338327950288419716939937510582097494459
		)

		fmt.Println(e, pi)
	}()

	// 带类型的常量
	func() {
		const noDelay time.Duration = 0
		const timeOut = 5 * time.Minute

		fmt.Printf("%T %[1]v \n", noDelay)
		fmt.Printf("%T %[1]v \n", timeOut)
		fmt.Printf("%T %[1]v \n", time.Minute)

	}()

	// 批量声明的常量
	func() {
		const (
			a = 1
			b
			c = 2
			d
		)
		fmt.Println(a, b, c, d)
	}()

	// iota 常量生成器
	func() {
		type Weekday int

		const (
			Sunday Weekday = iota
			Monday
			Tuesday
			Wednesday
			Thursday
			Friday
			Saturday
		)
		fmt.Println(Sunday,
			Monday,
			Tuesday,
			Wednesday,
			Thursday,
			Friday,
			Saturday)

		type Flags uint
		const (
			FlagUp Flags = 1 << iota
			FlagBroadcast
			FlagLoopback
			FlagPointToPoint
			FlagMulticast
		)
		fmt.Println(FlagUp, "\n",
			FlagBroadcast, "\n",
			FlagLoopback, "\n",
			FlagPointToPoint)

		const (
			_   = 1 << (10 * iota)
			KiB // 1024
			MiB // 1048576
			GiB // 1073741824
			TiB // 1099511627776             (exceeds 1 << 32)
			PiB // 1125899906842624
			EiB // 1152921504606846976
			ZiB // 1180591620717411303424    (exceeds 1 << 64)
			YiB // 1208925819614629174706176
		)

		fmt.Println(YiB / ZiB) // "1024"
	}()
	// 不确定类型的常量
	func() {
		var x float32 = math.Pi
		var y float64 = math.Pi
		var z complex128 = math.Pi
		fmt.Println(x, y, z)
	}()

}
