package postgres

import (
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type Equipment interface {
	CreateEquipment(string, int, string, string, string, int) (int, error)
	DeleteEquipment(int) error
	GetAllEquipment() ([]tp.Equipment, error)
	GetEquipment(int) (tp.Equipment, error)
	UpdateEquipment(int, string, int, string, string, string, int) error
}

func (pg *Store) CreateEquipment(title string, year int, make, modelNumber, description string, categoryId int) (int, error) {
	query := `INSERT INTO Equipment (title, year, make, model_number, description, category_id) VALUES ($1, $2, $3, $4, $5, $6) returning id`
	var id int
	err := pg.db.QueryRow(query, title, year, make, modelNumber, description, categoryId).Scan(&id)

	return id, err
}

func (pg *Store) DeleteEquipment(id int) error {
	query := `DELETE FROM Equipment WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllEquipment() ([]tp.Equipment, error) {
	var Equipment []tp.Equipment
	query := `SELECT id, title, year, make, model_number, description, category_id FROM Equipment`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e tp.Equipment
		err = rows.Scan(&e.Id, &e.Title, &e.Year, &e.Make, &e.ModelNumber, &e.Description, &e.CategoryId)
		if err != nil {
			return nil, err
		}
		Equipment = append(Equipment, e)
	}

	return Equipment, nil
}

func (pg *Store) GetEquipment(id int) (tp.Equipment, error) {
	var e tp.Equipment
	query := `SELECT id, title, year, make, model_number, description, category_id FROM Equipment WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&e.Id, &e.Title, &e.Year, &e.Make, &e.ModelNumber, &e.Description, &e.CategoryId)

	return e, err
}

func (pg *Store) UpdateEquipment(id int, title string, year int, make, modelNumber, description string, categoryId int) error {
	query := `UPDATE Equipment SET title = $1, year = $2, make = $3, model_number = $4, description = $6, category = $7 WHERE id = $8`
	_, err := pg.db.Exec(query, title, year, make, modelNumber, description, categoryId, id)

	return err
}
