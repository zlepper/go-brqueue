package main

import (
	"github.com/zlepper/go-brqueue"
	"log"
	"time"
)

func popOnly(threadCount, messageCount int) {
	before := time.Now()

	client, err := brqueue.NewClient("localhost", 6431)
	if err != nil {
		log.Panicln(err)
	}
	defer client.Close()

	after := time.Now()

	log.Printf("Pop-only Setup took '%s'\n", after.Sub(before).String())

	measure(threadCount, messageCount, "Pop-Ack Only", func() {
		for j := 0; j < messageCount; j++ {
			task, err := client.Pop([]string{}, true)
			if err != nil {
				log.Panicln(err)
			}

			err = client.Acknowledge(task)
			if err != nil {
				log.Panicln(err)
			}
		}
	})
}
