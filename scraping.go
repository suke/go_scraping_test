package main

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/yuin/charsetutil"
)

func main() {
	doc, err := goquery.NewDocument("http://www.creation.gr.jp/")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	doc.Find("div.pickup-topic-label").Each(func(_ int, s *goquery.Selection) {
		text := s.Text()
		text, err = charsetutil.DecodeString(text, "Windows-31J")
		if err != nil {
			fmt.Print("text decode failed")
		}
		r := regexp.MustCompile(`\d{1,2}年\d{1,2}月\d{1,2}日`)
		fmt.Println(r.FindAllStringSubmatch(text, -1))
	})
}
