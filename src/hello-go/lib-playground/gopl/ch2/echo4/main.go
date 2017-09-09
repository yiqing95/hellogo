package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("m" , false ,"omit trailing newline")
var sep = flag.String("s", " ",  "seprarator ")

func main()  {
	flag.Parse()
	fmt.Print(strings.Join( flag.Args() , *sep))
	if(! *n){
		fmt.Println( )
	}
}
