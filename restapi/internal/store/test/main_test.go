package test

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	st "github.com/jtcarden0001/personacmms/restapi/internal/store"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	log "github.com/sirupsen/logrus"
)

// TODO: there is alot of postgres leaked from the store implementation. The clean approach
// will be to decouple the store test code from the implementation.  This is good enough for now.

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	os.Setenv("DATABASE_USER", "test")
	os.Setenv("DATABASE_PASSWORD", "test")
	os.Setenv("DATABASE_NAME", "test")
	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.13-alpine3.20",
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", os.Getenv("DATABASE_PASSWORD")),
			fmt.Sprintf("POSTGRES_USER=%s", os.Getenv("DATABASE_USER")),
			fmt.Sprintf("POSTGRES_DB=%s", os.Getenv("DATABASE_NAME")),
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	defer resource.Close()

	hostAndPort := resource.GetHostPort("5432/tcp")
	host, port, err := net.SplitHostPort(hostAndPort)
	if err != nil {
		log.Fatalf("Invalid host:port format: %s", err)
	}

	os.Setenv("DATABASE_HOST", host)
	os.Setenv("DATABASE_PORT", port)
	databaseUrl := fmt.Sprintf("postgres://test:test@%s/test?sslmode=disable", hostAndPort)
	log.Println("Connecting to database on url: ", databaseUrl)
	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	var db *sql.DB
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
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
	schema, err := os.ReadFile("../postgres/ddl/init.sql")
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

func convertToSet(arr []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	return set
}

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

		if strings.Contains(field.Name, "Date") {
			expectedTime, ok1 := expectedField.(time.Time)
			actualTime, ok2 := actualField.(time.Time)
			if !ok1 || !ok2 {
				expectedTimePtr, ok1 := expectedField.(*time.Time)
				actualTimePtr, ok2 := actualField.(*time.Time)
				if ok1 && ok2 && expectedTimePtr != nil && actualTimePtr != nil {
					expectedTime = *expectedTimePtr
					actualTime = *actualTimePtr
				} else {
					t.Errorf("Compare failed: expected %v for field %s, got %v", expectedField, field.Name, actualField)
					continue
				}
			}
			if !expectedTime.Truncate(time.Second).Equal(actualTime.Truncate(time.Second)) {
				t.Errorf("Compare failed: expected %v for field %s, got %v", expectedTime, field.Name, actualTime)
			}
		} else {
			expectedValue := reflect.Indirect(reflect.ValueOf(expectedField)).Interface()
			actualValue := reflect.Indirect(reflect.ValueOf(actualField)).Interface()
			if !reflect.DeepEqual(expectedValue, actualValue) {
				t.Errorf("Create() failed: expected %v for field %s, got %v", expectedValue, field.Name, actualValue)
			}
		}
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
		Description: fmt.Sprintf("Category %s description", identifier),
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
		Size:  fmt.Sprintf("Tool %s Size", identifier),
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
		Description:   fmt.Sprintf("Asset %s description", identifier),
		GroupTitle:    groupTitle,
		CategoryTitle: categoryTitle,
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
		Description: fmt.Sprintf("Task %s description", identifier),
		Type:        tp.TaskTypePreventative,
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
		Title:              fmt.Sprintf("Task %s", identifier),
		UniqueInstructions: fmt.Sprintf("Task %s instructions", identifier),
		AssetId:            assetId,
		TaskTemplateId:     taskId,
	}
	assetTask, err := store.CreateTask(assetTask)
	if err != nil {
		t.Errorf("CreateTask() failed: %v", err)
	}
	return assetTask.Id
}
