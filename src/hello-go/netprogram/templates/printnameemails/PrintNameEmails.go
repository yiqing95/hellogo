package main

import (
	"hello-go/netprogram/util"
	"html/template"
	"os"
)

type Person struct {
	Name   string
	Emails []string
}

// 模板中使用变量
const templ = `{{$name := .Name}}
{{range .Emails }}
	Name is {{$name}}, email is {{.}}
{{end }}
`

func main() {
	person := Person{
		Name:   "jan",
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
	}

	t := template.New("Person template")
	t, err := t.Parse(templ)
	util.CheckError(err)

	err = t.Execute(os.Stdout, person)
	util.CheckError(err)
}
