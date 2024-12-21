package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateWorkOrderStatus(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	workOrderStatus := tp.WorkOrderStatus{Title: "New"}
	_, err := app.CreateWorkOrderStatus(workOrderStatus)
	assert.NoError(t, err)
}

func TestDeleteWorkOrderStatus(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	workOrderStatus := tp.WorkOrderStatus{Title: "New"}
	workOrderStatus, _ = db.CreateWorkOrderStatus(workOrderStatus)

	err := app.DeleteWorkOrderStatus(workOrderStatus.Title)
	assert.NoError(t, err)
}

func TestListWorkOrderStatuses(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	workOrderStatus1 := tp.WorkOrderStatus{Title: "New"}
	workOrderStatus2 := tp.WorkOrderStatus{Title: "In Progress"}
	db.CreateWorkOrderStatus(workOrderStatus1)
	db.CreateWorkOrderStatus(workOrderStatus2)

	workOrderStatuses, err := app.ListWorkOrderStatus()
	assert.NoError(t, err)
	assert.Len(t, workOrderStatuses, 2)
}

func TestGetWorkOrderStatus(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	workOrderStatus := tp.WorkOrderStatus{Title: "New"}
	workOrderStatus, _ = db.CreateWorkOrderStatus(workOrderStatus)

	_, err := app.GetWorkOrderStatus(workOrderStatus.Title)
	assert.NoError(t, err)
}

func TestUpdateWorkOrderStatus(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	workOrderStatus := tp.WorkOrderStatus{Title: "New"}
	workOrderStatus, _ = db.CreateWorkOrderStatus(workOrderStatus)

	workOrderStatus.Title = "In Progress"
	_, err := app.UpdateWorkOrderStatus("New", workOrderStatus)
	assert.NoError(t, err)
}
