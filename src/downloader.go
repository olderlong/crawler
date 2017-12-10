package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//LinkItem 链接对象
type LinkItem struct {
	//url 网页内的URL链接
	URL string
	//text URL链接的文字
	Text string
}

func GetPageSource(url string) (content string, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)
	return
}

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
