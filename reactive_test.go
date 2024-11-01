package goreactive

import (
	"sync"
	"testing"
	"time"
)

func TestReactiveVarSendsChangeMessage(t *testing.T) {
	myvar := NewReactiveVar("initial value")

	if myvar.Value != "initial value" {
		t.Fatalf(`Initialized ReactiveVar = %q, but has value %q`, "initial value", myvar.Value)
	}

	// Listen to changes
	var wg sync.WaitGroup
	var payload channelPayload
	wg.Add(1)
	go func(payload *channelPayload) {
		defer wg.Done()

		sub := updateBroker.subscribe("test")
		*payload = <-sub
		updateBroker.unsubscribe("test")
	}(&payload)
	<-time.After(time.Millisecond * 10) // Allow subscriber some time to hook in

	// Change the var
	myvar.Update("changed")

	// Wait for update listener
	wg.Wait()
	if payload.Value != "changed" {
		t.Fatalf(`channelPayload for variable incorrect, was %q, expected %q`, payload.Value, "changed")
	}
}
