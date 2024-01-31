package main

import (
	"github.com/rammyblog/pubsub"
)

func main() {

	ps := pubsub.New()
	subscribers := Sub(ps)
	Pub(ps, subscribers)

	// Keep the program running
	select {}
}
