package unit

import (
	"testing"

	a "github.com/jtcarden0001/personacmms/restapi/internal/app"
	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
)

var app a.App

func TestMain(m *testing.M) {
	// setup
	app = a.New(st.NewMock())

	m.Run()

	// teardown
}
