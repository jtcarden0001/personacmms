package main

import (
	"github.com/jtcarden0001/personacmms/webapi/internal/app"
	"github.com/jtcarden0001/personacmms/webapi/internal/httpapi"
	"github.com/jtcarden0001/personacmms/webapi/internal/store"
)

func main() {
	storeLayer := store.New()

	appLayer := app.New(storeLayer)

	httpLayer := httpapi.New(appLayer)

	httpLayer.Start()
}
