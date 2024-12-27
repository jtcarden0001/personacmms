package postgres_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func setupGroup(identifier int) tp.Group {
	return tp.Group{
		Id:    uuid.New(),
		Title: fmt.Sprintf("Group %d", identifier),
	}
}

func TestGroupCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testgroupcreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	g := setupGroup(1)

	// test
	createdGroup, err := store.CreateGroup(g)
	if err != nil {
		t.Errorf("CreateGroup() failed: %v", err)
	}

	utest.CompEntities(t, g, createdGroup)
}

func TestGroupDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testgroupdelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	g := setupGroup(1)
	createdGroup, err := store.CreateGroup(g)
	if err != nil {
		t.Errorf("TestGroupDelete: failed during setup. CreateGroup() failed: %v", err)
	}

	// test
	err = store.DeleteGroup(createdGroup.Id)
	if err != nil {
		t.Errorf("TestGroupDelete: DeleteGroup() failed: %v", err)
	}

	_, err = store.GetGroup(createdGroup.Id)
	if err == nil {
		t.Errorf("TestGroupDelete: GetGroup() returned nil error after deletion")
	}
}

func TestGroupGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testgroupget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	g := setupGroup(1)
	createGroup, err := store.CreateGroup(g)
	if err != nil {
		t.Errorf("TestGroupGet: failed during setup. CreateGroup() failed: %v", err)
	}

	// test
	getGroup, err := store.GetGroup(createGroup.Id)
	if err != nil {
		t.Errorf("GetGroup() failed: %v", err)
	}

	utest.CompEntities(t, createGroup, getGroup)
}

func TestGroupList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testgrouplist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	g1 := setupGroup(1)
	g2 := setupGroup(2)
	g3 := setupGroup(3)

	_, err := store.CreateGroup(g1)
	if err != nil {
		t.Errorf("TestGroupList: failed during setup. CreateGroup() failed: %v", err)
	}
	_, err = store.CreateGroup(g2)
	if err != nil {
		t.Errorf("TestGroupList: failed during setup. CreateGroup() failed: %v", err)
	}
	_, err = store.CreateGroup(g3)
	if err != nil {
		t.Errorf("TestGroupList: failed during setup. CreateGroup() failed: %v", err)
	}

	// test
	groups, err := store.ListGroups()
	if err != nil {
		t.Errorf("ListGroups() failed: %v", err)
	}

	if len(groups) != 3 {
		t.Errorf("ListGroups() returned %d groups, expected 3", len(groups))
	}
}

func TestGroupUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testgroupupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)

	g := setupGroup(1)
	createdGroup, err := store.CreateGroup(g)
	if err != nil {
		t.Errorf("TestGroupUpdate: failed during setup. CreateGroup() failed: %v", err)
	}

	// test
	g.Title = "Updated Title"
	updatedGroup, err := store.UpdateGroup(g)
	if err != nil {
		t.Errorf("UpdateGroup() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"Title"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createdGroup, updatedGroup, differentFields)
}
