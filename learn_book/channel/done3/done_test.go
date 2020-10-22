package channel

import (
	"fmt"
	"sync"
	"testing"
)

// 使用 sync.WaitGroup 来同步routine 完成任务
// doWorker封装抽象一下

type Worker struct {
	in chan int
	done func()
}

// 使用 chan 来通知完成工作
func doWorker(id int, worker Worker) {
	for n := range worker.in{
		// n, ok := <-c
		fmt.Printf("Worker %d received %c\n", id, n)
		worker.done()
	}

	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}
}

func CreateWorker(id int, wg *sync.WaitGroup) Worker{
	worker := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}

	go doWorker(id, worker)
	return worker
}


func chanDemo()  {
	var workers [10]Worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers{
		worker.in <- 'a' + i
	}

	for i, worker := range workers{
		worker.in <- 'A' + i
	}
	wg.Wait()
}


func bufferedChannel()  {
	c := make(chan int, 3)
	//go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
}

//func channelClose()  {
//	w := CreateWorker(0)
//	c := w.in
//	c <- 'a'
//	c <- 'b'
//	c <- 'c'
//	c <- 'd'
//	// sender close
//	close(c)
//	time.Sleep(time.Millisecond)
//
//}


func TestDemo(t *testing.T) {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}