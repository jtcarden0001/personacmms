package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimeUnit(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	timeUnit := tp.TimeUnit{Title: "day"}
	_, err := app.CreateTimeUnit(timeUnit)
	assert.NoError(t, err)
}

func TestDeleteTimeUnit(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	timeUnit := tp.TimeUnit{Title: "day"}
	timeUnit, _ = db.CreateTimeUnit(timeUnit)

	err := app.DeleteTimeUnit(timeUnit.Title)
	assert.NoError(t, err)
}

func TestListTimeUnits(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	timeUnit1 := tp.TimeUnit{Title: "day"}
	timeUnit2 := tp.TimeUnit{Title: "week"}
	db.CreateTimeUnit(timeUnit1)
	db.CreateTimeUnit(timeUnit2)

	timeUnits, err := app.ListTimeUnits()
	assert.NoError(t, err)
	assert.Len(t, timeUnits, 2)
}

func TestGetTimeUnit(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	timeUnit := tp.TimeUnit{Title: "day"}
	timeUnit, _ = db.CreateTimeUnit(timeUnit)

	_, err := app.GetTimeUnit(timeUnit.Title)
	assert.NoError(t, err)
}

func TestUpdateTimeUnit(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	timeUnit := tp.TimeUnit{Title: "day"}
	timeUnit, _ = db.CreateTimeUnit(timeUnit)

	timeUnit.Title = "week"
	_, err := app.UpdateTimeUnit("day", timeUnit)
	assert.NoError(t, err)
}
