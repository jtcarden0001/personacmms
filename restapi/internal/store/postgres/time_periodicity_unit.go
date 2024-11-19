package postgres

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type TimeUnit interface {
	CreateTimeUnit(string) (int, error)
	DeleteTimeUnit(int) error
	GetAllTimeUnit() ([]tp.TimeUnit, error)
	GetTimeUnit(int) (tp.TimeUnit, error)
	UpdateTimeUnit(int, string) error
}

func (pg *Store) CreateTimeUnit(title string) (int, error) {
	query := `INSERT INTO time_periodicity_unit (title) VALUES ($1) returning id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteTimeUnit(id int) error {
	query := `DELETE FROM time_periodicity_unit WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllTimeUnit() ([]tp.TimeUnit, error) {
	query := `SELECT id, title FROM time_periodicity_unit`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var units []tp.TimeUnit
	for rows.Next() {
		var tpu tp.TimeUnit
		err = rows.Scan(&tpu.Id, &tpu.Title)
		if err != nil {
			return nil, err
		}
		units = append(units, tpu)
	}

	return units, nil
}

func (pg *Store) GetTimeUnit(id int) (tp.TimeUnit, error) {
	query := `SELECT id, title FROM time_periodicity_unit WHERE id = $1`
	var tpu tp.TimeUnit
	err := pg.db.QueryRow(query, id).Scan(&tpu.Id, &tpu.Title)

	return tpu, err
}

func (pg *Store) UpdateTimeUnit(id int, title string) error {
	query := `UPDATE time_periodicity_unit SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, title, id)

	return err
}
