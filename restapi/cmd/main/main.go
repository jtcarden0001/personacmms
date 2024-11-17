package main

import (
	"os"

	api "github.com/jtcarden0001/personacmms/restapi/internal/api"
	application "github.com/jtcarden0001/personacmms/restapi/internal/app"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
)

func main() {
	store := getStore()
	app := application.New(store)
	http := api.New(app)
	http.Start()
}

// The store has two identical schemas, one for production and one for testing.
// This function returns the appropriate store based on the environment.
func getStore() st.Store {
	testEnv := os.Getenv("PERSONACMMSTESTENV")
	if testEnv == "true" {
		return st.NewTest()
	}

	return st.New()
}
