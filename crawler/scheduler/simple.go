package scheduler

import "learngo/gopl.io/crawler/types"

type SimpleScheduler struct {
	in chan types.Request
}

func (s *SimpleScheduler) WorkerChan() chan types.Request {
	return s.in
}

func (s *SimpleScheduler) WorkerReady(w chan types.Request) {
}

func (s *SimpleScheduler) Run() {
	s.in = make(chan types.Request)

}

func (s *SimpleScheduler) Submit(request types.Request) {
	go func() {
		s.in<- request
	}()
}
