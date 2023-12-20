package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type EquipmentCategory interface {
	CreateEquipmentCategory(string) (int, error)
	DeleteEquipmentCategory(int) error
	GetAllEquipmentCategory() ([]tp.EquipmentCategory, error)
	GetEquipmentCategory(int) (tp.EquipmentCategory, error)
	UpdateEquipmentCategory(int, string) error
}

type EquipmentCategoryTest interface {
	ResetSequenceEquipmentCategory(int) error
}

func (pg *Store) CreateEquipmentCategory(title string) (int, error) {
	// TODO: add validation to prevent sql injection
	query := `INSERT INTO equipment_category (title) VALUES ($1) returning id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteEquipmentCategory(id int) error {
	// TODO: add validation to prevent sql injection
	query := `DELETE FROM equipment_category WHERE id = $1`
	_, err := pg.db.Exec(query, id)
	return err
}

func (pg *Store) GetAllEquipmentCategory() ([]tp.EquipmentCategory, error) {
	// TODO: add validation to prevent sql injection
	var categories []tp.EquipmentCategory
	query := `SELECT id, title FROM equipment_category`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ec tp.EquipmentCategory
		err = rows.Scan(&ec.Id, &ec.Title)
		if err != nil {
			return nil, err
		}
		categories = append(categories, ec)
	}

	return categories, nil
}

func (pg *Store) GetEquipmentCategory(id int) (tp.EquipmentCategory, error) {
	// TODO: add validation to prevent sql injection
	var ec tp.EquipmentCategory
	query := `SELECT id, title FROM equipment_category WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&ec.Id, &ec.Title)

	return ec, err
}

func (pg *Store) UpdateEquipmentCategory(id int, title string) error {
	// TODO: add validation to prevent sql injection
	query := `UPDATE equipment_category SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, title, id)

	return err
}

func (pg *Store) ResetSequenceEquipmentCategory(id int) error {
	query := fmt.Sprintf("ALTER SEQUENCE equipment_category_id_seq RESTART WITH %d", id)
	_, err := pg.db.Query(query)
	return err
}
