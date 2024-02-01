package main

import (
	"github.com/rammyblog/pubsub"
)

func main() {

	ps := pubsub.New()
	subscribers := Sub(ps)
	Pub(ps, subscribers)

	for _, sub := range subscribers {
		go sub.PrintMessages()
	}
	// Keep the program running
	select {}
}
