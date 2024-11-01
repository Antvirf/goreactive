package goreactive

import (
	"context"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

func WebsocketServerHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("error for client connecting from %s: %s", r.RemoteAddr, err)
	}
	defer c.CloseNow()

	ctx := context.Background()
	ctx = c.CloseRead(ctx)

	// Add subscription
	sub := updateBroker.subscribe(r.RemoteAddr)

	for {
		select {
		case <-ctx.Done():
			updateBroker.unsubscribe(r.RemoteAddr)
			c.Close(websocket.StatusNormalClosure, "")
			return
		case val := <-sub:
			err = wsjson.Write(ctx, c, val)
			if err != nil {
				log.Printf("error sending to client %s: %s", r.RemoteAddr, err)
				return
			}
		}
	}
}
