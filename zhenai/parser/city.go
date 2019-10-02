package parser

import (
	"awesomeProject/crawler/config"
	"awesomeProject/crawler/engine"
	"log"
	"regexp"
)

var (
	cityRe    = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*?>([^<]+)</a>`
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(content []byte,_ string) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), Parser: NewProfileParser(string(m[2]))})
	}

	//下一页
	matches = cityUrlRe.FindAllSubmatch(
		content, -1)
	for _, m := range matches {
		log.Printf(string(m[1]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, config.ParseCity),
			})
	}
	return result
}
