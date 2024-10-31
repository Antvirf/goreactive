package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/antvirf/goreactive"
)

var pageData PageData

type PageData struct {
	Title                   string
	ReactiveVars            []*goreactive.ReactiveVar
	ReactiveVarsScriptBlock string
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("./index.html")).Execute(w, pageData)
}

func main() {
	// Create some reactive vars
	example_one := goreactive.NewReactiveVar("28")
	example_two := goreactive.NewReactiveVar("12")

	// Initialize our global state of data for example page
	pageData = PageData{
		Title: "Reactive Vars Example",
		ReactiveVars: []*goreactive.ReactiveVar{
			example_one,
			example_two,
		},
		ReactiveVarsScriptBlock: goreactive.WebsocketJavascriptBlock,
	}

	// Start updating variables randomly
	log.Println("Starting random variable updater in the background...")
	go RandomUpdatesToReactiveVars(example_one, example_two)

	// Set up an HTTP server as we normally would
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveIndex)
	mux.HandleFunc("/reactiveVarsWebsocket", goreactive.WebsocketServerHandler)
	handler := http.Handler(mux)
	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
