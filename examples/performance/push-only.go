package main

import (
	"github.com/zlepper/go-brqueue"
	"log"
	"time"
)

func pushOnly(threadCount, messageCount int) {
	before := time.Now()

	client, err := brqueue.NewClient("localhost", 6431)
	if err != nil {
		log.Panicln(err)
	}
	defer client.Close()

	after := time.Now()

	log.Printf("Push-only Setup took '%s'\n", after.Sub(before).String())

	measure(threadCount, messageCount, "Push Only", func() {
		for j := 0; j < messageCount; j++ {
			_, err := client.EnqueueRequest([]byte("Hello world1!"), brqueue.HighPriority, []string{})
			if err != nil {
				log.Panicln(err)
			}
		}
	})
}