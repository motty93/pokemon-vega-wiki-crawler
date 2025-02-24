package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	allUrl = "https://w.atwiki.jp/altair1/pages/19.html" // ポケモン図鑑のURL
)

// 図鑑リストを作成する
func main() {
	res, err := http.Get(allUrl)
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

	doc.Find("table tbody tr").Each(func(i int, row *goquery.Selection) {
		row.Find("td").Each(func(j int, cell *goquery.Selection) {
			text := strings.TrimSpace(cell.Text()) // No.を取得
			link := cell.Find("a")                 // <a>タグを取得
			href, exists := link.Attr("href")
			if text == "No." || text == "名前" || text == "" {
				return
			}

			// tdループのindexを表示
			fmt.Printf("index: %d\n", j)
			if exists {
				fmt.Printf("名前: %s, URL: %s\n\n", text, formatURL(href))
			} else {
				fmt.Printf("No.%s\n", text)
			}
		})
	})
}

// hrefをhttpsへ変換
func formatURL(href string) string {
	if strings.HasPrefix(href, "//") {
		return "https:" + href
	} else if strings.HasPrefix(href, "/") {
		return "https://w.atwiki.jp" + href
	}

	return href
}
