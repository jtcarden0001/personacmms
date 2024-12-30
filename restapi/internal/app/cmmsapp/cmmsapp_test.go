package cmmsapp

import (
	"strings"
	"testing"

	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
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

func initializeAppTest(t *testing.T, testName string) (*App, func(), error) {
	t.Parallel()
	store := utest.InitializeStore(strings.ToLower(testName))
	app := New(store)
	cleanup := func() {
		utest.CloseStore(store, testName)
	}

	return app, cleanup, nil
}
