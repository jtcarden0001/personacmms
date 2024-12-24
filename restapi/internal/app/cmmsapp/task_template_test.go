package cmmsapp

// func TestCreateTaskTemplate(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	taskTemplate := tp.TaskTemplate{Title: "template1"}
// 	createdTaskTemplate, err := app.CreateTaskTemplate(taskTemplate)
// 	assert.NoError(t, err)
// 	assert.Equal(t, taskTemplate.Title, createdTaskTemplate.Title)
// }

// func TestDeleteTaskTemplate(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	taskTemplate := tp.TaskTemplate{Title: "template1"}
// 	db.CreateTaskTemplate(taskTemplate)

// 	err := app.DeleteTaskTemplate("template1")
// 	assert.NoError(t, err)
// }

// func TestListTaskTemplates(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	taskTemplate1 := tp.TaskTemplate{Title: "template1"}
// 	taskTemplate2 := tp.TaskTemplate{Title: "template2"}
// 	db.CreateTaskTemplate(taskTemplate1)
// 	db.CreateTaskTemplate(taskTemplate2)

// 	taskTemplates, err := app.ListTaskTemplates()
// 	assert.NoError(t, err)
// 	assert.Len(t, taskTemplates, 2)
// }

// func TestGetTaskTemplate(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	taskTemplate := tp.TaskTemplate{Title: "template1"}
// 	db.CreateTaskTemplate(taskTemplate)

// 	retrievedTaskTemplate, err := app.GetTaskTemplate("template1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, taskTemplate.Title, retrievedTaskTemplate.Title)
// }

// func TestUpdateTaskTemplate(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	taskTemplate := tp.TaskTemplate{Title: "template1"}
// 	db.CreateTaskTemplate(taskTemplate)

// 	updatedTaskTemplate := tp.TaskTemplate{Title: "template1_updated"}
// 	_, err := app.UpdateTaskTemplate("template1", updatedTaskTemplate)
// 	assert.NoError(t, err)

// 	retrievedTaskTemplate, err := app.GetTaskTemplate("template1_updated")
// 	assert.NoError(t, err)
// 	assert.Equal(t, updatedTaskTemplate.Title, retrievedTaskTemplate.Title)
// }
