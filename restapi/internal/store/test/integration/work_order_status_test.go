package integration

import (
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestCreateWorkOrderStatus(t *testing.T) {
	t.Parallel()
	dbName := "testcreateworkorderstatus"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	workOrderStatus := tp.WorkOrderStatus{
		Title: "testworkorderstatus1",
	}

	returnedStat, err := store.CreateWorkOrderStatus(workOrderStatus)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returnedStat.Title != workOrderStatus.Title {
		t.Errorf("Create() failed: expected %s, got %s", workOrderStatus.Title, returnedStat.Title)
	}

	if returnedStat.Id == uuid.Nil {
		t.Errorf("Create() failed: expected non-empty ID, got empty")
	}
}

func TestDeleteWorkOrderStatus(t *testing.T) {
	t.Parallel()
	dbName := "testdeleteworkorderstatus"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	workOrderStatus := tp.WorkOrderStatus{
		Title: "testworkorderstatus1",
	}

	_, err := store.CreateWorkOrderStatus(workOrderStatus)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Delete
	err = store.DeleteWorkOrderStatus(workOrderStatus.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Confirm deletion
	_, err = store.GetWorkOrderStatus(workOrderStatus.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestListWorkOrderStatus(t *testing.T) {
	t.Parallel()
	dbName := "testlistworkorderstatus"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	initWOS, err := store.ListWorkOrderStatuses()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	count := len(initWOS)

	// Create
	workOrderStatus := tp.WorkOrderStatus{
		Title: "testworkorderstatus1",
	}

	_, err = store.CreateWorkOrderStatus(workOrderStatus)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	workOrderStatus.Title = "testworkorderstatus2"
	_, err = store.CreateWorkOrderStatus(workOrderStatus)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	workOrderStatuses, err := store.ListWorkOrderStatuses()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(workOrderStatuses) != 2+count {
		t.Errorf("List() failed: expected at least 1, got %d", len(workOrderStatuses))
	}
}

func TestGetWorkOrderStatus(t *testing.T) {
	t.Parallel()
	dbName := "testgetworkorderstatus"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	workOrderStatus := tp.WorkOrderStatus{
		Title: "testworkorderstatus1",
	}

	_, err := store.CreateWorkOrderStatus(workOrderStatus)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get
	returnedStat, err := store.GetWorkOrderStatus(workOrderStatus.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if returnedStat.Title != workOrderStatus.Title {
		t.Errorf("Get() failed: expected %s, got %s", workOrderStatus.Title, returnedStat.Title)
	}

	if returnedStat.Id == uuid.Nil {
		t.Errorf("Get() failed: expected non-empty ID, got empty")
	}
}

func TestUpdateWorkOrderStatus(t *testing.T) {
	t.Parallel()
	dbName := "testupdateworkorderstatus"
	store := initializeStore(dbName)
	defer closeStore(store, dbName)

	// Create
	workOrderStatus := tp.WorkOrderStatus{
		Title: "testworkorderstatus1",
	}

	_, err := store.CreateWorkOrderStatus(workOrderStatus)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	workOrderStatus.Title = "testworkorderstatus2"
	returnedStat, err := store.UpdateWorkOrderStatus("testworkorderstatus1", workOrderStatus)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	if returnedStat.Title != workOrderStatus.Title {
		t.Errorf("Update() failed: expected %s, got %s", workOrderStatus.Title, returnedStat.Title)
	}
}
