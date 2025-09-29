package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Hello, Go!")
	res, err := http.Get("https://www.google.com/")

	fmt.Println(res.Status)
	fmt.Println(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	title := doc.Find("title").Text()
	fmt.Println("Titre de la page:", title)

}
