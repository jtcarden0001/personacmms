package main

import (
	api "github.com/jtcarden0001/personacmms/restapi/internal/api"
	application "github.com/jtcarden0001/personacmms/restapi/internal/app"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
)

func main() {
	store := st.New()
	app := application.New(store)
	http := api.New(app)
	http.Start()
}
