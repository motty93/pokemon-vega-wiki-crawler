package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://w.atwiki.jp/altair1/pages/71.html"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Failed to fetch the page: %d %s", res.StatusCode, res.Status)
	}

	// HTMLをパース
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("取得したカラム名:")
	doc.Find("table tbody").Each(func(i int, s *goquery.Selection) {
		text := s.Find("tr td").Text()
		fmt.Println(text)
	})
}
