package postgres

import (
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type Consumable interface {
	CreateConsumable(string) (int, error)
	DeleteConsumable(int) error
	GetAllConsumable() ([]tp.Consumable, error)
	GetConsumable(int) (tp.Consumable, error)
	UpdateConsumable(int, string) error
}

func (pg *Store) CreateConsumable(title string) (int, error) {
	query := "INSERT INTO consumable (title) VALUES ($1) RETURNING id"
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteConsumable(id int) error {
	query := "DELETE FROM consumable WHERE id = $1"
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllConsumable() ([]tp.Consumable, error) {
	var consumables []tp.Consumable
	query := "SELECT id, title FROM consumable"
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c tp.Consumable
		err = rows.Scan(&c.Id, &c.Title)
		if err != nil {
			return nil, err
		}
		consumables = append(consumables, c)
	}

	return consumables, err
}

func (pg *Store) GetConsumable(id int) (tp.Consumable, error) {
	query := "SELECT id, title FROM consumable WHERE id = $1"
	var c tp.Consumable
	err := pg.db.QueryRow(query, id).Scan(&c.Id, &c.Title)

	return c, err
}

func (pg *Store) UpdateConsumable(id int, title string) error {
	query := "UPDATE consumable SET title = $1 WHERE id = $2"
	_, err := pg.db.Exec(query, title, id)

	return err
}
