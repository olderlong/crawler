package crawler

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// config.go 用来处理config.json配置文件，解析其中的参数

//ItemRule 每一项页面内容解析规则结构，与JSON文件中字段对应
type ItemRule struct {
	ItemName    string `json:"ItemName"`
	CSSSelector string `json:"CSSSelector"`
	AttrName    string `json:"AttrName,omitempty"`
}

// ParseRule 页面解析规则结构，每个ParseRule下的所有Item位于同一根节点下
type ParseRule struct {
	TopicName           string     `json:"TopicName"`
	RootNodeCSSSelector string     `json:"RootNodeCSSSelector"`
	ItemRules           []ItemRule `json:"ItemRules"`
}

// CrawlerLinkRule 爬取链接解析规则，包括URL模式和链接解析规则（ItemRule）
type CrawlerLinkRule struct {
	URLPattern string   `json:"URLPattern"`
	LinkRule   ItemRule `json:"LinkRule"`
}

// CrawlerConfigure 页面爬取解析配置
type CrawlerConfigure struct {
	ConfigureName    string            `json:"ConfigureName"`
	BaseURL          string            `json:"BaseURL"`
	CrawlerLinkRules []CrawlerLinkRule `json:"CrawlerLinkRules"`
	ParseRules       []ParseRule       `json:"ParseRules"`
}

// CrawlerConfigureParse 从配置文件中解析出页面爬取解析规则
func CrawlerConfigureParse(confJSONFilePath string) (CrawlerConfigure, error) {
	conf := CrawlerConfigure{}
	bytes, errFileRead := ioutil.ReadFile(confJSONFilePath)
	if errFileRead != nil {
		log.Fatal(errFileRead)
		return conf, errFileRead
	}
	errJSONParse := json.Unmarshal(bytes, &conf)
	if errJSONParse != nil {
		log.Fatal(errJSONParse)
		return conf, errFileRead
	}
	return conf, nil
}
