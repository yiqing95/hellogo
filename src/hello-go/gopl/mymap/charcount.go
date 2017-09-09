package mymap

import (
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"fmt"
	"unicode"
)

func mymain(){
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax+1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for{
		r, n , err := in.ReadRune()
		if(err == io.EOF){
			break
		}
		if(err != nil){
			fmt.Fprint(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid ++
			continue
		}
		counts[r] ++
		utflen[n] ++

	}
	fmt.Print("rune\tcount\n")
	for c, n := range counts{
		fmt.Printf("%q\t%d", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen{
		if i >0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid >0 {
		fmt.Printf("\n%d invalidate UTF-8 characters \n ", invalid)
	}
}
