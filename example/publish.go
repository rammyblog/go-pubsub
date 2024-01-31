package main

import (
	"time"

	"github.com/rammyblog/pubsub"
)

type Publisher struct {
}

var TOPICS = []string{"topic1", "topic2", "topic3"}

func Pub(ps *pubsub.PubSub, subscribers []pubsub.ISubscriber) {
	go func() {
		for i := 10; i <= 15; i++ {
			time.Sleep(2 * time.Second)
			ps.Publish(pubsub.Message{
				Topic: "topic1",
				Data:  "hello",
			})
		}
	}()

	go func() {
		for i := 10; i <= 15; i++ {
			time.Sleep(2 * time.Second)
			ps.Publish(pubsub.Message{
				Topic: "topic2",
				Data:  "yooo",
			})
		}
	}()

	go func() {
		for i := 10; i <= 15; i++ {
			time.Sleep(2 * time.Second)
			ps.Publish(pubsub.Message{
				Topic: "topic3",
				Data:  "hello there",
			})
		}
	}()

	time.Sleep(10 * time.Second)
	for _, t := range TOPICS {
		for _, s := range subscribers {
			ps.Unsubscribe(t, s)

		}
	}
}
