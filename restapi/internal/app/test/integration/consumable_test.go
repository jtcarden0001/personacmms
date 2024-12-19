package integration

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateConsumable(t *testing.T) {
	app := newApp("testcreateconsumable")
	consumable := tp.Consumable{Title: "Test Consumable"}
	createdConsumable, err := app.CreateConsumable(consumable)
	assert.NoError(t, err)
	assert.Equal(t, consumable.Title, createdConsumable.Title)
}

func TestDeleteConsumable(t *testing.T) {
	app := newApp("testdeleteconsumable")
	consumable := tp.Consumable{Title: "Test Consumable"}
	_, _ = app.CreateConsumable(consumable)
	err := app.DeleteConsumable(consumable.Title)
	assert.NoError(t, err)
}

func TestListConsumables(t *testing.T) {
	app := newApp("testlistconsumables")
	consumable := tp.Consumable{Title: "Test Consumable"}
	_, _ = app.CreateConsumable(consumable)
	consumables, err := app.ListConsumables()
	assert.NoError(t, err)
	assert.NotEmpty(t, consumables)
}

func TestGetConsumable(t *testing.T) {
	app := newApp("testgetconsumable")
	consumable := tp.Consumable{Title: "Test Consumable"}
	_, _ = app.CreateConsumable(consumable)
	retrievedConsumable, err := app.GetConsumable(consumable.Title)
	assert.NoError(t, err)
	assert.Equal(t, consumable.Title, retrievedConsumable.Title)
}

func TestUpdateConsumable(t *testing.T) {
	app := newApp("testupdateconsumable")
	consumable := tp.Consumable{Title: "Test Consumable"}
	_, _ = app.CreateConsumable(consumable)
	updatedConsumable := tp.Consumable{Title: "Updated Consumable"}
	_, err := app.UpdateConsumable(consumable.Title, updatedConsumable)
	assert.NoError(t, err)
	retrievedConsumable, _ := app.GetConsumable(updatedConsumable.Title)
	assert.Equal(t, updatedConsumable.Title, retrievedConsumable.Title)
}
