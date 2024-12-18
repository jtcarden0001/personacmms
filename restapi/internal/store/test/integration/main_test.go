package integration

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// TODO: there is alot of postgres leaked from the store implementation. The clean approach
// will be to decouple the store test code from the implementation.  This is good enough for now.

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

func initializeStore(dbName string) st.Store {
	conninfo := fmt.Sprintf("user=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"))
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		log.Fatal(err)
	}

	store := st.NewWithDb(dbName)
	schema, err := os.ReadFile("../../postgres/ddl/init.sql")
	if err != nil {
		log.Fatalf("Could not read schema file: %s", err)
	}

	err = store.Exec(string(schema))
	if err != nil {
		log.Fatalf("Could not execute schema: %s", err)
	}

	return store
}

func closeStore(store st.Store, dbName string) {
	// cannnot drop the database because the connection is still open
	// store.Exec(fmt.Sprintf("DROP DATABASE %s", dbName))
	_ = dbName
	err := store.Close()
	if err != nil {
		log.Fatalf("Could not close store: %s", err)
	}
}

// TODO: some optimization and code reduction to be had here in these comparison functions
func compEntities(t *testing.T, expected interface{}, actual interface{}) {
	// exclude no fields
	compEntitiesExcludeFields(t, expected, actual, make(map[string]struct{}))
}

func compEntitiesExcludeFields(t *testing.T, expected interface{}, actual interface{}, fields map[string]struct{}) {
	// Compare all properties except for the specified fields
	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	for i := 0; i < expectedValue.NumField(); i++ {
		field := expectedValue.Type().Field(i)
		if _, ok := fields[field.Name]; ok {
			continue
		}

		expectedField := expectedValue.Field(i).Interface()
		actualField := actualValue.Field(i).Interface()

		isPtr := reflect.TypeOf(expectedField).Kind() == reflect.Ptr
		if isPtr {
			comparePointers(t, expectedField, actualField, field)
		} else {
			compareValues(t, expectedField, actualField, field)
		}
	}
}

func comparePointers(t *testing.T, expectedField interface{}, actualField interface{}, field reflect.StructField) {
	if expectedField == nil && actualField != nil {
		t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
	} else if expectedField != nil && actualField == nil {
		t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
	} else if expectedField != nil && actualField != nil {
		if !reflect.DeepEqual(expectedField, actualField) {
			t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
		}
	}
}

func compareValues(t *testing.T, expectedField interface{}, actualField interface{}, field reflect.StructField) {
	if field.Type == reflect.TypeOf(time.Time{}) {
		expectedTime := expectedField.(time.Time)
		actualTime := actualField.(time.Time)
		if !expectedTime.Truncate(time.Second).Equal(actualTime.Truncate(time.Second)) {
			t.Errorf("Compare failed: expected %v for field %s, got %v", expectedTime, field.Name, actualTime)
		}
	} else if !reflect.DeepEqual(expectedField, actualField) {
		t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
	}
}

func compEntitiesFieldsShouldBeDifferent(t *testing.T, inital interface{}, updated interface{}, fields map[string]struct{}) {
	// Compare all properties make sure fields are different
	initalValue := reflect.ValueOf(inital)
	updatedValue := reflect.ValueOf(updated)

	for i := 0; i < initalValue.NumField(); i++ {
		field := initalValue.Type().Field(i)
		different := false
		if _, ok := fields[field.Name]; ok {
			different = true
		}

		initalField := initalValue.Field(i).Interface()
		updatedField := updatedValue.Field(i).Interface()
		// Time based comparisons have drift that causes DeepEqual to fail for some reason
		if field.Name == "Date" {
			if !initalField.(time.Time).Equal(updatedField.(time.Time)) {
				t.Errorf("Compare failed: expected %v for field %s to be the same, got %v", initalField, field.Name, updatedField)
			}
		} else {
			compResult := reflect.DeepEqual(initalField, updatedField)
			if different && compResult {
				t.Errorf("Compare failed: expected %v for field %s to be different, got %v", initalField, field.Name, updatedField)
			} else if !different && !compResult {
				t.Errorf("Compare failed: expected %v for field %s to be the same, got %v", initalField, field.Name, updatedField)
			}
		}
	}
}

func setupGroup(t *testing.T, store st.Store, identifier string) string {
	group := tp.Group{
		Title: fmt.Sprintf("Group %s", identifier),
	}
	group, err := store.CreateGroup(group)
	if err != nil {
		t.Errorf("CreateGroup() failed: %v", err)
	}
	return group.Title
}

func setupCategory(t *testing.T, store st.Store, identifier string) string {
	category := tp.Category{
		Title:       fmt.Sprintf("Category %s", identifier),
		Description: utest.ToPtr(fmt.Sprintf("Category %s description", identifier)),
	}
	category, err := store.CreateCategory(category)
	if err != nil {
		t.Errorf("CreateCategory() failed: %v", err)
	}
	return category.Title
}

func setupConsumable(t *testing.T, store st.Store, identifier string) tp.UUID {
	consumable := tp.Consumable{
		Title: fmt.Sprintf("Consumable %s", identifier),
	}
	consumable, err := store.CreateConsumable(consumable)
	if err != nil {
		t.Errorf("CreateConsumable() failed: %v", err)
	}
	return consumable.Id
}

func setupTool(t *testing.T, store st.Store, identifier string) tp.UUID {
	tool := tp.Tool{
		Title: fmt.Sprintf("Tool %s", identifier),
		Size:  utest.ToPtr(fmt.Sprintf("Tool %s Size", identifier)),
	}
	tool, err := store.CreateTool(tool)
	if err != nil {
		t.Errorf("CreateTool() failed: %v", err)
	}
	return tool.Id
}

func setupAsset(t *testing.T, store st.Store, identifier string) tp.UUID {
	groupTitle := setupGroup(t, store, identifier)
	categoryTitle := setupCategory(t, store, identifier)
	asset := tp.Asset{
		Title:         fmt.Sprintf("Asset %s", identifier),
		Description:   utest.ToPtr(fmt.Sprintf("Asset %s description", identifier)),
		GroupTitle:    groupTitle,
		CategoryTitle: &categoryTitle,
	}
	asset, err := store.CreateAsset(asset)
	if err != nil {
		t.Errorf("CreateAsset() failed: %v", err)
	}
	return asset.Id
}

func setupTaskTemplate(t *testing.T, store st.Store, identifier string) tp.UUID {
	task := tp.TaskTemplate{
		Title:       fmt.Sprintf("Task %s", identifier),
		Description: utest.ToPtr(fmt.Sprintf("Task %s description", identifier)),
		Type:        utest.ToPtr(tp.TaskTypePreventative),
	}
	task, err := store.CreateTaskTemplate(task)
	if err != nil {
		t.Errorf("CreateTaskTemplate() failed: %v", err)
	}
	return task.Id
}

func setupTask(t *testing.T, store st.Store, identifier string) tp.UUID {
	assetId := setupAsset(t, store, identifier)
	taskId := setupTaskTemplate(t, store, identifier)
	assetTask := tp.Task{
		Title:          fmt.Sprintf("Task %s", identifier),
		Instructions:   utest.ToPtr(fmt.Sprintf("Task %s instructions", identifier)),
		AssetId:        assetId,
		TaskTemplateId: &taskId,
	}
	assetTask, err := store.CreateTask(assetTask)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}
	return assetTask.Id
}
