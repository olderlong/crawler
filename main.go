package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/olderlong/crawler/src"
)

func main() {

	url := "http://blog.csdn.net/"
	url1 := "http://www.mamicode.com"
	fmt.Println(url)
	crawler.URLQueue.PushBack(url)
	crawler.URLQueue.PushBack(url1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	crawler.IsRunning = true
	go crawler.StartDownload()
	time.Sleep(3)
	go crawler.StartParse()

	s := <-c
	fmt.Println("Got signal:", s) //Got signal: terminated

	crawler.IsRunning = false
	// // code := crawler.GetPage(crawler.GetURLFromQueue())
	// // println(code)

	// resp := crawler.GetRespFromQueue()

	// items := crawler.PageParse(resp)
	// // items := crawler.LinkParse(url)

	// for _, item := range items {
	// 	// fmt.Printf("URL:\t %s \t\t Link Text:\t %s\n", item.URL, item.Text)
	// 	crawler.ResponseQueue.PushBack(item)
	// }

	// fmt.Println(url + "中共有" + strconv.Itoa(len(items)) + "链接")
	// for crawler.ResponseQueue.Len() > 0 {
	// 	it := crawler.ResponseQueue.Front()
	// 	crawler.ResponseQueue.Remove(it)
	// 	item, _ := it.Value.(crawler.LinkItem)

	// 	fmt.Printf("URL:\t %s \t\t Link Text:\t %s\n", item.URL, item.Text)
	// }
}
