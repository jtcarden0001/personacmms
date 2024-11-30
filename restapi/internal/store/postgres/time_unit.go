package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var timeUnitTableName = "timeunit"

func (pg *Store) CreateTimeUnit(tu tp.TimeUnit) (tp.TimeUnit, error) {
	//TODO: allow for group creation with a specified id ?
	id := uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2) returning id`, timeUnitTableName)
	_, err := pg.db.Exec(query, id, tu.Title)
	if err != nil {
		return tp.TimeUnit{}, handleDbError(err, "time-unit")
	}

	tu.Id = id
	return tu, nil
}

func (pg *Store) DeleteTimeUnit(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, timeUnitTableName)
	_, err := pg.db.Exec(query, title)

	return handleDbError(err, "time-unit")
}

func (pg *Store) ListTimeUnits() ([]tp.TimeUnit, error) {
	var timeUnits = []tp.TimeUnit{}
	query := fmt.Sprintf(`SELECT id, title FROM %s`, timeUnitTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return timeUnits, handleDbError(err, "time-unit")
	}
	defer rows.Close()

	for rows.Next() {
		var tu tp.TimeUnit
		err = rows.Scan(&tu.Id, &tu.Title)
		if err != nil {
			return nil, handleDbError(err, "time-unit")
		}
		timeUnits = append(timeUnits, tu)
	}

	return timeUnits, nil
}

func (pg *Store) GetTimeUnit(title string) (tp.TimeUnit, error) {
	var tu tp.TimeUnit
	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE title = $1`, timeUnitTableName)
	err := pg.db.QueryRow(query, title).Scan(&tu.Id, &tu.Title)
	if err != nil {
		return tp.TimeUnit{}, handleDbError(err, "time-unit")
	}

	return tu, nil
}

func (pg *Store) UpdateTimeUnit(title string, tu tp.TimeUnit) (tp.TimeUnit, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE title = $2`, timeUnitTableName)
	_, err := pg.db.Exec(query, tu.Title, title)
	if err != nil {
		return tp.TimeUnit{}, handleDbError(err, "time-unit")
	}

	return tu, nil
}
