package cmmsapp

// import (
// 	"testing"

// 	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
// 	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateWorkOrder(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	workOrder := tp.WorkOrder{TaskId: task.Id, StatusTitle: "new"}
// 	_, err := app.CreateWorkOrder(group.Title, asset.Title, task.Id.String(), workOrder)
// 	assert.NoError(t, err)
// }

// func TestDeleteWorkOrder(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	workOrder := tp.WorkOrder{TaskId: task.Id, StatusTitle: "new"}
// 	workOrder, _ = db.CreateWorkOrder(workOrder)

// 	err := app.DeleteWorkOrder(group.Title, asset.Title, task.Id.String(), workOrder.Id.String())
// 	assert.NoError(t, err)
// }

// func TestListWorkOrders(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	workOrder1 := tp.WorkOrder{TaskId: task.Id, StatusTitle: "new"}
// 	workOrder2 := tp.WorkOrder{TaskId: task.Id, StatusTitle: "in progress"}
// 	db.CreateWorkOrder(workOrder1)
// 	db.CreateWorkOrder(workOrder2)

// 	workOrders, err := app.ListWorkOrders(group.Title, asset.Title, task.Id.String())
// 	assert.NoError(t, err)
// 	assert.Len(t, workOrders, 2)
// }

// func TestGetWorkOrder(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	workOrder := tp.WorkOrder{TaskId: task.Id, StatusTitle: "new"}
// 	workOrder, _ = db.CreateWorkOrder(workOrder)

// 	_, err := app.GetWorkOrder(group.Title, asset.Title, task.Id.String(), workOrder.Id.String())
// 	assert.NoError(t, err)
// }

// func TestUpdateWorkOrder(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	group, asset, task := setupGroupAssetTask(app)
// 	workOrder := tp.WorkOrder{TaskId: task.Id, StatusTitle: "new"}
// 	workOrder, _ = db.CreateWorkOrder(workOrder)

// 	workOrder.StatusTitle = "in progress"
// 	_, err := app.UpdateWorkOrder(group.Title, asset.Title, task.Id.String(), workOrder.Id.String(), workOrder)
// 	assert.NoError(t, err)
// }
