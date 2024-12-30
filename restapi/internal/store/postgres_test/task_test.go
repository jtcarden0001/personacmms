package postgres_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestTaskCreate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtaskcreate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	aId := setupAssetInStore(t, store)
	tk := utest.SetupTask(1, aId, true)

	// test
	createdTask, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}

	utest.CompEntities(t, tk, createdTask)
}

func TestTaskDelete(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtaskdelete"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	aId := setupAssetInStore(t, store)
	tk := utest.SetupTask(1, aId, true)
	createdTask, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTaskDelete: failed during setup. CreateTask() failed: %v", err)
	}

	// test
	err = store.DeleteTask(createdTask.Id)
	if err != nil {
		t.Errorf("TestTaskDelete: DeleteTask() failed: %v", err)
	}

	_, err = store.GetTask(createdTask.Id)
	if err == nil {
		t.Errorf("TestTaskDelete: GetTask() returned nil error after deletion")
	}
}

func TestTaskGet(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtaskget"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	aId := setupAssetInStore(t, store)
	tk := utest.SetupTask(1, aId, true)
	createTask, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTaskGet: failed during setup. CreateTask() failed: %v", err)
	}

	// test
	getTask, err := store.GetTask(createTask.Id)
	if err != nil {
		t.Errorf("GetTask() failed: %v", err)
	}

	utest.CompEntities(t, createTask, getTask)
}

func TestTaskList(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtasklist"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	aId := setupAssetInStore(t, store)
	tk1 := utest.SetupTask(1, aId, true)
	tk2 := utest.SetupTask(2, aId, true)
	tk3 := utest.SetupTask(3, aId, true)

	_, err := store.CreateTask(tk1)
	if err != nil {
		t.Errorf("TestTaskList: failed during setup. CreateTask() failed: %v", err)
	}
	_, err = store.CreateTask(tk2)
	if err != nil {
		t.Errorf("TestTaskList: failed during setup. CreateTask() failed: %v", err)
	}
	_, err = store.CreateTask(tk3)
	if err != nil {
		t.Errorf("TestTaskList: failed during setup. CreateTask() failed: %v", err)
	}

	// test
	tasks, err := store.ListTasks()
	if err != nil {
		t.Errorf("ListTasks() failed: %v", err)
	}

	if len(tasks) != 3 {
		t.Errorf("ListTasks() failed: expected 3 tasks, got %d", len(tasks))
	}

	taskMap := map[uuid.UUID]tp.Task{
		tk1.Id: tk1,
		tk2.Id: tk2,
		tk3.Id: tk3,
	}

	for _, task := range tasks {
		expectedTask, ok := taskMap[task.Id]
		if !ok {
			t.Errorf("ListTasks() failed: unexpected task with ID %v", task.Id)
		}
		utest.CompEntities(t, expectedTask, task)
	}
}

func TestTaskUpdate(t *testing.T) {
	t.Parallel()

	// setup
	dbname := "testtaskupdate"
	store := utest.InitializeStore(dbname)
	defer utest.CloseStore(store, dbname)
	aId := setupAssetInStore(t, store)
	tk := utest.SetupTask(1, aId, true)
	createTask, err := store.CreateTask(tk)
	if err != nil {
		t.Errorf("TestTaskUpdate: failed during setup. CreateTask() failed: %v", err)
	}

	// test
	tk.Title = "Updated Task Title"
	tk.Instructions = utest.ToPtr("Updated Task Instructions")

	updatedTask, err := store.UpdateTask(tk)
	if err != nil {
		t.Errorf("UpdateTask() failed: %v", err)
	}

	differentFields := utest.ConvertStrArrToSet([]string{"Title", "Instructions"})
	utest.CompEntitiesFieldsShouldBeDifferent(t, createTask, updatedTask, differentFields)
}

func setupAssetInStore(t *testing.T, st store.Store) (assetId uuid.UUID) {
	a := utest.SetupAsset(1, true)
	_, err := st.CreateAsset(a)
	if err != nil {
		t.Fatalf("TestTaskUpdate: failed during setup. CreateAsset() failed: %v", err)
	}

	return a.Id
}
