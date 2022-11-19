// Package private maintains the group of handlers for private access.
package private

import (
	"context"
	"net/http"
	"web-server/pkg/web"

	"go.uber.org/zap"
)

// Handlers manages the set of bar ledger endpoints.
type Handlers struct {
	Log    *zap.SugaredLogger
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
