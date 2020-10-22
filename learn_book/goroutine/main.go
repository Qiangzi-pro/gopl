package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func(i int) { // race condition
			for  {
				fmt.Printf("hello goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//fmt.Println(a)
}

