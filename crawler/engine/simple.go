package engine

import (
	"learngo/gopl.io/crawler/fetcher"
	"learngo/gopl.io/crawler/types"
	"log"
)

type SimpleEngine struct {}

func (SimpleEngine) Run(seeds ...types.Request) {
	requests := make([]types.Request, 0)
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		log.Printf("fetching url: %s", req.Url)

		parseResult, err := worker(req)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item: %v", item)
		}
	}
}