package goreactive

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMessageBrokerPublish(t *testing.T) {
	updateBroker := newBroker()
	numSubscribers := 3

	// Set up subs
	var wg sync.WaitGroup
	for i := range numSubscribers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sub := updateBroker.subscribe(fmt.Sprintf("%d", i))
			<-sub
			updateBroker.unsubscribe(fmt.Sprintf("%d", i))
		}()
	}

	<-time.After(20 * time.Millisecond) // Allow time for subs to come online

	// Assert on number of subs
	if len(updateBroker.subscribers) != numSubscribers {
		t.Fatalf(`Broker has %d subscribers; expected %d`, len(updateBroker.subscribers), numSubscribers)
	}

	// Send a message
	go updateBroker.publish(channelPayload{
		Key:   "testvalue",
		Value: "testvalue",
	})

	// Wait for all goroutines to finish, after which they unsubscribe
	wg.Wait()

	// Assert on num subs now being zero
	if len(updateBroker.subscribers) != 0 {
		t.Fatalf(`Broker has %d subscribers; expected %d`, len(updateBroker.subscribers), 0)
	}
}

func TestMessageBrokerPublishDoesNotBlock(t *testing.T) {
	updateBroker.publish(channelPayload{
		Key:   "testvalue",
		Value: "testvalue",
	})
	// The fact that this test case exits confirms we're fine :)
}
