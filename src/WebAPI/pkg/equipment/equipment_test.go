package equipment

import "testing"

func TestCreateGetDeleteGet(t *testing.T) {
	// Create
	id, err := Create("test equipment", "test description")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Get
	_, err = Get(id)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	// Delete
	err = Delete(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = Get(id)
	if err == nil {
		t.Errorf("Get() should have failed")
	}
}
