package crawler

//downloader.go 下载页面，获取源代码
import (
	"log"
	"net/http"
	"time"
)

func GetPage(url string) {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err1 := client.Get(url)

	if err1 != nil {
		log.Fatalf("%s 获取错误：%s", url, err1)
		return
	}
	ResponseQueue.PushBack(resp)
}
func GetURLFromQueue() string {
	if URLQueue.Len() > 0 {
		item := URLQueue.Front()
		URLQueue.Remove(item)
		return item.Value.(string)
	} else {
		return ""
	}
}
func StartDownload() {
	for IsRunning {
		url := GetURLFromQueue()
		if url != "" {
			log.Println("Start download ...")
			go GetPage(url)
			// IsRunning = false
		}
	}
}
