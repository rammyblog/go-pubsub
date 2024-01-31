package main

import (
	"fmt"

	"github.com/rammyblog/pubsub"
)

type Subscriber struct {
}

func Sub(ps *pubsub.PubSub) []pubsub.ISubscriber {
	subscriber1 := ps.Subscribe("topic1", &Subscriber{})
	subscriber2 := ps.Subscribe("topic2", &Subscriber{})
	subscriber3 := ps.Subscribe("topic3", &Subscriber{})

	return []pubsub.ISubscriber{subscriber1, subscriber2, subscriber3}
}

func (s *Subscriber) Notify(msg pubsub.Message) {
	// do something with the message
	fmt.Println("Subscriber notified: ", msg.Topic, msg.Data)
}
