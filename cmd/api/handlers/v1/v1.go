// Package v1 contains the full set of handler functions and routes supported by the v1 web api.
package v1

import (
	"net/http"
	"web-server/cmd/api/handlers/v1/private"
	"web-server/cmd/api/handlers/v1/public"
	"web-server/pkg/events"
	"web-server/pkg/web"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const version = "v1"

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log    *zap.SugaredLogger
	Events *events.Events
}

// PublicRoutes binds all the version 1 public routes.
func PublicRoutes(app *web.App, cfg Config) {
	pbl := public.Handlers{
		Log:    cfg.Log,
		WS:     websocket.Upgrader{},
		Events: cfg.Events,
	}

	app.Handle(http.MethodGet, version, "/ws", pbl.EventsWebsocket)
	app.Handle(http.MethodGet, version, "/get", pbl.Get)
	app.Handle(http.MethodPost, version, "/post", pbl.Post)
}

// PrivateRoutes binds all the version 1 private routes.
func PrivateRoutes(app *web.App, cfg Config) {
	prv := private.Handlers{
		Log: cfg.Log,
	}

	app.Handle(http.MethodGet, version, "/get", prv.Get)
	app.Handle(http.MethodPost, version, "/post", prv.Post)
}
