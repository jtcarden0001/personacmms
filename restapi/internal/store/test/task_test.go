package test

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func TestTaskCreate(t *testing.T) {
	store := InitializeStore("testpreventativetaskcreate")

	// Create
	task := tp.Task{
		Title:       "testtask1",
		Description: "test description",
		Type:        tp.TaskTypePreventative,
	}

	returntask, err := store.CreateTask(task)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	compareEntitiesExcludingId(t, task, returntask)
}

func TestTaskDelete(t *testing.T) {
	store := InitializeStore("testpreventativetaskdelete")

	// Delete
	task := tp.Task{
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
	store := InitializeStore("testpreventativetasklist")

	// List
	tasks, err := store.ListTasks()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	// create a map of the tasks title: tp.task
	taskMap := make(map[string]tp.Task)
	for _, task := range tasks {
		taskMap[task.Title] = task
	}

	count := len(tasks)

	taskMap["testtask1"] = tp.Task{
		Title:       "testtask1",
		Description: "test description",
		Type:        tp.TaskTypePreventative,
	}

	taskMap["testtask2"] = tp.Task{
		Title:       "testtask2",
		Description: "test description",
		Type:        tp.TaskTypeCorrective,
	}

	// Create the tasks
	_, err = store.CreateTask(taskMap["testtask1"])
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	_, err = store.CreateTask(taskMap["testtask2"])
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	newTasks, err := store.ListTasks()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(newTasks) != count+2 {
		t.Errorf("List() failed: expected %d, got %d", count+2, len(newTasks))
	}

	newTaskMap := make(map[string]tp.Task)
	for _, task := range newTasks {
		newTaskMap[task.Title] = task
	}

	for title, task := range taskMap {
		compareEntitiesExcludingId(t, task, newTaskMap[title])
	}
}

func TestTaskUpdateGet(t *testing.T) {
	store := InitializeStore("testpreventativetaskupdate")

	// Update
	task := tp.Task{
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
