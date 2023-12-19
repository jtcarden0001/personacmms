package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
)

func (pg *Store) CreateEquipment(title string, year int, make, modelNumber, description string, categoryId int) (int, error) {
	// TODO: add validation to prevent sql injection
	query := `INSERT INTO Equipment (title, year, make, model_number, description, category_id) VALUES ($1, $2, $3, $4, $5, $6) returning id`
	var id int
	err := pg.db.QueryRow(query, title, year, make, modelNumber, description, categoryId).Scan(&id)

	return id, err
}

func (pg *Store) DeleteEquipment(id int) error {
	// TODO: add validation to prevent sql injection
	query := `DELETE FROM Equipment WHERE id = $1`
	_, err := pg.db.Exec(query, id)
	return err
}

func (pg *Store) GetAllEquipment() ([]tp.Equipment, error) {
	// TODO: add validation to prevent sql injection
	var Equipment []tp.Equipment
	query := `SELECT id, title, year, make, model_number, description, category_id FROM Equipment`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}

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
	// TODO: add validation to prevent sql injection
	var e tp.Equipment
	query := `SELECT id, title, year, make, model_number, description, category_id FROM Equipment WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&e.Id, &e.Title, &e.Year, &e.Make, &e.ModelNumber, &e.Description, &e.CategoryId)
	return e, err
}

func (pg *Store) UpdateEquipment(id int, title string, year int, make, modelNumber, description string, categoryId int) error {
	// TODO: add validation to prevent sql injection
	query := `UPDATE Equipment SET title = $1, year = $2, make = $3, model_number = $4, description = $6, category = $7 WHERE id = $8`
	_, err := pg.db.Exec(query, title, year, make, modelNumber, description, categoryId, id)

	return err
}

// TODO: maybe make a separate create function that takes a category ID.
func (pg *Store) UpdateEquipmentCategoryFK(equipmentId, categoryId int) error {
	// TODO: add validation to prevent sql injection
	query := `UPDATE Equipment SET category_id = $1 WHERE id = $6`
	_, err := pg.db.Exec(query, categoryId, equipmentId)

	return err
}

func (pg *Store) ResetSequenceEquipment(id int) error {
	query := fmt.Sprintf("ALTER SEQUENCE equipment_id_seq RESTART WITH %d", id)
	_, err := pg.db.Query(query)
	return err
}
