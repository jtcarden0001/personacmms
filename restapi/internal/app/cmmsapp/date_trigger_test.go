package cmmsapp

// func setupGroupAssetTask(app *App) (tp.Group, tp.Asset, tp.Task) {
// 	db := app.db.(*mock.MockStore)
// 	group, _ := db.CreateGroup(tp.Group{Title: "group1"})
// 	asset, _ := db.CreateAsset(tp.Asset{Title: "asset1", GroupTitle: "group1"})
// 	task1, _ := app.CreateTask(group.Title, asset.Title, tp.Task{Title: "task1"})
// 	return group, asset, task1
// }

// func TestCreateDateTrigger(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	// setup group, asset, and task
// 	_, _, task := setupGroupAssetTask(app)

// 	dateTrigger := tp.DateTrigger{TaskId: task.Id}
// 	createdDateTrigger, err := app.CreateDateTrigger("group1", "asset1", task.Id.String(), dateTrigger)
// 	assert.NoError(t, err)
// 	assert.Equal(t, dateTrigger.TaskId, createdDateTrigger.TaskId)
// }

// func TestDeleteDateTrigger(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	// setup group, asset, and task
// 	_, _, task := setupGroupAssetTask(app)
// 	dateTrigger, _ := db.CreateDateTrigger(tp.DateTrigger{TaskId: task.Id})

// 	err := app.DeleteDateTrigger("group1", "asset1", task.Id.String(), dateTrigger.Id.String())
// 	assert.NoError(t, err)
// }

// func TestListDateTriggers(t *testing.T) {
// 	db := mock.New()
// 	app := &App{db: db}

// 	// setup group, asset, and task
// 	_, _, task := setupGroupAssetTask(app)
// 	db.CreateDateTrigger(tp.DateTrigger{TaskId: task.Id})
// 	db.CreateDateTrigger(tp.DateTrigger{TaskId: task.Id})

// 	dateTriggers, err := app.ListDateTriggers("group1", "asset1", task.Id.String())
// 	assert.NoError(t, err)
// 	assert.Len(t, dateTriggers, 2)
// }
