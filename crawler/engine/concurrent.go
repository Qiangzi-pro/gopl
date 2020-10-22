package engine
//
//import (
//	"learngo/gopl.io/crawler/types"
//	"log"
//)
//
//type Concurrent struct {
//	Scheduler   Scheduler
//	WorkerCount int
//}
//
//type Scheduler interface {
//	ConfigurationMasterWorkerChan(chan types.Request)
//	Submit(request types.Request)
//	WorkerReady(w chan types.Request)
//}
//
//func (e Concurrent) Run(seeds ...types.Request) {
//	in := make(chan types.Request)
//	out := make(chan types.ParseResult)
//
//	e.Scheduler.ConfigurationMasterWorkerChan(in)
//
//	for i := 0; i < e.WorkerCount; i++ {
//		CreateWorker(in, out)
//	}
//
//	for _, req := range seeds {
//		e.Scheduler.Submit(req)
//	}
//
//	for {
//		parseResult := <-out
//		for _, item := range parseResult.Items {
//			log.Printf("Got item: %v", item)
//		}
//
//		for _, req := range parseResult.Requests {
//			e.Scheduler.Submit(req)  // 这里容易出现循环等待
//		}
//	}
//}
//
//func CreateWorker(in <-chan types.Request, out chan<- types.ParseResult) {
//	go func() {
//		for {
//			req := <-in
//			parseResult, err := worker(req)
//			if err != nil {
//				continue
//			}
//			out <-parseResult
//		}
//	}()
//}
