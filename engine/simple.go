package engine

import (
	"awesomeProject/crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult,error) {
	content, err := fetcher.Fetcher(r.Url)
	if err != nil {
		log.Printf("fetching:error ,url:%s,%v",r.Url,err)
		return ParseResult{},err
	}
	return  r.Parser.Parse(content,r.Url),nil
}

