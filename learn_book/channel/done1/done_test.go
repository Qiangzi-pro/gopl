package channel

import (
	"fmt"
	"testing"
	"time"
)

type Worker struct {
	in chan int
	done chan bool
}

// 使用 chan 来通知完成工作
func doWorker(id int, c chan int, done chan bool) {
	for n := range c{
		// n, ok := <-c
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() {
			done <- true
		}()
	}

	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}
}

func CreateWorker(id int) Worker{
	worker := Worker{
		in: make(chan int),
		done: make(chan bool),
	}

	go doWorker(id, worker.in, worker.done)
	return worker
}


func chanDemo()  {
	var workers [10]Worker

	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i)
	}

	for i, worker := range workers{
		worker.in <- 'a' + i
	}

	for i, worker := range workers{
		worker.in <- 'A' + i
	}

	for _, worker := range workers{
		<-worker.done
		<-worker.done
	}
}


func bufferedChannel()  {
	c := make(chan int, 3)
	//go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
}

func channelClose()  {
	w := CreateWorker(0)
	c := w.in
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// sender close
	close(c)
	time.Sleep(time.Millisecond)

}


func TestDemo(t *testing.T) {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}