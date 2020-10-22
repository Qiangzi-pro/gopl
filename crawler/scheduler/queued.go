package scheduler

import "learngo/gopl.io/crawler/types"

type QueuedScheduler struct {
	requestChan chan types.Request
	workerChan chan chan types.Request
}

func (s *QueuedScheduler) WorkerChan() chan types.Request {
	return make(chan types.Request)
}

func (s *QueuedScheduler) Submit(request types.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) WorkerReady(w chan types.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan types.Request)
	s.workerChan = make(chan chan types.Request)

	go func() {
		var requestQ []types.Request
		var workerQ []chan types.Request
		for {
			var activeRequest types.Request
			var activeWorker chan types.Request
			if len(requestQ) > 0 &&
				len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case req := <-s.requestChan:
				requestQ = append(requestQ, req)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}



