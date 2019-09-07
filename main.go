package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/zhenai/parser"
)
func main() {
	engine.Run(engine.Request{
		Url:"https://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}

