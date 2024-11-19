package test

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestTaskCreate(t *testing.T) {
	store := InitializeStore("testtaskcreate")

	// Create
	task := types.Task{
		Title:       "testtask1",
		Description: "test description",
	}

	returntask, err := store.CreateTask(task)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	if returntask.Title != task.Title {
		t.Errorf("Create() failed: expected %s, got %s", task.Title, returntask.Title)
	}

	if returntask.Description != task.Description {
		t.Errorf("Create() failed: expected %s, got %s", task.Description, returntask.Description)
	}
}

func TestTaskDelete(t *testing.T) {
	store := InitializeStore("testtaskdelete")

	// Delete
	task := types.Task{
		Title:       "testtask1",
		Description: "test description",
	}
	_, err := store.CreateTask(task)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	err = store.DeleteTask(task.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = store.GetTask(task.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestTaskList(t *testing.T) {
	store := InitializeStore("testtasklist")

	// List
	tasks, err := store.ListTasks()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("ListTask() failed: expected 0, got %d", len(tasks))
	}

	// Create
	task := types.Task{
		Title:       "testtask1",
		Description: "test description",
	}
	_, err = store.CreateTask(task)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	task.Title = "testtask2"
	_, err = store.CreateTask(task)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	tasks, err = store.ListTasks()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("ListTask() failed: expected 2, got %d", len(tasks))
	}
}

func TestTaskUpdateGet(t *testing.T) {
	store := InitializeStore("testtaskupdate")

	// Update
	task := types.Task{
		Title:       "testtask1",
		Description: "test description",
	}
	_, err := store.CreateTask(task)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	task.Description = "new description"
	returntask, err := store.UpdateTask(task.Title, task)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	gettask, err := store.GetTask(task.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	if returntask.Title != task.Title || gettask.Title != task.Title {
		t.Errorf("Update() failed: expected %s, got %s", task.Title, returntask.Title)
	}

	if returntask.Description != task.Description || gettask.Description != task.Description {
		t.Errorf("Update() failed: expected %s, got %s", task.Description, returntask.Description)
	}
}
