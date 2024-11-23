package test

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestPreventativeTaskCreate(t *testing.T) {
	store := InitializeStore("testpreventativetaskcreate")

	// Create
	preventativeTask := types.PreventativeTask{
		Title:       "testpreventativeTask1",
		Description: "test description",
	}

	returnpreventativeTask, err := store.CreatePreventativeTask(preventativeTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returnpreventativeTask.Title != preventativeTask.Title {
		t.Errorf("Create() failed: expected %s, got %s", preventativeTask.Title, returnpreventativeTask.Title)
	}

	if returnpreventativeTask.Description != preventativeTask.Description {
		t.Errorf("Create() failed: expected %s, got %s", preventativeTask.Description, returnpreventativeTask.Description)
	}
}

func TestPreventativeTaskDelete(t *testing.T) {
	store := InitializeStore("testpreventativetaskdelete")

	// Delete
	preventativeTask := types.PreventativeTask{
		Title:       "testpreventativeTask1",
		Description: "test description",
	}
	_, err := store.CreatePreventativeTask(preventativeTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	err = store.DeletePreventativeTask(preventativeTask.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = store.GetPreventativeTask(preventativeTask.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestPreventativeTaskList(t *testing.T) {
	store := InitializeStore("testpreventativetasklist")

	// List
	preventativeTasks, err := store.ListPreventativeTasks()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(preventativeTasks) != 0 {
		t.Errorf("ListPreventativeTask() failed: expected 0, got %d", len(preventativeTasks))
	}

	// Create
	preventativeTask := types.PreventativeTask{
		Title:       "testpreventativeTask1",
		Description: "test description",
	}
	_, err = store.CreatePreventativeTask(preventativeTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	preventativeTask.Title = "testpreventativeTask2"
	_, err = store.CreatePreventativeTask(preventativeTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	preventativeTasks, err = store.ListPreventativeTasks()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(preventativeTasks) != 2 {
		t.Errorf("ListPreventativeTask() failed: expected 2, got %d", len(preventativeTasks))
	}
}

func TestPreventativeTaskUpdateGet(t *testing.T) {
	store := InitializeStore("testpreventativetaskupdate")

	// Update
	preventativeTask := types.PreventativeTask{
		Title:       "testpreventativeTask1",
		Description: "test description",
	}
	_, err := store.CreatePreventativeTask(preventativeTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	preventativeTask.Description = "new description"
	returnpreventativeTask, err := store.UpdatePreventativeTask(preventativeTask.Title, preventativeTask)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	getpreventativeTask, err := store.GetPreventativeTask(preventativeTask.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if returnpreventativeTask.Title != preventativeTask.Title || getpreventativeTask.Title != preventativeTask.Title {
		t.Errorf("Update() failed: expected %s, got %s", preventativeTask.Title, returnpreventativeTask.Title)
	}

	if returnpreventativeTask.Description != preventativeTask.Description || getpreventativeTask.Description != preventativeTask.Description {
		t.Errorf("Update() failed: expected %s, got %s", preventativeTask.Description, returnpreventativeTask.Description)
	}
}
