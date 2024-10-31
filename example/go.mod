module github.com/antvirf/goreactive/example

replace github.com/antvirf/goreactive => ./../

go 1.22.4

require (
	github.com/antvirf/goreactive v0.0.0-00010101000000-000000000000
	github.com/coder/websocket v1.8.12
)

require github.com/google/uuid v1.6.0 // indirect
