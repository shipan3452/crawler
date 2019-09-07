package parser

import (
	"awesomeProject/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+")[^>]*?>([^<]+)</a>`

func ParseCityList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc:engine.NilParser})
	}
	return result
}