package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestGroupCreate(t *testing.T) {
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

func TestGroupList(t *testing.T) {
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
