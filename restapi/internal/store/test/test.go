package test

import "github.com/jtcarden0001/personacmms/restapi/internal/store"

var testStore = store.NewTest()

func teardownTable(tableName string, lastId *int) error {
	err := testStore.CleanTable(tableName)
	if err != nil {
		return err
	}

	if lastId != nil {
		err = testStore.ResetSequence(tableName, *lastId)
	}

	return err
}
