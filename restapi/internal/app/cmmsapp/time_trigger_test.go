package cmmsapp

// import (
// 	"testing"

// 	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
// 	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateTimeTrigger(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	timeTrigger := tp.TimeTrigger{TaskId: task.Id, TimeUnit: "day", Quantity: 1}
// 	_, err := app.CreateTimeTrigger(group.Title, asset.Title, task.Id.String(), timeTrigger)
// 	assert.NoError(t, err)
// }

// func TestDeleteTimeTrigger(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	timeTrigger := tp.TimeTrigger{TaskId: task.Id, TimeUnit: "day", Quantity: 1}
// 	timeTrigger, _ = db.CreateTimeTrigger(timeTrigger)

// 	err := app.DeleteTimeTrigger(group.Title, asset.Title, task.Id.String(), timeTrigger.Id.String())
// 	assert.NoError(t, err)
// }

// func TestListTimeTriggers(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	timeTrigger1 := tp.TimeTrigger{TaskId: task.Id, TimeUnit: "day", Quantity: 1}
// 	timeTrigger2 := tp.TimeTrigger{TaskId: task.Id, TimeUnit: "week", Quantity: 2}
// 	db.CreateTimeTrigger(timeTrigger1)
// 	db.CreateTimeTrigger(timeTrigger2)

// 	timeTriggers, err := app.ListTimeTriggers(group.Title, asset.Title, task.Id.String())
// 	assert.NoError(t, err)
// 	assert.Len(t, timeTriggers, 2)
// }

// func TestGetTimeTrigger(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	timeTrigger := tp.TimeTrigger{TaskId: task.Id, TimeUnit: "day", Quantity: 1}
// 	timeTrigger, _ = db.CreateTimeTrigger(timeTrigger)

// 	_, err := app.GetTimeTrigger(group.Title, asset.Title, task.Id.String(), timeTrigger.Id.String())
// 	assert.NoError(t, err)
// }

// func TestUpdateTimeTrigger(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	timeTrigger := tp.TimeTrigger{TaskId: task.Id, TimeUnit: "day", Quantity: 1}
// 	timeTrigger, _ = db.CreateTimeTrigger(timeTrigger)

// 	timeTrigger.TimeUnit = "week"
// 	timeTrigger.Quantity = 2
// 	_, err := app.UpdateTimeTrigger(group.Title, asset.Title, task.Id.String(), timeTrigger.Id.String(), timeTrigger)
// 	assert.NoError(t, err)
// }
