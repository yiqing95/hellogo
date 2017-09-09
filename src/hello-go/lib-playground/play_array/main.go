package main

import "fmt"

func main()  {

	p := make([][]byte,5)
	// s := "hi  hello ya "
        /*
	format := ""
	if _, err := fmt.Sscanf(s, format, &p[0], &p[1], &p[2], &p[3], &p[4]); err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	*/
	fmt.Printf("p[0]:%p p[1]:%p\n",p[0],p[1])

	p[1]  =  []byte("hi")
	p[2]  =  []byte("hi")
	p[3]  =  []byte("hi")
	p[4]  =  []byte("hi")

	fmt.Println(p)

	fmt.Printf("p[0]:%p p[1]:%p\n",p[0],p[1])

	var a int
	fmt.Printf("%p", a)
	/*
	fmt.Sscanf("123","%d",&a)
	fmt.Printf("%p","123")
	fmt.Printf("%p", a)*/
	fmt.Println(a)
}
