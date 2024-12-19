package integration

import (
	"log"
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/app"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

var cmmsapp app.App

func TestMain(m *testing.M) {
	pool, resource, err := utest.CreateDockerTestPostgres()
	if err != nil {
		log.Fatalf("Could not create docker test postgres: %s", err)
	}
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
		resource.Close()
	}()

	store := utest.InitializeStore("cmmsapptest")
	cmmsapp = app.New(store)

	// run tests
	m.Run()
}
