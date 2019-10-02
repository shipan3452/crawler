package engine

import "awesomeProject/crawler/config"


type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

/**
*一个页面
 */
type Request struct {
	Url        string
	Parser     Parser
}

/**
* 页面解析
 */
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type ParserFunc func(
	contents []byte, url string) ParseResult


type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(
	contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (
	name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}


type NilParser struct{}

func (NilParser) Parse(
	_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (
	name string, args interface{}) {
	return config.NilParser, nil
}


/**
*调度器
 */
type Scheduler interface {
	Submit(Request)
	SetChannel(c chan Request)
	WorkerIsReady(c chan Request)
}
