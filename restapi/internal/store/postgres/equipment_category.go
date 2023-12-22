package postgres

import (
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type EquipmentCategory interface {
	CreateEquipmentCategory(string) (int, error)
	DeleteEquipmentCategory(int) error
	GetAllEquipmentCategory() ([]tp.EquipmentCategory, error)
	GetEquipmentCategory(int) (tp.EquipmentCategory, error)
	UpdateEquipmentCategory(int, string) error
}

func (pg *Store) CreateEquipmentCategory(title string) (int, error) {
	query := `INSERT INTO equipment_category (title) VALUES ($1) returning id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteEquipmentCategory(id int) error {
	query := `DELETE FROM equipment_category WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllEquipmentCategory() ([]tp.EquipmentCategory, error) {
	query := `SELECT id, title FROM equipment_category`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []tp.EquipmentCategory
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
	query := `SELECT id, title FROM equipment_category WHERE id = $1`
	var ec tp.EquipmentCategory
	err := pg.db.QueryRow(query, id).Scan(&ec.Id, &ec.Title)

	return ec, err
}

func (pg *Store) UpdateEquipmentCategory(id int, title string) error {
	query := `UPDATE equipment_category SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, title, id)

	return err
}
