# crawler
用Golang写的爬虫，以此学习Go语言

# 思路
1. 爬虫分为下载器、解析器和解析配置文件；
2. 下载器负责下载页面，页面获取正常是返回http.Response对象，将该对象放到全局队列，后期考虑用消息总线实现；
3. 解析器利用goquery库进行页面解析，解析规则有配置文件传入

=============================================
将命令行参数os.Args作为参数利用FlagSet进行解析时给FlagSet.Parse()函数传递的参数是os.Args[1:]要忽略第一个参数