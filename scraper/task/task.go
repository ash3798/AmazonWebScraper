package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/ash3798/AmazonWebScraper/scraper/config"
	colly "github.com/gocolly/colly/v2"
)

const (
	post = "POST"
)

type Product struct {
	Name         string
	ImageURL     string
	Description  string
	Price        string
	TotalReviews int
}

type ProductInfo struct {
	Url     string
	Product Product
}

type UrlInfo struct {
	Url string `json:"url"`
}

func ScrapeAndSend(url string) {
	log.Println("url received is : ", url)

	productInfo := doScrape(url)

	data, err := json.Marshal(productInfo)
	if err != nil {
		log.Println("error while marsheling the scrape results. Error : ", err.Error())
	}

	err = sendToPersist(data)
	if err != nil {
		log.Println("error while sending the scrape results. Error : ", err.Error())
	}
}

func doScrape(visitLink string) ProductInfo {
	product := Product{}

	c := colly.NewCollector(
		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		colly.CacheDir("./amazonps_cache"),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting ", r.URL.String())
	})

	c.OnHTML(`#titleSection`, func(e *colly.HTMLElement) {
		title := e.ChildText("#productTitle")
		//log.Println(title)
		product.Name = title
	})

	c.OnHTML(`.a-unordered-list.a-nostyle.a-horizontal.list.maintain-height`, func(e *colly.HTMLElement) {
		imageLink := e.ChildAttr(`.a-dynamic-image`, "src")
		//log.Println(imageLink)
		product.ImageURL = imageLink
	})

	c.OnHTML(`#price`, func(e *colly.HTMLElement) {
		price := e.ChildText(`#priceblock_ourprice`)
		//log.Println(price)
		product.Price = price
	})

	c.OnHTML(`div#feature-bullets > ul`, func(e *colly.HTMLElement) {
		description := ""

		e.ForEach(`span.a-list-item`, func(index int, elem *colly.HTMLElement) {
			if index < 1 {
				return
			}
			if elem.Text != " " {
				temp := strings.Replace(elem.Text, "\n", "", -1)
				description += temp
				description += "."
			}
		})
		//log.Println(description)
		product.Description = description
	})

	numReviewsCaptured := false
	c.OnHTML(`a#acrCustomerReviewLink`, func(e *colly.HTMLElement) {
		if !numReviewsCaptured {
			reviewCount := e.ChildText(`span#acrCustomerReviewText`)

			reviewCount = strings.Replace(reviewCount, ",", "", -1)
			r, _ := regexp.Compile(`^[\d]+`)

			reviewTotal, err := strconv.Atoi(r.FindString(reviewCount))
			if err != nil {
				log.Println("error while converting string to int, Error: ", err.Error())
			}
			//log.Println(reviewTotal)

			product.TotalReviews = reviewTotal
			numReviewsCaptured = true
		}
	})

	c.Visit(visitLink)

	productInfo := ProductInfo{Url: visitLink, Product: product}
	log.Printf("product info scraped : %+v", productInfo)

	return productInfo
}

func sendToPersist(data []byte) error {
	resp, err := http.Post(config.Manager.PersistServiceURL, http.DetectContentType(data), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed status code returned for request")
	}
	log.Println("succesfully sent the scrape results to persist")
	return nil
}
