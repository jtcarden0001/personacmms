package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetNotExists(t *testing.T) {
	_, err := testStore.GetAsset(0)
	errString := "sql: no rows in result set"
	if err.Error() != errString {
		t.Errorf("Get() should have failed with error: %s", errString)
	}
}

func TestDeleteNotExists(t *testing.T) {
	err := testStore.DeleteAsset(0)
	if err != nil {
		t.Errorf("Delete() record that does not exist should produce no error")
	}
}

// TODO: parse out the initialization code so that we only wipe the db and initialize foreign keys once per test run.
func TestCreateGetDeleteGet(t *testing.T) {
	// initialization
	ecId, err := initAssetFKs()
	if err != nil {
		t.Errorf("initAssetFKs() failed: %v", err)
		return
	}

	//// Test
	// Create
	id, err := testStore.CreateAsset("test asset", 2023, "test make", "test model number", "test description", ecId)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get
	_, err = testStore.GetAsset(id)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteAsset(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = testStore.GetAsset(id)
	if err == nil {
		t.Errorf("Get() should have failed")
	}

	err = teardownAsset(id, ecId)
	if err != nil {
		t.Errorf("teardownAsset() failed: %v", err)
	}
}

func TestGetAllCreateCreateGetAll(t *testing.T) {
	// initialization
	ecId, err := initAssetFKs()
	if err != nil {
		t.Errorf("initAssetFKs() failed: %v", err)
		return
	}

	//// Test
	// Get all
	e1, err := testStore.GetAllAsset()
	if err != nil {
		t.Errorf("GetAll() failed: %v", err)
	}
	origLength := len(e1)

	// Create
	id1, err := testStore.CreateAsset("test asset 1", 2023, "test make 1", "test model number 1", "test description 1", ecId)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Create
	var id2 int
	id2, err = testStore.CreateAsset("test asset 2", 2023, "test make 2", "test model number 2", "test description 2", ecId)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get all
	var e2 []tp.Asset
	e2, err = testStore.GetAllAsset()
	if err != nil {
		t.Errorf("GetAll() failed: %v", err)
	}
	assert.Equal(t, origLength+2, len(e2))

	// Delete
	err = testStore.DeleteAsset(id1)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteAsset(id2)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	err = teardownAsset(id1, ecId)
	if err != nil {
		t.Errorf("teardownAsset() failed: %v", err)
	}
}

func initAssetFKs() (int, error) {
	ecId, err := testStore.CreateCategory("test asset category")
	if err != nil {
		return 0, err
	}

	return int(ecId.ID()), nil
}

func teardownAsset(id int, ecId int) error {
	err := teardownTable("asset", &id)
	if err != nil {
		return err
	}

	err = teardownTable("asset_category", &ecId)
	return err
}
