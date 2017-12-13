package crawler

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
)

type Crawler struct {
	StartURL string
}

func (this *Crawler) Run(args []string) {
	fmt.Printf("%v", args)
	fSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	var ParseConfigFilePath = fSet.String("config", "", "页面解析配置文件完整路径")
	fSet.Parse(args[1:])

	ParseConfigs = GetParseConfigs(*ParseConfigFilePath)
	URLQueue.PushBack(this.StartURL)

	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt, os.Kill)

	IsRunning = true
	go StartDownload()
	go StartParse()

	s := <-c
	IsRunning = false
	fmt.Println("Got signal:", s) //Got signal: terminated
}
