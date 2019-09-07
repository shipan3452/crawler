package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 2,
	}
	e.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
