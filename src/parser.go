package crawler

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parser.go 解析页面内容

// LinkParse 根据Url获取网页中当前域名下的URL链接
func LinkParse(baseURL string) []LinkItem {

	doc, err := goquery.NewDocument(baseURL)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	u, _ := url.Parse(baseURL)
	site := u.Scheme + "://" + u.Host

	var items = make([]LinkItem, 16, 32)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		var item LinkItem
		item.URL, _ = s.Attr("href")
		if !strings.Contains(item.URL, "http") {
			item.URL = site + item.URL
		}
		if strings.Contains(item.URL, site) {
			item.Text = s.Text()
			items = append(items, item)
		}
	})
	return items
}
