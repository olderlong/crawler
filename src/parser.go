package crawler

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parser.go 解析页面内容

// LinkParse 根据Url获取网页中当前域名下的URL链接
func PageParse(resp *http.Response) []LinkItem {

	doc, err := goquery.NewDocumentFromResponse(resp) //这里已经将resp关闭了，对resp的处理要放在前面
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// u ,_:=resp.Location()	//location函数从resp的Header中读取Location字段的值，有时会为空，这里不采用
	u := resp.Request.URL
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
