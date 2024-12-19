package integration

import (
	"testing"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateGroup(t *testing.T) {
	app := newApp("testcreategroup")
	group := tp.Group{Title: "Test Group"}
	createdGroup, err := app.CreateGroup(group)
	assert.NoError(t, err)
	assert.Equal(t, group.Title, createdGroup.Title)
}

func TestDeleteGroup(t *testing.T) {
	app := newApp("testdeletegroup")
	group := tp.Group{Title: "Test Group"}
	_, _ = app.CreateGroup(group)
	err := app.DeleteGroup(group.Title)
	assert.NoError(t, err)
}

func TestListGroups(t *testing.T) {
	app := newApp("testlistgroups")
	group := tp.Group{Title: "Test Group"}
	_, _ = app.CreateGroup(group)
	groups, err := app.ListGroups()
	assert.NoError(t, err)
	assert.NotEmpty(t, groups)
}

func TestGetGroup(t *testing.T) {
	app := newApp("testgetgroup")
	group := tp.Group{Title: "Test Group"}
	_, _ = app.CreateGroup(group)
	retrievedGroup, err := app.GetGroup(group.Title)
	assert.NoError(t, err)
	assert.Equal(t, group.Title, retrievedGroup.Title)
}

func TestUpdateGroup(t *testing.T) {
	app := newApp("testupdategroup")
	group := tp.Group{Title: "Test Group"}
	_, _ = app.CreateGroup(group)
	updatedGroup := tp.Group{Title: "Updated Group"}
	_, err := app.UpdateGroup(group.Title, updatedGroup)
	assert.NoError(t, err)
	retrievedGroup, _ := app.GetGroup(updatedGroup.Title)
	assert.Equal(t, updatedGroup.Title, retrievedGroup.Title)
}
