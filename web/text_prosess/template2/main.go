package main

import (
	"fmt"
	"html/template"
	"os"
)

// Friend struct
type Friend struct {
	Fname string
}

// Person struct
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	t := template.New("fieldname example")

	t, _ = t.Parse(`hello {{.UserName}}!
            {{range .Emails}}
                an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
                my friend name is {{.Fname}}
            {{end}}
			{{end}}
			`)

	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	p := Person{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}

	t.Execute(os.Stdout, p)
	fmt.Println(p.Emails[0])
	fmt.Println(*(p.Friends[0]))
}
