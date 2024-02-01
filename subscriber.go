package pubsub

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Subscriber struct {
	id       string
	topics   []string
	messages chan *Message
	mu       sync.Mutex
	active   bool
}

func generateUniqueID() string {
	id := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000000)
	return strconv.Itoa(id)
}

func NewSubscriber() *Subscriber {
	id := generateUniqueID()

	return &Subscriber{
		id:       id,
		topics:   make([]string, 0),
		messages: make(chan *Message),
		active:   true,
	}
}

func (s *Subscriber) GetID() string {
	return s.id
}

func (s *Subscriber) GetTopics() []string {
	return s.topics
}

func (s *Subscriber) AddTopic(topic string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.topics = append(s.topics, topic)
}

func (s *Subscriber) RemoveTopic(topic string) {
	// maybe a map would be better here
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, t := range s.topics {
		if t == topic {
			s.topics = append(s.topics[:i], s.topics[i+1:]...)
		}
	}
}

func (s *Subscriber) Destroy() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.active = false
	close(s.messages)
}

func (s *Subscriber) Notify(msg *Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.active {
		s.messages <- msg
	}
}

func (s *Subscriber) GetMessages() chan *Message {
	return s.messages
}

func (s *Subscriber) PrintMessages() {
	for {
		select {
		case msg := <-s.messages:
			fmt.Println("Subscriber: " + s.id + " received message: " + msg.Data.(string))
		default:
			// No message to receive, continue the loop
			continue
		}
	}
}
