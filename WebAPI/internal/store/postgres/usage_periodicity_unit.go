package postgres

import (
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type UsagePeriodicityUnit interface {
	CreateUsagePeriodicityUnit(string) (int, error)
	DeleteUsagePeriodicityUnit(int) error
	GetAllUsagePeriodicityUnit() ([]tp.UsagePeriodicityUnit, error)
	GetUsagePeriodicityUnit(int) (tp.UsagePeriodicityUnit, error)
	UpdateUsagePeriodicityUnit(int, string) error
}

func (pg *Store) CreateUsagePeriodicityUnit(title string) (int, error) {
	query := `INSERT INTO usage_periodicity_unit (title) VALUES ($1) returning id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteUsagePeriodicityUnit(id int) error {
	query := `DELETE FROM usage_periodicity_unit WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllUsagePeriodicityUnit() ([]tp.UsagePeriodicityUnit, error) {
	query := `SELECT id, title FROM usage_periodicity_unit`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []tp.UsagePeriodicityUnit
	for rows.Next() {
		var upu tp.UsagePeriodicityUnit
		err = rows.Scan(&upu.Id, &upu.Title)
		if err != nil {
			return nil, err
		}
		units = append(units, upu)
	}

	return units, nil
}

func (pg *Store) GetUsagePeriodicityUnit(id int) (tp.UsagePeriodicityUnit, error) {
	query := `SELECT id, title FROM usage_periodicity_unit WHERE id = $1`
	var upu tp.UsagePeriodicityUnit
	err := pg.db.QueryRow(query, id).Scan(&upu.Id, &upu.Title)

	return upu, err
}

func (pg *Store) UpdateUsagePeriodicityUnit(id int, name string) error {
	query := `UPDATE usage_periodicity_unit SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, name, id)

	return err
}
