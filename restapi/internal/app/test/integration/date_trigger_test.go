package integration

// func TestCreateDateTrigger(t *testing.T) {
// 	app := newApp("TestCreateDateTrigger")
// 	dateTrigger := tp.DateTrigger{Title: "Test DateTrigger"}
// 	createdDateTrigger, err := app.CreateDateTrigger("Test Group", "Test Asset", "Test Task", dateTrigger)
// 	assert.NoError(t, err)
// 	assert.Equal(t, dateTrigger.Title, createdDateTrigger.Title)
// }

// func TestDeleteDateTrigger(t *testing.T) {
// 	app := newApp("TestDeleteDateTrigger")
// 	dateTrigger := tp.DateTrigger{Title: "Test DateTrigger"}
// 	createdDateTrigger, _ := app.CreateDateTrigger("Test Group", "Test Asset", "Test Task", dateTrigger)
// 	err := app.DeleteDateTrigger("Test Group", "Test Asset", "Test Task", createdDateTrigger.Id.String())
// 	assert.NoError(t, err)
// }

// func TestListDateTriggers(t *testing.T) {
// 	app := newApp("TestListDateTriggers")
// 	dateTrigger := tp.DateTrigger{Title: "Test DateTrigger"}
// 	_, _ = app.CreateDateTrigger("Test Group", "Test Asset", "Test Task", dateTrigger)
// 	dateTriggers, err := app.ListDateTriggers("Test Group", "Test Asset", "Test Task")
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, dateTriggers)
// }

// func TestGetDateTrigger(t *testing.T) {
// 	app := newApp("TestGetDateTrigger")
// 	dateTrigger := tp.DateTrigger{Title: "Test DateTrigger"}
// 	createdDateTrigger, _ := app.CreateDateTrigger("Test Group", "Test Asset", "Test Task", dateTrigger)
// 	retrievedDateTrigger, err := app.GetDateTrigger("Test Group", "Test Asset", "Test Task", createdDateTrigger.Id.String())
// 	assert.NoError(t, err)
// 	assert.Equal(t, dateTrigger.Title, retrievedDateTrigger.Title)
// }

// func TestUpdateDateTrigger(t *testing.T) {
// 	app := newApp("TestUpdateDateTrigger")
// 	dateTrigger := tp.DateTrigger{Title: "Test DateTrigger"}
// 	createdDateTrigger, _ := app.CreateDateTrigger("Test Group", "Test Asset", "Test Task", dateTrigger)
// 	updatedDateTrigger := tp.DateTrigger{Title: "Updated DateTrigger"}
// 	_, err := app.UpdateDateTrigger("Test Group", "Test Asset", "Test Task", createdDateTrigger.Id.String(), updatedDateTrigger)
// 	assert.NoError(t, err)
// 	retrievedDateTrigger, _ := app.GetDateTrigger("Test Group", "Test Asset", "Test Task", createdDateTrigger.Id.String())
// 	assert.Equal(t, updatedDateTrigger.Title, retrievedDateTrigger.Title)
// }
