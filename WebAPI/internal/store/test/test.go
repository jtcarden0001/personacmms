package test

import "github.com/jtcarden0001/personacmms/webapi/internal/store"

var testStore = store.NewTest()

func initStore() error {
	err := testStore.CleanTestStore()
	if err != nil {
		return err
	}

	return nil
}
