package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// config.go 用来处理config.json配置文件，解析其中的参数

//Rule 解析规则结构，与JSON文件中字段对应
type Rule struct {
	ItemName    string `json:"ItemName"`
	CSSSelector string `json:"CSSSelector"`
	AttrName    string `json:"AttrName,omitempty"`
}

//ParseConfig 解析配置结构，与JSON文件中字段对应
type ParseConfig struct {
	URLPattern string `json:"URLPattern"`
	Rules      []Rule `json:"Rules"`
}

// GetParseConfigs 读取json格式的配置文件，从中解析出页面解析配置列表
func GetParseConfigs(configJSONFilePath string) []ParseConfig {
	bytes, err := ioutil.ReadFile(configJSONFilePath)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	var configs []ParseConfig
	err = json.Unmarshal(bytes, &configs)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	return configs
}
