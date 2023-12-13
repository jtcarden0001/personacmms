package main

import (
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/app"
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/db"
	"github.com/jtcarden0001/personacmms/projects/webapi/internal/http"
)

func main() {
	storeLayer := db.New()

	appLayer := app.New(storeLayer)

	httpLayer := http.New(appLayer)

	httpLayer.Start()
}
