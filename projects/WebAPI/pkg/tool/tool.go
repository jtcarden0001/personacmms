package tool

import (
	imp "github.com/jtcarden0001/personacmms/projects/webapi/pkg/tool/postgres"
	tp "github.com/jtcarden0001/personacmms/projects/webapi/pkg/types"
)

func Create(title string) (int, error) {
	return imp.Create(title)
}

func Update(id int, title string) error {
	return imp.Update(id, title)
}

func Get(id int) (tp.Tool, error) {
	return imp.Get(id)
}

func GetAll() ([]tp.Tool, error) {
	return imp.GetAll()
}

func Delete(id int) error {
	return imp.Delete(id)
}

func ResetSequence(id int) error {
	return imp.ResetSequence(id)
}
