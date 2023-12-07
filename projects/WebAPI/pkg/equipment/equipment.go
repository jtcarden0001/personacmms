package equipment

import (
	imp "github.com/jtcarden0001/personacmms/projects/webapi/pkg/equipment/postgres"
	tp "github.com/jtcarden0001/personacmms/projects/webapi/pkg/types"
)

// Create creates a new equipment record in the database
func Create(title string, description string) (int, error) {
	return imp.Create(title, description)
}

// Update updates an equipment record in the database
func Update(id int, title string, description string) error {
	return imp.Update(id, title, description)
}

// Get returns an equipment record from the database
func Get(id int) (tp.Equipment, error) {
	return imp.Get(id)
}

// GetAll returns all equipment records from the database
func GetAll() ([]tp.Equipment, error) {
	return imp.GetAll()
}

func Delete(id int) error {
	return imp.Delete(id)
}

func ResetSequence(id int) error {
	return imp.ResetSequence(id)
}
