package cmmsapp

// import (
// 	"testing"

// 	"github.com/jtcarden0001/personacmms/restapi/internal/store/mock"
// 	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateTool(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	tool := tp.Tool{Title: "Hammer"}
// 	_, err := app.CreateTool(tool)
// 	assert.NoError(t, err)
// }

// func TestDeleteTool(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	tool := tp.Tool{Title: "Hammer"}
// 	tool, _ = db.CreateTool(tool)

// 	err := app.DeleteTool(tool.Title)
// 	assert.NoError(t, err)
// }

// func TestListTools(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	tool1 := tp.Tool{Title: "Hammer"}
// 	tool2 := tp.Tool{Title: "Wrench"}
// 	db.CreateTool(tool1)
// 	db.CreateTool(tool2)

// 	tools, err := app.ListTools()
// 	assert.NoError(t, err)
// 	assert.Len(t, tools, 2)
// }

// func TestGetTool(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	tool := tp.Tool{Title: "Hammer"}
// 	tool, _ = db.CreateTool(tool)

// 	_, err := app.GetTool(tool.Title)
// 	assert.NoError(t, err)
// }

// func TestUpdateTool(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	tool := tp.Tool{Title: "Hammer"}
// 	tool, _ = db.CreateTool(tool)

// 	tool.Title = "Wrench"
// 	_, err := app.UpdateTool("Hammer", tool)
// 	assert.NoError(t, err)
// }
