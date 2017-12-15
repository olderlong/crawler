package crawler

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parser.go 解析页面内容

// Parser 页面解析类
type Parser struct {
	//是否运行标志，用于解析线程控制
	IsRunning bool
	// 解析配置信息
	ParserConfig CrawlerConfigure
	// 解析对应的网站URL，主要用于不全相对当前网站的相对链接
	SiteURL        string
	CurrentPageURL string
	BaseURL        string
	// 待爬取链接解析参数
	CrawlerLinkRules []CrawlerLinkRule
	// 爬取内容的解析参数
	ItemParseRules []ParseRule
}

// LinkParse 根据Url获取网页中当前域名下的URL链接
func (p *Parser) PageParse(resp *http.Response) {

	doc, err := goquery.NewDocumentFromResponse(resp) //这里已经将resp关闭了，对resp的处理要放在前面
	if err != nil {
		log.Fatal(err)
		return
	}
	// u ,_:=resp.Location()	//location函数从resp的Header中读取Location字段的值，有时会为空，这里不采用
	u := resp.Request.URL
	p.SiteURL = u.Scheme + "://" + u.Host
	p.CurrentPageURL = u.String()

	go p.ParseLinks(*doc)

	p.ParseItems(*doc)
}

//ParseLinks 从页面中解析出需要爬取的链接
func (p *Parser) ParseLinks(doc goquery.Document) {
	var pattern = ""
	for _, crawlerLinkRule := range p.CrawlerLinkRules {
		pattern = crawlerLinkRule.URLPattern
		linkRule := crawlerLinkRule.LinkRule
		doc.Find(linkRule.CSSSelector).Each(func(_ int, s *goquery.Selection) {
			link, isExist := s.Attr(linkRule.AttrName)
			if isExist {
				// 补全URL
				if !strings.HasPrefix(strings.ToLower(link), "http") {
					if !strings.HasPrefix(strings.ToLower(link), "/") {
						link = p.SiteURL + "/" + link
					} else {
						link = p.SiteURL + link
					}
				}
				// 检查URL是否符合爬取规则
				if p.CheckURLPatternMatch(link, pattern) && p.IsInURLQueue(link) {
					URLQueue.PushBack(link)
				}
			}
		})
	}
}

//CheckURLPatternMatch 检查URL与爬取链接模式是否匹配
func (p *Parser) CheckURLPatternMatch(url string, pattern string) bool {
	return true
}

//IsInURLQueue 检查URL是否已经在爬取队列里
func (p *Parser) IsInURLQueue(url string) bool {
	return false
}

//ParseItems 页面内容解析检查URL是否已经在爬取队列里
func (p *Parser) ParseItems(doc goquery.Document) PageParseResult {
	var pageRes PageParseResult
	for _, itemParseRules := range p.ItemParseRules {
		//一个页面可以解析多个主题，每个主题在一个HTMLNode节点下
		var topicRes TopicParseResult
		doc.Find(itemParseRules.RootNodeCSSSelector).Each(func(_ int, sNode *goquery.Selection) {
			for _, itemRule := range itemParseRules.ItemRules {
				sNode.Find(itemRule.CSSSelector).Each(func(_ int, sItem *goquery.Selection) {
					var itemValue string
					var isExist = false
					if itemRule.AttrName == "Text" {
						itemValue = sNode.Text()
					} else {
						itemValue, isExist = sNode.Attr(itemRule.AttrName)
						if !isExist {
							itemValue = ""
						}
					}
					topicRes[itemRule.ItemName] = itemValue
				})
			}
		})
		pageRes[itemParseRules.TopicName] = topicRes
	}
	log.Printf("%v\n", pageRes)
	return pageRes
}

//GetRespFromQueue 从响应队列中获取响应
func (p *Parser) GetRespFromQueue() *http.Response {
	if ResponseQueue.Len() > 0 {
		item := ResponseQueue.Front()
		ResponseQueue.Remove(item)
		return item.Value.(*http.Response)
	} else {
		return nil
	}
}

//InitConfigure 初始化配置
func (p *Parser) InitConfigure() {
	p.ItemParseRules = p.ParserConfig.ParseRules
	p.CrawlerLinkRules = p.ParserConfig.CrawlerLinkRules
	p.BaseURL = p.ParserConfig.BaseURL
}

// Start 启动页面解析线程
func (p *Parser) Start() {
	p.IsRunning = true
	p.InitConfigure()
	for p.IsRunning {
		resp := p.GetRespFromQueue()
		if resp != nil {
			log.Println("Start parsing, base url is", p.BaseURL)
			go p.PageParse(resp)
		}
	}
}
