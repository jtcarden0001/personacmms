package postgres_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func setupWorkOrder(identifier int) tp.WorkOrder {
	return tp.WorkOrder{
		Id:              uuid.New(),
		Title:           "Work Order " + strconv.Itoa(identifier),
		CreatedDate:     time.Now().AddDate(0, 0, -identifier),
		CompletedDate:   utest.ToPtr(time.Now().AddDate(0, 0, identifier)),
		Instructions:    utest.ToPtr("Instructions " + strconv.Itoa(identifier)),
		Notes:           utest.ToPtr("Notes " + strconv.Itoa(identifier)),
		CumulativeMiles: utest.ToPtr(identifier * 100),
		CumulativeHours: utest.ToPtr(identifier * 10),
		Status:          tp.WorkOrderStatusComplete,
	}
}

func TestWorkOrderCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testworkordercreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	wo := setupWorkOrder(1)

	// test
	createdWorkOrder, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}

	utest.CompEntities(t, wo, createdWorkOrder)
}

func TestWorkOrderDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testworkorderdelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	wo := setupWorkOrder(1)
	createdWorkOrder, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("TestWorkOrderDelete: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	// test
	err = store.DeleteWorkOrder(createdWorkOrder.Id)
	if err != nil {
		t.Errorf("DeleteWorkOrder() failed: %v", err)
	}

	_, err = store.GetWorkOrder(createdWorkOrder.Id)
	if err == nil {
		t.Errorf("GetWorkOrder() returned nil error after deletion")
	}
}

func TestWorkOrderGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testworkorderget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	wo := setupWorkOrder(1)
	createWorkOrder, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("TestWorkOrderGet: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	// test
	getWorkOrder, err := store.GetWorkOrder(createWorkOrder.Id)
	if err != nil {
		t.Errorf("GetWorkOrder() failed: %v", err)
	}

	utest.CompEntities(t, wo, getWorkOrder)
}

func TestWorkOrderList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testworkorderlist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	wo1 := setupWorkOrder(1)
	wo2 := setupWorkOrder(2)
	wo3 := setupWorkOrder(3)

	_, err := store.CreateWorkOrder(wo1)
	if err != nil {
		t.Errorf("TestWorkOrderList: failed during setup. CreateWorkOrder() failed: %v", err)
	}
	_, err = store.CreateWorkOrder(wo2)
	if err != nil {
		t.Errorf("TestWorkOrderList: failed during setup. CreateWorkOrder() failed: %v", err)
	}
	_, err = store.CreateWorkOrder(wo3)
	if err != nil {
		t.Errorf("TestWorkOrderList: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	// test
	workOrders, err := store.ListWorkOrders()
	if err != nil {
		t.Errorf("ListWorkOrders() failed: %v", err)
	}

	if len(workOrders) != 3 {
		t.Errorf("ListWorkOrders() failed: expected 3 work orders, got %d", len(workOrders))
	}

	workOrderMap := map[uuid.UUID]tp.WorkOrder{
		wo1.Id: wo1,
		wo2.Id: wo2,
		wo3.Id: wo3,
	}

	for _, workOrder := range workOrders {
		expectedWorkOrder, ok := workOrderMap[workOrder.Id]
		if !ok {
			t.Errorf("ListWorkOrders() failed: unexpected work order with ID %v", workOrder.Id)
		}
		utest.CompEntities(t, expectedWorkOrder, workOrder)
	}
}

func TestWorkOrderUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testworkorderupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	wo := setupWorkOrder(1)
	createWorkOrder, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("TestWorkOrderUpdate: failed during setup. CreateWorkOrder() failed: %v", err)
	}

	// test
	wo.Notes = utest.ToPtr("Updated Work Order Notes")
	wo.CumulativeMiles = utest.ToPtr(200)
	wo.CumulativeHours = utest.ToPtr(20)

	updatedWorkOrder, err := store.UpdateWorkOrder(wo)
	if err != nil {
		t.Errorf("UpdateWorkOrder() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"Notes", "CumulativeMiles", "CumulativeHours"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createWorkOrder, updatedWorkOrder, differentFields)
}
