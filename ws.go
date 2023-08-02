package poego

import (
	"context"
	"fmt"

	"github.com/fasthttp/websocket" // indirect
)

func (c *Client) websocket(ctx context.Context) {
	conn, _, err := websocket.DefaultDialer.Dial("", c.conf.Headers)
	if err != nil {
		c.wsErr = err
		return
	}

	defer conn.Close()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := conn.ReadMessage()
			fmt.Println("message:", message)
			if err != nil {
				c.wsErr = err
				return
			}
		}
	}
}

func (c *Client) websocketURL() string { return "" }
