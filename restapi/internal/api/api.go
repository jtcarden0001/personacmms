package api

import (
	imp "github.com/jtcarden0001/personacmms/restapi/internal/api/v1/gin"
	app "github.com/jtcarden0001/personacmms/restapi/internal/app"
)

// Api layer hosts the http routing, handling, and serving logic and forwards requests to the App layer.
type Api interface {
	Start()
}

func New(injectedApp app.App) Api {
	return imp.New(injectedApp)
}
