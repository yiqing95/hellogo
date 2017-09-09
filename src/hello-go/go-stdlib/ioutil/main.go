package main


import(
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"os"
)

func main(){
	r := strings.NewReader("Go is a genral-language designed with system programing in mind ")

	b , err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s" , b)

	readFile()

	tmpDir()
}


func readFile(){
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}

func tmpDir(){
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)
	fmt.Println(dir)
}