package crawler

import "container/list"

// 公共类型的定义

//LinkItem 链接对象
type LinkItem struct {
	//url 网页内的URL链接
	URL string
	//text URL链接的文字
	Text string
}

// 下面是用到的一些全局队列变量

var ResponseQueue = list.New()

var URLQueue = list.New()

var IsRunning = false
var ParseConfigs []ParseConfig
