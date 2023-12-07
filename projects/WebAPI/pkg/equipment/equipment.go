package equipment

import (
	imp "github.com/jtcarden0001/personacmms/projects/webapi/pkg/equipment/postgres"
	tp "github.com/jtcarden0001/personacmms/projects/webapi/pkg/types"
)

func Create(title string, description string) (int, error) {
	return imp.Create(title, description)
}

func Update(id int, title string, description string) error {
	return imp.Update(id, title, description)
}

func Get(id int) (tp.Equipment, error) {
	return imp.Get(id)
}

func GetAll() ([]tp.Equipment, error) {
	return imp.GetAll()
}

func Delete(id int) error {
	return imp.Delete(id)
}

func ResetSequence(id int) error {
	return imp.ResetSequence(id)
}
