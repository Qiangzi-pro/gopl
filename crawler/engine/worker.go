package engine

import (
	"learngo/gopl.io/crawler/fetcher"
	"learngo/gopl.io/crawler/types"
	"log"
)

func worker(req types.Request) (types.ParseResult, error) {
	content, err := fetcher.Fetcher(req.Url, req.Gatherer, nil)
	if err != nil {
		log.Printf("fetcher: error url %s: %v", req.Url, err)
		return types.ParseResult{}, err
	}

	parseResult := req.ParseFunc(content)
	return parseResult, nil
}