package main

import (
	"learngo/gopl.io/crawler/engine"
	"learngo/gopl.io/crawler/persist"
	"learngo/gopl.io/crawler/scheduler"
	"learngo/gopl.io/crawler/types"
	gather2 "learngo/gopl.io/crawler/zhenai/gather"
	"learngo/gopl.io/crawler/zhenai/parser"
)

func main() {
	e := engine.Concurrent{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 5,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun/",
		ParseFunc: parser.ParseCityList,
		Gatherer:  gather2.CityListGather{},
	})
}
