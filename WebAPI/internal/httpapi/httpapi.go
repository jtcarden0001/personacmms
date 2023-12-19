package httpapi

import (
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/app"
	imp "github.com/jtcarden0001/personacmms/projects/webapi/internal/httpapi/gin"
)

type HttpApi interface {
	Start()
}

func New(injectedApp app.App) *imp.HttpApi {
	return imp.New(injectedApp)
}
