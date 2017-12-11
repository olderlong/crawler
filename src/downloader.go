package crawler

//downloader.go 下载页面，获取源代码
import (
	"net/http"
)

func GetPage(url string) (resp *http.Response, statusCode int) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}

	statusCode = resp.StatusCode
	return
}
