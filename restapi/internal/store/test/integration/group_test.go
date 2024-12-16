package integration

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestGroupCreate(t *testing.T) {
	t.Parallel()
	dbname := "testgroupcreate"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Create
	group := types.Group{
		Title: "testgroup1",
	}

	returnGroup, err := store.CreateGroup(group)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returnGroup.Title != group.Title {
		t.Errorf("Create() failed: expected %s, got %s", group.Title, returnGroup.Title)
	}

	if returnGroup.Id == uuid.Nil {
		t.Errorf("Create() failed: expected non-nil id")
	}
}

func TestGroupDelete(t *testing.T) {
	t.Parallel()
	dbname := "testgroupdelete"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Delete
	group := types.Group{
		Title: "testgroup1",
	}
	_, err := store.CreateGroup(group)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	err = store.DeleteGroup(group.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = store.GetGroup(group.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestGroupDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbname := "testgroupdeletenotfound"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	err := store.DeleteGroup("notfound")
	if err == nil {
		t.Errorf("DeleteGroup() should have failed")
	}
}

func TestGroupList(t *testing.T) {
	t.Parallel()
	dbname := "testgrouplist"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// List
	groups, err := store.ListGroups()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(groups) != 0 {
		t.Errorf("List() failed: expected 0, got %d", len(groups))
	}

	// Create
	group := types.Group{
		Title: "testgroup1",
	}
	_, err = store.CreateGroup(group)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	group.Title = "testgroup2"
	_, err = store.CreateGroup(group)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	groups, err = store.ListGroups()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(groups) != 2 {
		t.Errorf("List() failed: expected 2, got %d", len(groups))
	}
}

func TestGroupUpdateGet(t *testing.T) {
	t.Parallel()
	dbname := "testgroupupdateget"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	// Update
	group := types.Group{
		Title: "testgroup1",
	}
	_, err := store.CreateGroup(group)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	group.Title = "testgroup2"
	returnGroup, err := store.UpdateGroup("testgroup1", group)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Get
	getGroup, err := store.GetGroup(group.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if returnGroup.Title != group.Title || getGroup.Title != group.Title {
		t.Errorf("Get() failed: expected %s, got %s", group.Title, returnGroup.Title)
	}
}

func TestGroupUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbname := "testgroupupdatenotfound"
	store := initializeStore(dbname)
	defer closeStore(store, dbname)

	group := types.Group{
		Title: "notfound",
	}
	_, err := store.UpdateGroup("notfound", group)
	if err == nil {
		t.Errorf("UpdateGroup() should have failed")
	}
}
