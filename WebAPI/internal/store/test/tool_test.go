package test

import (
	"testing"
)

func TestToolCreateUpdateDelete(t *testing.T) {
	// Create
	id, err := testStore.CreateTool("test tool", "test size")
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	err = testStore.UpdateTool(id, "test tool updated", "test size")
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteTool(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	err = testStore.ResetSequence("tool", id)
	if err != nil {
		t.Errorf("ResetSequence() failed: %v", err)
	}
}
