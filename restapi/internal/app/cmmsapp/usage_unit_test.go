package cmmsapp

// import (
// 	"testing"

// 	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
// 	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateUsageUnit(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	usageUnit := tp.UsageUnit{Title: "hour", Type: "time"}
// 	_, err := app.CreateUsageUnit(usageUnit)
// 	assert.NoError(t, err)
// }

// func TestDeleteUsageUnit(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	usageUnit := tp.UsageUnit{Title: "hour", Type: "time"}
// 	usageUnit, _ = db.CreateUsageUnit(usageUnit)

// 	err := app.DeleteUsageUnit(usageUnit.Title)
// 	assert.NoError(t, err)
// }

// func TestListUsageUnits(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	usageUnit1 := tp.UsageUnit{Title: "hour", Type: "time"}
// 	usageUnit2 := tp.UsageUnit{Title: "minute", Type: "time"}
// 	db.CreateUsageUnit(usageUnit1)
// 	db.CreateUsageUnit(usageUnit2)

// 	usageUnits, err := app.ListUsageUnits()
// 	assert.NoError(t, err)
// 	assert.Len(t, usageUnits, 2)
// }

// func TestGetUsageUnit(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	usageUnit := tp.UsageUnit{Title: "hour", Type: "time"}
// 	usageUnit, _ = db.CreateUsageUnit(usageUnit)

// 	_, err := app.GetUsageUnit(usageUnit.Title)
// 	assert.NoError(t, err)
// }

// func TestUpdateUsageUnit(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	usageUnit := tp.UsageUnit{Title: "hour", Type: "time"}
// 	usageUnit, _ = db.CreateUsageUnit(usageUnit)

// 	usageUnit.Title = "minute"
// 	_, err := app.UpdateUsageUnit("hour", usageUnit)
// 	assert.NoError(t, err)
// }
