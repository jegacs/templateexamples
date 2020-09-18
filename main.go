package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
)

type Provider struct {
	Name            string
	WhithoutAccount bool
	ProviderError   bool
	Value           float64
}

type Info struct {
	Organization string
	ReportType   string
	Providers    []Provider
}

func main() {
	fmt.Println("generando archivo")
	t, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	info := Info{
		Organization: "paquititous.com",
		ReportType:   "Semanal",
		Providers: []Provider{
			Provider{
				Name:            "shopify",
				WhithoutAccount: false,
				ProviderError:   false,
				Value:           8.5,
			},
			Provider{
				Name:            "google",
				WhithoutAccount: false,
				ProviderError:   false,
				Value:           -15.6,
			},
			Provider{
				Name:            "facebook",
				WhithoutAccount: false,
				ProviderError:   false,
				Value:           18.1,
			},
		},
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, info); err != nil {
		panic(err)
	}

	//message := buf.String()

	//fmt.Println(message)

	ioutil.WriteFile("output.html", buf.Bytes(), 0644)
}
