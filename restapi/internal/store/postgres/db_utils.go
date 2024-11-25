package postgres

func processDbError(err error) error {
	if err != nil {
		return err
	}

	// TODO: process different errors from postgres (i.e. validation, not found, etc.)
	return err
}
