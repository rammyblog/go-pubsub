package pubsub

import (
	"sync"
)

type Message struct {
	Topic string
	Data  interface{}
}

type ISubscriber interface {
	Notify(msg Message)
}

type PubSub struct {
	mu          sync.Mutex
	Subscribers map[string][]ISubscriber
}

func New() *PubSub {
	return &PubSub{
		Subscribers: make(map[string][]ISubscriber),
	}
}

func (ps *PubSub) Subscribe(topic string, subscriber ISubscriber) ISubscriber {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.Subscribers[topic] = append(ps.Subscribers[topic], subscriber)
	return subscriber
}

func (ps *PubSub) Unsubscribe(topic string, subscriber ISubscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	subscribers := ps.Subscribers[topic]
	for i, sub := range subscribers {
		if sub == subscriber {
			ps.Subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
			return
		}
	}

}

func (ps *PubSub) Publish(msg Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	subscribers := ps.Subscribers[msg.Topic]
	for _, sub := range subscribers {
		go func(s ISubscriber) {
			s.Notify(msg)
		}(sub)
	}

}
