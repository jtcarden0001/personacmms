package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateUsageTrigger(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task := setupGroupAssetTask(app)
	usageTrigger := tp.UsageTrigger{TaskId: task.Id, UsageUnit: "hour", Quantity: 10}
	_, err := app.CreateUsageTrigger(group.Title, asset.Title, task.Id.String(), usageTrigger)
	assert.NoError(t, err)
}

func TestDeleteUsageTrigger(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task := setupGroupAssetTask(app)
	usageTrigger := tp.UsageTrigger{TaskId: task.Id, UsageUnit: "hour", Quantity: 10}
	usageTrigger, _ = db.CreateUsageTrigger(usageTrigger)

	err := app.DeleteUsageTrigger(group.Title, asset.Title, task.Id.String(), usageTrigger.Id.String())
	assert.NoError(t, err)
}

func TestListUsageTriggers(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task := setupGroupAssetTask(app)
	usageTrigger1 := tp.UsageTrigger{TaskId: task.Id, UsageUnit: "hour", Quantity: 10}
	usageTrigger2 := tp.UsageTrigger{TaskId: task.Id, UsageUnit: "minute", Quantity: 20}
	db.CreateUsageTrigger(usageTrigger1)
	db.CreateUsageTrigger(usageTrigger2)

	usageTriggers, err := app.ListUsageTriggers(group.Title, asset.Title, task.Id.String())
	assert.NoError(t, err)
	assert.Len(t, usageTriggers, 2)
}

func TestGetUsageTrigger(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task := setupGroupAssetTask(app)
	usageTrigger := tp.UsageTrigger{TaskId: task.Id, UsageUnit: "hour", Quantity: 10}
	usageTrigger, _ = db.CreateUsageTrigger(usageTrigger)

	_, err := app.GetUsageTrigger(group.Title, asset.Title, task.Id.String(), usageTrigger.Id.String())
	assert.NoError(t, err)
}

func TestUpdateUsageTrigger(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group, asset, task := setupGroupAssetTask(app)
	usageTrigger := tp.UsageTrigger{TaskId: task.Id, UsageUnit: "hour", Quantity: 10}
	usageTrigger, _ = db.CreateUsageTrigger(usageTrigger)

	usageTrigger.UsageUnit = "minute"
	usageTrigger.Quantity = 20
	_, err := app.UpdateUsageTrigger(group.Title, asset.Title, task.Id.String(), usageTrigger.Id.String(), usageTrigger)
	assert.NoError(t, err)
}
