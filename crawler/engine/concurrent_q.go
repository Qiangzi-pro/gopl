package engine

import (
	"learngo/gopl.io/crawler/types"
)

type Concurrent struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	// 问sche要chan
	WorkerChan() chan types.Request
	Submit(request types.Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(w chan types.Request)
}

func (e Concurrent) Run(seeds ...types.Request) {
	out := make(chan types.ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		CreateWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, req := range seeds {
		e.Scheduler.Submit(req)
	}

	for {
		parseResult := <-out
		for _, item := range parseResult.Items {
			// 该项目 ItemSaver的速度比Fetcher快
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, req := range parseResult.Requests {
			// 布隆过滤器？
			if isDuplicate(req.Url) {
				continue
			}

			e.Scheduler.Submit(req)
		}
	}
}

func CreateWorker(in chan types.Request, out chan<- types.ParseResult, s ReadyNotifier) {
	go func() {
		for {
			s.WorkerReady(in)
			req := <-in
			parseResult, err := worker(req)
			if err != nil {
				continue
			}
			out <-parseResult
		}
	}()
}


var urlBloomFilter int

func isDuplicate(url string) bool {
	return false
}