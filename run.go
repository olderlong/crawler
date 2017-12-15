package crawler

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
)

type Crawler struct {
	StartURL string
}

func (this *Crawler) Run(args []string) {
	fmt.Printf("%v", args)
	fSet := flag.NewFlagSet(args[0], flag.ExitOnError)
	var ParseConfigFilePath = fSet.String("conf", "", "页面解析配置文件完整路径")
	fSet.Parse(args[1:])

	conf, err := CrawlerConfigureParse(*ParseConfigFilePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	URLQueue.PushBack(this.StartURL)

	var parser = Parser{ParserConfig: conf, BaseURL: this.StartURL}
	c := make(chan os.Signal, 0)
	signal.Notify(c, os.Interrupt, os.Kill)

	IsRunning = true
	go StartDownload()
	go parser.Start()

	s := <-c
	IsRunning = false
	fmt.Println("Got signal:", s) //Got signal: terminated
}
