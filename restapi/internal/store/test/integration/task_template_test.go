package integration

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestTaskTemplateCreate(t *testing.T) {
	t.Parallel()
	dbName := "testtasktemplatecreate"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Create
	taskTemplate := tp.TaskTemplate{
		Title:       "testtasktemplate1",
		Description: utest.ToPtr("test description"),
		Type:        utest.ToPtr(tp.TaskTypePreventative),
	}

	returnTaskTemplate, err := store.CreateTaskTemplate(taskTemplate)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	fieldsToExclude := utest.ConvertToSet([]string{"Id"})
	compEntitiesExcludeFields(t, taskTemplate, returnTaskTemplate, fieldsToExclude)
}

func TestTaskTemplateDelete(t *testing.T) {
	t.Parallel()
	dbName := "testtasktemplatedelete"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Delete
	taskTemplate := tp.TaskTemplate{
		Title:       "testtasktemplate1",
		Description: utest.ToPtr("test description"),
	}
	_, err := store.CreateTaskTemplate(taskTemplate)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	err = store.DeleteTaskTemplate(taskTemplate.Title)
	if err != nil {
		t.Errorf("Delete() failed: %v", err)
	}

	// Get
	_, err = store.GetTaskTemplate(taskTemplate.Title)
	if err == nil {
		t.Errorf("Get() failed: expected error, got nil")
	}
}

func TestTaskTemplateDeleteNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testtasktemplatedeletenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	err := store.DeleteTaskTemplate("notfound")
	if err == nil {
		t.Errorf("DeleteTaskTemplate() should have failed")
	}
}

func TestTaskTemplateList(t *testing.T) {
	t.Parallel()
	dbName := "testtasktemplatelist"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// List
	taskTemplates, err := store.ListTaskTemplates()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	// create a map of the taskTemplates title: tp.TaskTemplate
	taskTemplateMap := make(map[string]tp.TaskTemplate)
	for _, taskTemplate := range taskTemplates {
		taskTemplateMap[taskTemplate.Title] = taskTemplate
	}

	count := len(taskTemplates)

	taskTemplateMap["testtasktemplate1"] = tp.TaskTemplate{
		Title:       "testtasktemplate1",
		Description: utest.ToPtr("test description"),
		Type:        utest.ToPtr(tp.TaskTypePreventative),
	}

	taskTemplateMap["testtasktemplate2"] = tp.TaskTemplate{
		Title:       "testtasktemplate2",
		Description: utest.ToPtr("test description"),
		Type:        utest.ToPtr(tp.TaskTypeCorrective),
	}

	// Create the taskTemplates
	_, err = store.CreateTaskTemplate(taskTemplateMap["testtasktemplate1"])
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	_, err = store.CreateTaskTemplate(taskTemplateMap["testtasktemplate2"])
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	// List
	newTaskTemplates, err := store.ListTaskTemplates()
	if err != nil {
		t.Errorf("List() failed: %v", err)
	}

	if len(newTaskTemplates) != count+2 {
		t.Errorf("List() failed: expected %d, got %d", count+2, len(newTaskTemplates))
	}

	newTaskTemplateMap := make(map[string]tp.TaskTemplate)
	for _, taskTemplate := range newTaskTemplates {
		newTaskTemplateMap[taskTemplate.Title] = taskTemplate
	}

	for title, taskTemplate := range taskTemplateMap {
		fieldsToExclude := utest.ConvertToSet([]string{"Id"})
		compEntitiesExcludeFields(t, taskTemplate, newTaskTemplateMap[title], fieldsToExclude)
	}
}

func TestTaskTemplateUpdateGet(t *testing.T) {
	t.Parallel()
	dbName := "testtasktemplateupdateget"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	// Update
	taskTemplate := tp.TaskTemplate{
		Title:       "testtasktemplate1",
		Description: utest.ToPtr("test description"),
	}
	createTaskTemplate, err := store.CreateTaskTemplate(taskTemplate)
	if err != nil {
		t.Errorf("Create() failed: %v", err)
	}

	taskTemplate.Description = utest.ToPtr("new description")
	returnTaskTemplate, err := store.UpdateTaskTemplate(taskTemplate.Title, taskTemplate)
	if err != nil {
		t.Errorf("Update() failed: %v", err)
	}

	diffFields := utest.ConvertToSet([]string{"Description"})
	compEntitiesExcludeFields(t, createTaskTemplate, returnTaskTemplate, diffFields)

	getTaskTemplate, err := store.GetTaskTemplate(taskTemplate.Title)
	if err != nil {
		t.Errorf("Get() failed: %v", err)
	}

	compEntities(t, returnTaskTemplate, getTaskTemplate)
}

func TestTaskTemplateUpdateNotFound(t *testing.T) {
	t.Parallel()
	dbName := "testtasktemplateupdatenotfound"
	store := utest.InitializeStore(dbName)
	defer utest.CloseStore(store, dbName)

	taskTemplate := tp.TaskTemplate{
		Title:       "notfound",
		Description: utest.ToPtr("test description"),
	}
	_, err := store.UpdateTaskTemplate(taskTemplate.Title, taskTemplate)
	if err == nil {
		t.Errorf("UpdateTaskTemplate() should have failed")
	}
}
