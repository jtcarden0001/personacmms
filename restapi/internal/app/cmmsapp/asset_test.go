package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestCreateAsset(t *testing.T) {
	app := &App{db: mock.New()}

	a := utest.SetupAsset(1, false)
	_, err := app.CreateAsset(a)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}
}
