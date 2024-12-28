package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	utest "github.com/jtcarden0001/personacmms/restapi/internal/utils/test"
)

func TestConsumableCreate(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupConsumable(1, false)
	creC, err := app.CreateConsumable(c)
	if err != nil {
		t.Fatalf("CreateConsumable() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, c, creC, diffFields)
}

func TestConsumableExists(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupConsumable(1, false)
	creCat, err := app.CreateConsumable(c)
	if err != nil {
		t.Fatalf("TestConsumableExistsOnAbsentConsumable: failed during setup. CreateConsumable() failed: %v", err)
	}

	exists, err := app.consumableExists(creCat.Id)
	if err != nil {
		t.Errorf("ConsumableExists() failed: %v", err)
	}

	if !exists {
		t.Errorf("ConsumableExists() returned false for existing consumable")
	}

	exists, err = app.consumableExists(creCat.Id)
	if err != nil {
		t.Errorf("ConsumableExists() failed: %v", err)
	}

	if !exists {
		t.Errorf("ConsumableExists() returned false for existing consumable")
	}
}

func TestConsumableDelete(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Fatalf("TestConsumableDelete: failed during setup. CreateConsumable() failed: %v", err)
	}

	err = app.DeleteConsumable(createdConsumable.Id.String())
	if err != nil {
		t.Fatalf("TestConsumableDelete: failed during delete. DeleteConsumable() failed: %v", err)
	}

	exists, err := app.consumableExists(createdConsumable.Id)
	if err != nil {
		t.Errorf("ConsumableExists() failed: %v", err)
	}

	if exists {
		t.Errorf("ConsumableExists() returned true for deleted consumable")
	}
}

func TestConsumableGet(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Fatalf("TestConsumableGet: failed during setup. CreateConsumable() failed: %v", err)
	}

	gotConsumable, err := app.GetConsumable(createdConsumable.Id.String())
	if err != nil {
		t.Fatalf("GetConsumable() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, createdConsumable, gotConsumable, diffFields)
}

func TestConsumableList(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Fatalf("TestConsumableList: failed during setup. CreateConsumable() failed: %v", err)
	}

	consumables, err := app.ListConsumables()
	if err != nil {
		t.Fatalf("ListConsumables() failed: %v", err)
	}

	if len(consumables) != 1 {
		t.Fatalf("ListConsumables() returned [%d] consumables, expected 1", len(consumables))
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, createdConsumable, consumables[0], diffFields)
}

func TestConsumableUpdate(t *testing.T) {
	app := &App{db: mock.New()}

	c := utest.SetupConsumable(1, false)
	createdConsumable, err := app.CreateConsumable(c)
	if err != nil {
		t.Fatalf("TestConsumableUpdate: failed during setup. CreateConsumable() failed: %v", err)
	}

	updatedConsumable := utest.SetupConsumable(1, false)
	updatedConsumable.Id = createdConsumable.Id
	updatedConsumable.Title = "Updated Title"

	updatedConsumable, err = app.UpdateConsumable(updatedConsumable.Id.String(), updatedConsumable)
	if err != nil {
		t.Fatalf("UpdateConsumable() failed: %v", err)
	}

	gotConsumable, err := app.GetConsumable(updatedConsumable.Id.String())
	if err != nil {
		t.Fatalf("GetConsumable() failed: %v", err)
	}

	diffFields := utest.ConvertStrArrToSet([]string{"Id"})
	utest.CompEntitiesExcludeFields(t, updatedConsumable, gotConsumable, diffFields)
}
