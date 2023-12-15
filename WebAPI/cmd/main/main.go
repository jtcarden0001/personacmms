package main

import (
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/app"
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/httpapi"
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/store"
)

func main() {
	storeLayer := store.New()

	appLayer := app.New(storeLayer)

	httpLayer := httpapi.New(appLayer)

	httpLayer.Start()
}
