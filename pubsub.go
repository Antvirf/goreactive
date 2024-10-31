package goreactive

import (
	"log"
	"sync"
)

// channelPayloads are the structs serialized and sent to each open websocket/client listener directly.
type channelPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// messageBroker is the internal object responsible for sending out a channelPayload to all subscribers.
type messageBroker struct {
	quit        chan struct{}
	subscribers map[string]chan channelPayload
	closed      bool
	mu          sync.Mutex
}

// newBroker returns an initialized messageBroker, ready for use.
func newBroker() *messageBroker {
	return &messageBroker{
		mu:          sync.Mutex{},
		quit:        make(chan struct{}),
		subscribers: make(map[string]chan channelPayload, 0),
		closed:      false,
	}
}

// publish sends a channelPayload to every subscriber over their respective channel.
func (b *messageBroker) publish(payload channelPayload) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	for _, channel := range b.subscribers {
		// Non-blocking send to all subscribers
		select {
		case channel <- payload:
		default:
			continue // skip if receiver not ready or any other problem.
		}
	}
}

// subscribe adds a listener with specified ID to the broker.
func (b *messageBroker) subscribe(listenerId string) <-chan channelPayload {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return nil
	}

	ch := make(chan channelPayload)
	b.subscribers[listenerId] = ch
	return ch
}

// unsubscribe removes a listener with specified ID to the broker.
func (b *messageBroker) unsubscribe(listenerId string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}
	delete(b.subscribers, listenerId)
	log.Printf("unsubbed: %s", listenerId)
}

// closes the messageBroker, and all its active channels.
func (b *messageBroker) close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}
	b.closed = true
	close(b.quit)
	for _, channel := range b.subscribers {
		close(channel)
	}
}
