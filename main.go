package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func retrieveSiteMapURL(websiteUrl string) string {
	if !strings.HasSuffix(websiteUrl, "/") {
		websiteUrl += "/"
	}
	websiteUrl += "robots.txt"

	resp, err := http.Get(websiteUrl)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(err, " ", resp.Status)
		return ""
	}
	scanner := bufio.NewScanner(resp.Body)
	var sitemapURL string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fmt.Println((line))
		if strings.HasPrefix(strings.ToLower(line), "sitemap:") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				sitemapURL = strings.TrimSpace(parts[1])
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lecture :", err)
	}
	return sitemapURL
}

func main() {

	sitemapUrl := retrieveSiteMapURL("https://facebook.com/")
	fmt.Println(sitemapUrl)
	fmt.Println("Hello, Go!")
	res, err := http.Get("https://raidlight.com/sitemap.xml")

	fmt.Println(res.Status)
	fmt.Println(err)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	title := doc.Find("title").Text()
	fmt.Println("Titre de la page:", title)

}
