package postgres

import (
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type TimePeriodicityUnit interface {
	CreateTimePeriodicityUnit(string) (int, error)
	DeleteTimePeriodicityUnit(int) error
	GetAllTimePeriodicityUnit() ([]tp.TimePeriodicityUnit, error)
	GetTimePeriodicityUnit(int) (tp.TimePeriodicityUnit, error)
	UpdateTimePeriodicityUnit(int, string) error
}

func (pg *Store) CreateTimePeriodicityUnit(title string) (int, error) {
	query := `INSERT INTO time_periodicity_unit (title) VALUES ($1) returning id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteTimePeriodicityUnit(id int) error {
	query := `DELETE FROM time_periodicity_unit WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllTimePeriodicityUnit() ([]tp.TimePeriodicityUnit, error) {
	query := `SELECT id, title FROM time_periodicity_unit`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []tp.TimePeriodicityUnit
	for rows.Next() {
		var tpu tp.TimePeriodicityUnit
		err = rows.Scan(&tpu.Id, &tpu.Title)
		if err != nil {
			return nil, err
		}
		units = append(units, tpu)
	}

	return units, nil
}

func (pg *Store) GetTimePeriodicityUnit(id int) (tp.TimePeriodicityUnit, error) {
	query := `SELECT id, title FROM time_periodicity_unit WHERE id = $1`
	var tpu tp.TimePeriodicityUnit
	err := pg.db.QueryRow(query, id).Scan(&tpu.Id, &tpu.Title)

	return tpu, err
}

func (pg *Store) UpdateTimePeriodicityUnit(id int, title string) error {
	query := `UPDATE time_periodicity_unit SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, title, id)

	return err
}
