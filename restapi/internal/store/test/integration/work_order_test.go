package integration

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestWorkOrderCreate(t *testing.T) {
	t.Parallel()
	dbName := "testworkordercreate"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	atId := setupTask(t, store, "1")
	wo := getTestWorkOrder(atId, "1")

	// Create a new work order
	returnedWo, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}

	fieldsToExclude := convertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, wo, returnedWo, fieldsToExclude)
}

func TestWorkOrderDelete(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderdelete"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	atId := setupTask(t, store, "1")
	wo := getTestWorkOrder(atId, "1")

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

func TestWorkOrderList(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderlist"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// List
	wos, err := store.ListWorkOrders()
	if err != nil {
		t.Errorf("ListWorkOrders() failed: %v", err)
	}
	// make a map
	wosMap := make(map[tp.UUID]tp.WorkOrder)
	for _, wo := range wos {
		wosMap[wo.Id] = wo
	}

	atId := setupTask(t, store, "1")

	// Create a new work order
	wo1 := getTestWorkOrder(atId, "1")
	cwo1, err := store.CreateWorkOrder(wo1)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}
	wosMap[cwo1.Id] = cwo1

	wo2 := getTestWorkOrder(atId, "2")
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
		compEntities(t, wo, wosMap[wo.Id])
	}
}

func TestWorkOrderUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testworkorderupdateget"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	atId := setupTask(t, store, "1")
	wo := getTestWorkOrder(atId, "1")

	// Create a new work order
	cwo, err := store.CreateWorkOrder(wo)
	if err != nil {
		t.Errorf("CreateWorkOrder() failed: %v", err)
	}

	// Update the work order
	wo = getTestWorkOrder(atId, "2")
	wo.Id = cwo.Id

	updatedWo, err := store.UpdateWorkOrder(cwo.Id, wo)
	if err != nil {
		t.Errorf("UpdateWorkOrder() failed: %v", err)
	}

	fieldsShouldBeDifferent := convertToSet([]string{"CreatedDate", "CompletedDate", "Notes", "CumulativeMiles", "CumulativeHours"})
	compEntitiesFieldsShouldBeDifferent(t, cwo, updatedWo, fieldsShouldBeDifferent)

	// Get the work order
	returnedWo, err := store.GetWorkOrder(updatedWo.Id)
	if err != nil {
		t.Errorf("GetWorkOrder() failed: %v", err)
	}

	compEntities(t, updatedWo, returnedWo)
}

// different values except assetTaskId and statusTitle
func getTestWorkOrder(assetTaskId tp.UUID, id string) tp.WorkOrder {
	return tp.WorkOrder{
		CreatedDate:     time.Now(),
		CompletedDate:   func(t time.Time) *time.Time { return &t }(time.Now().UTC().Truncate(time.Millisecond)),
		Notes:           func(s string) *string { return &s }(fmt.Sprintf("Test work order %s", id)),
		CumulativeMiles: func() *int { miles, _ := strconv.Atoi(fmt.Sprintf("20%s", id)); return &miles }(),
		CumulativeHours: func() *int { hours, _ := strconv.Atoi(fmt.Sprintf("20%s", id)); return &hours }(),
		TaskId:          assetTaskId,
		StatusTitle:     tp.WorkOrderStatusComplete,
	}
}
