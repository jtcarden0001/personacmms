package cmmsapp

import (
	"testing"

	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateGroup(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Title: "group1"}
	createdGroup, err := app.CreateGroup(group)
	assert.NoError(t, err)
	assert.Equal(t, group.Title, createdGroup.Title)
}

func TestDeleteGroup(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Title: "group1"}
	db.CreateGroup(group)

	err := app.DeleteGroup("group1")
	assert.NoError(t, err)
}

func TestListGroups(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group1 := tp.Group{Title: "group1"}
	group2 := tp.Group{Title: "group2"}
	db.CreateGroup(group1)
	db.CreateGroup(group2)

	groups, err := app.ListGroups()
	assert.NoError(t, err)
	assert.Len(t, groups, 2)
}

func TestGetGroup(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Title: "group1"}
	db.CreateGroup(group)

	retrievedGroup, err := app.GetGroup("group1")
	assert.NoError(t, err)
	assert.Equal(t, group.Title, retrievedGroup.Title)
}

func TestUpdateGroup(t *testing.T) {
	db := mock.New()
	app := &App{db: db}

	group := tp.Group{Title: "group1"}
	db.CreateGroup(group)

	updatedGroup := tp.Group{Title: "group1_updated"}
	_, err := app.UpdateGroup("group1", updatedGroup)
	assert.NoError(t, err)

	retrievedGroup, err := app.GetGroup("group1_updated")
	assert.NoError(t, err)
	assert.Equal(t, updatedGroup.Title, retrievedGroup.Title)
}
