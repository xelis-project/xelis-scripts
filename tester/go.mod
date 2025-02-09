module tester

go 1.21.6

require (
	github.com/google/uuid v1.6.0
	github.com/xelis-project/xelis-go-sdk v0.4.4
)

replace github.com/xelis-project/xelis-go-sdk => ../../xelis-go-sdk

require (
	github.com/creachadair/jrpc2 v0.44.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)
