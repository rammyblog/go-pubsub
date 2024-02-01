package main

import (
	"github.com/rammyblog/pubsub"
)

func Sub(ps *pubsub.PubSub) []pubsub.ISubscriber {

	subscriber1 := pubsub.NewSubscriber()
	subscriber2 := pubsub.NewSubscriber()
	subscriber3 := pubsub.NewSubscriber()

	subscriber1.AddTopic("topic1")
	subscriber2.AddTopic("topic2")
	subscriber3.AddTopic("topic3")

	ps.Subscribe("topic1", subscriber1)
	ps.Subscribe("topic2", subscriber2)
	ps.Subscribe("topic3", subscriber3)

	return []pubsub.ISubscriber{subscriber1, subscriber2, subscriber3}
}
