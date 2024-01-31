## GO PUBSUb package

## Installation

```bash
go get github.com/rammyblog/pubsub
```

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/rammyblog/pubsub"
)

type CustomSubscriber struct {
	Name string
}

func (cs *CustomSubscriber) Notify(msg pubsub.Message) {
	fmt.Printf("[%s] Received message on topic %s: %v\n", cs.Name, msg.Topic, msg.Data)
}

func main() {
	ps := pubsub.New()

	// Create subscribers
	subscriber1 := &CustomSubscriber{Name: "Subscriber 1"}
	subscriber2 := &CustomSubscriber{Name: "Subscriber 2"}

	// Subscribe subscribers to topics
	ps.Subscribe("topic1", subscriber1)
	ps.Subscribe("topic2", subscriber2)

	// Publish messages
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Second)
			ps.Publish(pubsub.Message{Topic: "topic1", Data: i})
		}
	}()

	go func() {
		for i := 10; i <= 15; i++ {
			time.Sleep(2 * time.Second)
			ps.Publish(pubsub.Message{Topic: "topic2", Data: i})
		}
	}()

	// Unsubscribe subscriber1 after some time
	time.Sleep(4 * time.Second)
	ps.Unsubscribe("topic1", subscriber1)

	// Keep the program running
	select {}
}

```

## Testing

To run the tests, execute the following command:

```bash
go test ./...
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
