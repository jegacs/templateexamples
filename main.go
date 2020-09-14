package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
)

type Info struct {
	Info string
	Link string
	Variable bool
}

func main() {
	fmt.Println("la clasica")
	t, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	info := Info{
		Info: "hola mariiiii",
		Link: "www.avenidasiempreviva.com",
		Variable: true,
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, info); err != nil {
		panic(err)
	}

	//message := buf.String()


	//fmt.Println(message)

	ioutil.WriteFile("output.html", buf.Bytes(), 0644)
}
