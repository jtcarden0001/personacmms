package postgres

import (
	"errors"

	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type Consumable interface {
	CreateConsumable(string) (int, error)
	DeleteConsumable(int) error
	GetAllConsumable() ([]tp.EquipmentCategory, error)
	GetConsumable(int) (tp.EquipmentCategory, error)
	UpdateConsumable(int, string) error
}

type ConsumableTest interface {
	ResetSequenceConsumable(int) error
}

func (pg *Store) CreateConsumable(title string) (int, error) {
	query := "INSERT INTO consumable (title) VALUES ($1) RETURNING id"
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteConsumable(id int) error {
	return errors.New("not implemented")
}

func (pg *Store) GetAllConsumable() ([]tp.EquipmentCategory, error) {
	return nil, errors.New("not implemented")
}

func (pg *Store) GetConsumable(id int) (tp.EquipmentCategory, error) {
	return tp.EquipmentCategory{}, errors.New("not implemented")
}

func (pg *Store) UpdateConsumable(id int, title string) error {
	return errors.New("not implemented")
}

func (pg *Store) ResetSequenceConsumable(id int) error {
	return errors.New("not implemented")
}
