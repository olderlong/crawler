package crawler

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parser.go 解析页面内容

// LinkParse 根据Url获取网页中当前域名下的URL链接
func PageParse(resp *http.Response) {

	doc, err := goquery.NewDocumentFromResponse(resp) //这里已经将resp关闭了，对resp的处理要放在前面
	if err != nil {
		log.Fatal(err)
		return
	}
	// u ,_:=resp.Location()	//location函数从resp的Header中读取Location字段的值，有时会为空，这里不采用
	u := resp.Request.URL
	site := u.Scheme + "://" + u.Host

	for _, parseConfig := range ParseConfigs {
		if strings.Contains(site, parseConfig.URLPattern) {
			ParseOne(doc, site, parseConfig)
		}
	}
	// var items = make([]LinkItem, 16, 32)

	// doc.Find("a").Each(func(_ int, s *goquery.Selection) {
	// 	var item LinkItem
	// 	item.URL, _ = s.Attr("href")
	// 	if !strings.Contains(item.URL, "http") {
	// 		item.URL = site + item.URL
	// 		log.Println(item.URL)
	// 	}
	// 	if strings.Contains(item.URL, site) {
	// 		item.Text = s.Text()
	// 		items = append(items, item)
	// 	}
	// })

	// return items
}
func ParseOne(doc *goquery.Document, site string, parseConfig ParseConfig) {
	for _, rule := range parseConfig.Rules {
		if parseConfig.Rules[0].ItemName == "Link" {
			LinksParse(doc, site, rule)
			continue
		}
		ItemParse(doc, site, rule)
	}
}
func LinksParse(doc *goquery.Document, site string, rule Rule) {

	var items = make([]LinkItem, 16, 32)

	doc.Find(rule.CSSSelector).Each(func(_ int, s *goquery.Selection) {
		var item LinkItem
		item.URL, _ = s.Attr(rule.AttrName)
		if !strings.Contains(item.URL, "http") {
			item.URL = site + item.URL
			log.Println(item.URL)
		}
		if strings.Contains(item.URL, site) {
			item.Text = s.Text()
			items = append(items, item)
		}
		log.Println(item.URL)
		URLQueue.PushBack(item.URL)
	})

	// return items
}
func ItemParse(doc *goquery.Document, site string, rule Rule) []map[string]string {

	var items = make([]map[string]string, 16, 32)

	doc.Find(rule.CSSSelector).Each(func(_ int, s *goquery.Selection) {
		var item map[string]string
		item["ItemName"] = rule.ItemName

		item["Value"] = s.Text()
		url, _ := s.Attr("href")
		if !strings.Contains(url, "http") {
			item["URL"] = site + url
		}
		item["Value"] = s.Text()
		items = append(items, item)
	})
	return items
}
func GetRespFromQueue() *http.Response {
	if ResponseQueue.Len() > 0 {
		item := ResponseQueue.Front()
		ResponseQueue.Remove(item)
		return item.Value.(*http.Response)
	} else {
		return nil
	}
}

func StartParse() {
	for IsRunning {
		resp := GetRespFromQueue()
		if resp != nil {
			log.Println("Start parsing ...")
			go PageParse(resp)
			// IsRunning = false

		}
	}
}
