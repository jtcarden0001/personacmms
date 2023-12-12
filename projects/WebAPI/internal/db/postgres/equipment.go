package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
)

func (pg *Store) CreateEquipment(title string, description string) (int, error) {
	// TODO: add validation to prevent sql injection
	query := `INSERT INTO Equipment (title, description) VALUES ($1, $2) returning id`
	var id int
	err := pg.db.QueryRow(query, title, description).Scan(&id)

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
	query := `SELECT id, title, description FROM Equipment`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e tp.Equipment
		err = rows.Scan(&e.Id, &e.Title, &e.Description)
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
	query := `SELECT id, title, description FROM Equipment WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&e.Id, &e.Title, &e.Description)
	return e, err
}

func (pg *Store) UpdateEquipment(id int, title string, description string) error {
	// TODO: add validation to prevent sql injection
	query := `UPDATE Equipment SET title = $1, description = $2 WHERE id = $3`
	_, err := pg.db.Exec(query, title, description, id)

	return err
}

func (pg *Store) ResetSequenceEquipment(id int) error {
	query := fmt.Sprintf("ALTER SEQUENCE equipment_id_seq RESTART WITH %d", id)
	_, err := pg.db.Query(query)
	return err
}
