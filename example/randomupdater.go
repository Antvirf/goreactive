package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/antvirf/goreactive"
)

func RandomUpdatesToReactiveVars(vars ...*goreactive.ReactiveVar) {
	var wg sync.WaitGroup

	for _, val := range vars {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				randValue := rand.IntN(2000)
				<-time.After(time.Duration(randValue) * time.Millisecond)
				val.Update(fmt.Sprintf("%d", randValue))
			}
		}()

	}
	wg.Wait()
}
