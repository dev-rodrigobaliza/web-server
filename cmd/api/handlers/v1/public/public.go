// Package public maintains the group of handlers for public access.
package public

import (
	"context"
	"net/http"
	"time"
	"web-server/pkg/events"
	"web-server/pkg/web"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Handlers manages the set of bar ledger endpoints.
type Handlers struct {
	Log    *zap.SugaredLogger
	WS     websocket.Upgrader
	Events *events.Events
}

// Events handles a web socket to provide events to a client.
func (h Handlers) EventsWebsocket(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("web value missing from context")
	}

	// Need this to handle CORS on the websocket.
	h.WS.CheckOrigin = func(r *http.Request) bool { return true }

	// This upgrades the HTTP connection to a websocket connection.
	c, err := h.WS.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer c.Close()

	// This provides a channel for receiving events from the blockchain.
	ch := h.Events.Acquire(v.TraceID)
	defer h.Events.Release(v.TraceID)

	// Starting a ticker to send a ping message over the websocket.
	ticker := time.NewTicker(time.Second)

	// Block waiting for events from the blockchain or ticker.
	for {
		select {
		case msg, wd := <-ch:

			// If the channel is closed, release the websocket.
			if !wd {
				return nil
			}

			if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return err
			}

		case <-ticker.C:
			if err := c.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				return nil
			}
		}
	}
}

// Get example.
func (h Handlers) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Do some work and make the appropriate response.

	resp := struct {
		Status string `json:"status"`
	}{
		Status: "get response",
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// Post example.
func (h Handlers) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("web value missing from context")
	}

	// Decode the JSON in the post call into a model (depends on what you want to process here, some request).

	h.Log.Infow("<put method name here>", "traceid", v.TraceID, "some label", "<put here some data received>", "raw request", v)

	// Do some work and make the appropriate response.

	resp := struct {
		Status string `json:"status"`
	}{
		Status: "post response",
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}
