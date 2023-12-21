package httpapi

import (
	"github.com/jtcarden0001/personacmms/webapi/internal/app"
	imp "github.com/jtcarden0001/personacmms/webapi/internal/httpapi/gin"
)

type HttpApi interface {
	Start()
}

func New(injectedApp app.App) HttpApi {
	return imp.New(injectedApp)
}
