package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestCorrectiveTaskCreate(t *testing.T) {
	store := InitializeStore("testcorrectivetaskcreate")

	// Create
	correctiveTask := types.CorrectiveTask{
		Title:       "testcorrectiveTask1",
		Description: "test description",
	}

	returncorrectiveTask, err := store.CreateCorrectiveTask(correctiveTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returncorrectiveTask.Title != correctiveTask.Title {
		t.Errorf("Create() failed: expected %s, got %s", correctiveTask.Title, returncorrectiveTask.Title)
	}

	if returncorrectiveTask.Description != correctiveTask.Description {
		t.Errorf("Create() failed: expected %s, got %s", correctiveTask.Description, returncorrectiveTask.Description)
	}

	if returncorrectiveTask.Id == uuid.Nil {
		t.Errorf("Create() failed: expected non-empty ID, got empty")
	}
}

func TestCorrectiveTaskDelete(t *testing.T) {
	store := InitializeStore("testcorrectivetaskdelete")

	// Delete
	correctiveTask := types.CorrectiveTask{
		Title:       "testcorrectiveTask1",
		Description: "test description",
	}
	_, err := store.CreateCorrectiveTask(correctiveTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	err = store.DeleteCorrectiveTask(correctiveTask.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = store.GetCorrectiveTask(correctiveTask.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestCorrectiveTaskList(t *testing.T) {
	store := InitializeStore("testcorrectivetasklist")

	// List
	// Create multiple tasks
	tasks := []types.CorrectiveTask{
		{Title: "testcorrectiveTask1", Description: "test description 1"},
		{Title: "testcorrectiveTask2", Description: "test description 2"},
	}

	for _, task := range tasks {
		_, err := store.CreateCorrectiveTask(task)
		if err != nil {
			t.Errorf("Create() failed: %v", err)
		}
	}

	// List tasks
	returnTasks, err := store.ListCorrectiveTasks()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	// Check if the tasks are listed
	if len(returnTasks) != len(tasks) {
		t.Errorf("List() failed: expected %d tasks, got %d", len(tasks), len(returnTasks))
	}

	taskMap := make(map[string]types.CorrectiveTask)
	for _, task := range returnTasks {
		taskMap[task.Title] = task
	}

	for _, task := range tasks {
		if returnTask, ok := taskMap[task.Title]; !ok {
			t.Errorf("List() failed: task %s not found", task.Title)
		} else {
			if returnTask.Description != task.Description {
				t.Errorf("List() failed: expected %s, got %s", task.Description, returnTask.Description)
			}
		}
	}
}

func TestCorrectiveTaskUpdateGet(t *testing.T) {
	store := InitializeStore("testcorrectivetaskupdateget")

	// Create
	correctiveTask := types.CorrectiveTask{
		Title:       "testcorrectiveTask1",
		Description: "test description",
	}
	_, err := store.CreateCorrectiveTask(correctiveTask)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// Update
	correctiveTask.Description = "updated description"
	_, err = store.UpdateCorrectiveTask(correctiveTask.Title, correctiveTask)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	// Get
	returncorrectiveTask, err := store.GetCorrectiveTask(correctiveTask.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if returncorrectiveTask.Description != correctiveTask.Description {
		t.Errorf("Get() failed: expected %s, got %s", correctiveTask.Description, returncorrectiveTask.Description)
	}
}
