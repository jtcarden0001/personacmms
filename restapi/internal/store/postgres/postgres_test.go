package postgres_test

import (
	"testing"

	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	pool, resource, err := utest.CreateDockerTestPostgres()
	if err != nil {
		log.Fatalf("Could not create docker test postgres: %s", err)
	}
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
		resource.Close()
	}()

	// run tests
	m.Run()
}

// func setupGroup(t *testing.T, store st.Store, identifier string) string {
// 	group := tp.Group{
// 		Title: fmt.Sprintf("Group %s", identifier),
// 	}
// 	group, err := store.CreateGroup(group)
// 	if err != nil {
// 		t.Errorf("CreateGroup() failed: %v", err)
// 	}
// 	return group.Title
// }

// func setupCategory(t *testing.T, store st.Store, identifier string) string {
// 	category := tp.Category{
// 		Title:       fmt.Sprintf("Category %s", identifier),
// 		Description: utest.ToPtr(fmt.Sprintf("Category %s description", identifier)),
// 	}
// 	category, err := store.CreateCategory(category)
// 	if err != nil {
// 		t.Errorf("CreateCategory() failed: %v", err)
// 	}
// 	return category.Title
// }

// func setupConsumable(t *testing.T, store st.Store, identifier string) uuid.UUID {
// 	consumable := tp.Consumable{
// 		Title: fmt.Sprintf("Consumable %s", identifier),
// 	}
// 	consumable, err := store.CreateConsumable(consumable)
// 	if err != nil {
// 		t.Errorf("CreateConsumable() failed: %v", err)
// 	}
// 	return consumable.Id
// }

// func setupTool(t *testing.T, store st.Store, identifier string) uuid.UUID {
// 	tool := tp.Tool{
// 		Title: fmt.Sprintf("Tool %s", identifier),
// 		Size:  utest.ToPtr(fmt.Sprintf("Tool %s Size", identifier)),
// 	}
// 	tool, err := store.CreateTool(tool)
// 	if err != nil {
// 		t.Errorf("CreateTool() failed: %v", err)
// 	}
// 	return tool.Id
// }

// func setupTaskTemplate(t *testing.T, store st.Store, identifier string) uuid.UUID {
// 	task := tp.TaskTemplate{
// 		Title:       fmt.Sprintf("Task %s", identifier),
// 		Description: utest.ToPtr(fmt.Sprintf("Task %s description", identifier)),
// 		Type:        utest.ToPtr(tp.TaskTypePreventative),
// 	}
// 	task, err := store.CreateTaskTemplate(task)
// 	if err != nil {
// 		t.Errorf("CreateTaskTemplate() failed: %v", err)
// 	}
// 	return task.Id
// }

// func setupTask(t *testing.T, store st.Store, identifier string) uuid.UUID {
// 	assetId := setupAsset(t, store, identifier)
// 	taskId := setupTaskTemplate(t, store, identifier)
// 	assetTask := tp.Task{
// 		Title:          fmt.Sprintf("Task %s", identifier),
// 		Instructions:   utest.ToPtr(fmt.Sprintf("Task %s instructions", identifier)),
// 		AssetId:        assetId,
// 		TaskTemplateId: &taskId,
// 	}
// 	assetTask, err := store.CreateTask(assetTask)
// 	if err != nil {
// 		t.Errorf("CreateTask() failed: %v", err)
// 	}
// 	return assetTask.Id
// }
