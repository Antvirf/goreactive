// `goreactive` allows you to build web applications such as dashboards that update in real-time without writing any JavaScript. Use standard Go templates, embed them with `ReactiveVar` objects and the library takes care of the rest. Updates are pushed to clients over Websockets.

package goreactive

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

var updateBroker *messageBroker

// Initializes the message broker used to fan out updates to all subscribed clients
func init() {
	updateBroker = newBroker()
}

// ReactiveVar stores string data in a concurrency-safe way.
type ReactiveVar struct {
	identifier string
	value      string
	mu         sync.Mutex
}

func NewReactiveVar(value string) *ReactiveVar {
	return &ReactiveVar{
		identifier: uuid.NewString(),
		value:      value,
	}
}

// Update provides a concurrency-safe way to update the value owned by a ReactiveVar, and triggers a change notification to all subscribed clients.
func (r *ReactiveVar) Update(newValue string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.value = newValue
	updateBroker.publish(
		channelPayload{
			Key:   r.identifier,
			Value: r.value, // Read new val from object itself
		})
}

// String formats the ReactiveVar as a HTML span for direct usage in HTML templates.
func (r *ReactiveVar) String() string {
	return fmt.Sprintf(`<span id="%s">%s</span>`, r.identifier, r.value)
}
