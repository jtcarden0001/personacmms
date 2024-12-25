package integration

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestWorkOrderCreate(t *testing.T) {
	t.Parallel()
	dbName := "testworkordercreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	atId := setupTask(t, store, "1")
	wo := getTestWorkOrder(atId, 1)

	// Create a new work order
	returnedWo, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}

	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, wo, returnedWo, fieldsToExclude)
}

func TestWorkOrderDelete(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderdelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	atId := setupTask(t, store, "1")
	wo := getTestWorkOrder(atId, 1)

	// Create a new work order
	returnedWo, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}

	// Delete the work order
	err = store.DeleteWorkOrder(returnedWo.Id)
	if err != nil {
		t.Errorf("DeleteWorkOrder() failed: %v", err)
	}

	// Confirm deletion
	_, err = store.GetWorkOrder(returnedWo.Id)
	if err == nil {
		t.Errorf("GetWorkOrder() failed: expected error, got nil")
	}
}

func TestWorkOrderDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderdeletenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	err := store.DeleteWorkOrder(uuid.New())
	if err == nil {
		t.Errorf("DeleteWorkOrder() failed: expected error, got nil")
	}
}

func TestWorkOrderList(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderlist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// List
	wos, err := store.ListWorkOrders()
	if err != nil {
		t.Errorf("ListWorkOrders() failed: %v", err)
	}
	// make a map
	wosMap := make(map[uuid.UUID]tp.WorkOrder)
	for _, wo := range wos {
		wosMap[wo.Id] = wo
	}

	atId := setupTask(t, store, "1")

	// Create a new work order
	wo1 := getTestWorkOrder(atId, 1)
	cwo1, err := store.CreateWorkOrder(wo1)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}
	wosMap[cwo1.Id] = cwo1

	wo2 := getTestWorkOrder(atId, 2)
	cwo2, err := store.CreateWorkOrder(wo2)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}
	wosMap[cwo2.Id] = cwo2

	// List all work orders
	workOrders, err := store.ListWorkOrders()
	if err != nil {
		t.Errorf("ListWorkOrders() failed: %v", err)
	}

	if len(workOrders) != len(wosMap) {
		t.Errorf("ListWorkOrders() failed: expected %d work orders, got %d", len(wosMap), len(workOrders))
	}

	for _, wo := range workOrders {
		utest.CompEntities(t, wo, wosMap[wo.Id])
	}
}

func TestWorkOrderUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderupdateget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	atId := setupTask(t, store, "1")
	wo1 := getTestWorkOrder(atId, 1)

	// Create a new work order
	cwo, err := store.CreateWorkOrder(wo1)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}

	// Update the work order
	wo2 := getTestWorkOrder(atId, 2)
	wo2.Id = cwo.Id

	updatedWo, err := store.UpdateWorkOrder(cwo.Id, wo2)
	if err != nil {
		t.Errorf("UpdateWorkOrder() failed: %v", err)
	}

	fieldsShouldBeDifferent := utest.ConvertToSet([]string{"CreatedDate", "CompletedDate", "CumulativeMiles", "CumulativeHours"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, cwo, updatedWo, fieldsShouldBeDifferent)

	// Get the work order
	returnedWo, err := store.GetWorkOrder(updatedWo.Id)
	if err != nil {
		t.Errorf("GetWorkOrder() failed: %v", err)
	}

	utest.CompEntities(t, updatedWo, returnedWo)
}

func TestWorkOrderUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderupdatenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	wo := getTestWorkOrder(uuid.New(), 1)
	_, err := store.UpdateWorkOrder(uuid.New(), wo)
	if err == nil {
		t.Errorf("UpdateWorkOrder() failed: expected error, got nil")
	}
}

// different values except assetTaskId and statusTitle
func getTestWorkOrder(assetTaskId uuid.UUID, id int) tp.WorkOrder {
	return tp.WorkOrder{
		CreatedDate:     time.Now().Add(time.Duration(id) * time.Hour),
		CompletedDate:   func(t time.Time) *time.Time { return &t }(time.Now().Add(time.Duration(id) * time.Hour).UTC().Truncate(time.Millisecond)),
		Notes:           nil, // func(s string) *string { return &s }(fmt.Sprintf("Test work order %s", id)),
		CumulativeMiles: func() *int { miles, _ := strconv.Atoi(fmt.Sprintf("20%d", id)); return &miles }(),
		CumulativeHours: func() *int { hours, _ := strconv.Atoi(fmt.Sprintf("20%d", id)); return &hours }(),
		TaskId:          assetTaskId,
		StatusTitle:     tp.WorkOrderStatusComplete,
	}
}
