package main

import (
	"fmt"
	"sync"
)

type Buffer struct {
	count int
	mu    sync.Mutex
}

var buffer Buffer
var cond = sync.NewCond(&buffer.mu)
var size = 3

func producer(buffer *Buffer) {
	for {
		buffer.mu.Lock()
		for buffer.count == size {
			cond.Wait()
		}
		buffer.count++
		fmt.Print("(")
		buffer.mu.Unlock()
		cond.Broadcast()
	}
}

func consumer(buffer *Buffer) {
	for {
		buffer.mu.Lock()
		for buffer.count == 0 {
			cond.Wait()
		}
		buffer.count--
		fmt.Print(")")
		buffer.mu.Unlock()
		cond.Broadcast()
	}

}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(&buffer)
	go producer(&buffer)
	go consumer(&buffer)
	go consumer(&buffer)
	wg.Wait()
}
