package main

import (
	"log"
	"sync"
	"time"
)

func measure(threadCount, messageCount int, name string, cb func()) {

	before := time.Now()

	var wg sync.WaitGroup

	wg.Add(threadCount)

	for j := 0; j < threadCount; j++ {
		go func() {
			cb()

			wg.Done()
		}()
	}

	wg.Wait()

	after := time.Now()

	log.Printf("%s: Time taken for %d message: %s\n", name, messageCount * threadCount, after.Sub(before).String())
}


func main() {
	threadCount, messageCount := 100, 10000
	pushOnly(threadCount, messageCount)
	popOnly(threadCount, messageCount)
	pushPop(threadCount, messageCount)
}
