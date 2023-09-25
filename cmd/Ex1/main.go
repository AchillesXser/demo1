package main

import (
	"html/template"
	"os"
)

type UserMata struct {
	Visitors int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
		Age  int16
		Mata UserMata
		Bio  string
	}{
		Name: "john lee",
		Age:  12,
		Mata: UserMata{
			Visitors: 12,
		},
		Bio: `<script>alert("你被攻击了")</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
