package channel

import (
	"fmt"
	"testing"
	"time"
)

func worker(id int, c chan int) {
	for n := range c{
		// n, ok := <-c
		fmt.Printf("Worker %d received %c\n", id, n)
	}

	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}
}

func CreateWorker(id int) chan<- int{
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo()  {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
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
	c := CreateWorker(0)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// sender close
	close(c)

	time.Sleep(time.Millisecond)
}


func TestDemo(t *testing.T) {
	//chanDemo()
	//bufferedChannel()
	channelClose()
	//c := make(chan int)
	//c <- 'a'
}