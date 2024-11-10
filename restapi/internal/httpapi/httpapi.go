package httpapi

import (
	app "github.com/jtcarden0001/personacmms/restapi/internal/app"
	imp "github.com/jtcarden0001/personacmms/restapi/internal/httpapi/v1/gin"
)

// HttpApi layer hosts the http routing, handling, and serving logic and forwards requests to the App layer.
type HttpApi interface {
	Start()
}

func New(injectedApp app.App) HttpApi {
	return imp.New(injectedApp)
}
