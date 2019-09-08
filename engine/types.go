package engine

/**
*一个页面
 */
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

/**
* 页面解析
 */
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

/**
*调度器
 */
type Scheduler interface {
	Submit(Request)
	SetChannel(c chan Request)
	WorkerIsReady(c chan Request)
}
