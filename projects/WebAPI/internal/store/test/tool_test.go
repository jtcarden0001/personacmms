package test

import (
	"fmt"
	"testing"
)

func TestCreateUpdateDelete(t *testing.T) {
	// Create
	id, err := testStore.CreateTool("test tool")
	fmt.Println("test tool id:", id)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	err = testStore.UpdateTool(id, "test tool updated")
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Delete
	err = testStore.DeleteTool(id)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	err = testStore.ResetSequenceTool(id)
	if err != nil {
		t.Errorf("ResetSequence() failed: %v", err)
	}
}
