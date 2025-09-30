package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	URL         string
	Name        string
	Category    string
	Price       string
	Rating      string
	Reviews     string
	Image       string
	Description string
}

func (p Product) String() string {
	return fmt.Sprintf(
		"Product    : %s\n"+
			"Name       : %s\n"+
			"Category   : %s\n"+
			"Price      : %s\n"+
			"Rating     : %s\n"+
			"Reviews    : %s\n"+
			"Image URL  : %s\n"+
			"Description: %s\n",
		p.URL, p.Name, p.Category, p.Price, p.Rating, p.Reviews, p.Image, p.Description,
	)
}

func GetProductURLs(sitemapURL string) ([]string, error) {
	defaultNS := "http://www.sitemaps.org/schemas/sitemap/0.9"
	res, err := http.Get(sitemapURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	decoder := xml.NewDecoder(res.Body)
	var urls []string
	skipFirst := true

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		element, conversionOK := token.(xml.StartElement)
		if conversionOK {
			if element.Name.Local == "loc" && element.Name.Space == defaultNS {
				if skipFirst {
					skipFirst = false
					decoder.DecodeElement(new(string), &element)
					continue
				}
				var loc string
				if err := decoder.DecodeElement(&loc, &element); err != nil {
					return nil, err
				}
				urls = append(urls, loc)
				if len(urls) >= 100 {
					break
				}
			}
		}
	}
	return urls, nil
}

func FetchProductDetails(productURL string) (*Product, error) {
	res, err := http.Get(productURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	product := &Product{}
	product.URL = productURL
	product.Name = doc.Find(".product__title h1").Text()
	product.Category = doc.Find(".product__title h2").Text()
	product.Price = doc.Find(".price-item .money").Text()
	product.Rating = doc.Find(".jdgm-prev-badge").AttrOr("data-average-rating", "")
	product.Reviews = doc.Find(".jdgm-prev-badge").AttrOr("data-number-of-reviews", "")
	product.Image = doc.Find(".product__media img.image-magnify-lightbox").AttrOr("src", "")
	product.Description = doc.Find(".product__description").Text()

	return product, nil
}

func main() {
	sitemapProductsURL := "https://raidlight.com/sitemap_products_1.xml?from=5762755920036&to=10169533137237"
	urls, err := GetProductURLs(sitemapProductsURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	// UTILISATION DE LA PUISSANCE DU GOROUTINES DE GO ...
	// ca ne respecte pas robots.txt car on fait trop de requetes en meme temps on doit attendre au moins 2 secondes entre chaque requete
	/*var waitGroup sync.WaitGroup
	for _, url := range urls {
		waitGroup.Add(1)
		go func(productURL string) {
			defer waitGroup.Done()
			product, err := FetchProductDetails(productURL)
			if err != nil {
				fmt.Println("Error fetching product details:", err)
				return
			}
			fmt.Print(product)
			fmt.Println("--------------------------------------------------")
		}(url)
	}
	waitGroup.Wait()*/

	// UTILISATION CLASSIQUE avec une attente de 2 secondes entre chaque requete
	for _, url := range urls {
		product, err := FetchProductDetails(url)
		if err != nil {
			fmt.Println("Error fetching product details:", err)
			continue
		}
		fmt.Print(product)
		fmt.Println("--------------------------------------------------")
		time.Sleep(2 * time.Second)
	}
}
