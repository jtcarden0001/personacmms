package integration

import (
	"log"
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/app"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

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

	// run tests
	m.Run()
}

func NewApp(name string) app.App {
	store := utest.InitializeStore(name)
	return app.New(store)
}
