# goreactive - reactive template variables with WebSockets

[![Go Reference](https://pkg.go.dev/badge/github.com/antvirf/goreactive.svg)](https://pkg.go.dev/github.com/antvirf/goreactive)

`goreactive` allows you to build web applications such as dashboards that update in real-time without writing any JavaScript. Use standard Go templates, embed them with `ReactiveVar` objects and the library takes care of the rest. Updates are pushed to clients over Websockets.


## Install

```bash
go get github.com/antvirf/goreactive
```

## Example

```go
package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"text/template"
	"time"

	"github.com/antvirf/goreactive"
)

var pageData PageData

type PageData struct {
	ReactiveVarsScriptBlock string
	ReactiveVars            []*goreactive.ReactiveVar
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	template.Must(template.New("index").Parse(`
  {{ .ReactiveVarsScriptBlock }}
  <h1>Goreactive example</h1>
  <ul>
  {{ range .ReactiveVars}}
    <li>
      {{ . }}
    </li>
  {{ end  }}
  </ul>
    `)).Execute(w, pageData)
}

func main() {
	// Create a reactive var
	example_one := goreactive.NewReactiveVar("28")

	// Set up data to pass to our templates
	pageData = PageData{
		ReactiveVars: []*goreactive.ReactiveVar{example_one},
		ReactiveVarsScriptBlock: goreactive.WebsocketJavascriptBlock,
	}

	// Start updating our variable in the background
	go func() {
		for {
			<-time.After(time.Duration(rand.IntN(1000)) * time.Millisecond)
			example_one.Update(fmt.Sprintf("%d", rand.IntN(1000)))
		}
	}()

	// Set up an HTTP server as we normally would
	// Add handler for template + websocket
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveIndex)
	mux.HandleFunc("/reactiveVarsWebsocket", goreactive.WebsocketServerHandler)

	handler := http.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
```

## Major missing features

- Server-side
  - Authentication
  - Review 'nicer'/proper ways to handle closing the messageBroker channel etc. when the application is stopped
  - Some way to customise the printing format of variables, or allow further customisation of what is sent to the client
  - Tests + load tests (how many websocket subscriptions can we handle?)
- Client-side
  - Inform client on error/disconnect
  - Automatic attempts to redirect to client

