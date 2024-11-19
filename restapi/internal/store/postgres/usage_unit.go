package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type UsageUnit interface {
	CreateUsageUnit(tp.UsageUnit) (tp.UsageUnit, error)
	DeleteUsageUnit(string) error
	ListUsageUnits() ([]tp.UsageUnit, error)
	GetUsageUnit(string) (tp.UsageUnit, error)
	UpdateUsageUnit(string, tp.UsageUnit) (tp.UsageUnit, error)
}

var usageUnitTableName = "usageunit"

func (pg *Store) CreateUsageUnit(uu tp.UsageUnit) (tp.UsageUnit, error) {
	//TODO: allow for group creation with a specified id ?
	id := uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2) returning id`, usageUnitTableName)
	_, err := pg.db.Exec(query, id.String(), uu.Title)
	if err != nil {
		return tp.UsageUnit{}, processDbError(err)
	}

	uu.Id = id
	return uu, nil
}

func (pg *Store) DeleteUsageUnit(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, usageUnitTableName)
	_, err := pg.db.Exec(query, title)

	return processDbError(err)
}

func (pg *Store) ListUsageUnits() ([]tp.UsageUnit, error) {
	var usageUnits = []tp.UsageUnit{}
	query := fmt.Sprintf(`SELECT id, title FROM %s`, usageUnitTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return usageUnits, processDbError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var uu tp.UsageUnit
		err = rows.Scan(&uu.Id, &uu.Title)
		if err != nil {
			return nil, err
		}
		usageUnits = append(usageUnits, uu)
	}

	return usageUnits, nil
}

func (pg *Store) GetUsageUnit(title string) (tp.UsageUnit, error) {
	var uu tp.UsageUnit
	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE title = $1`, usageUnitTableName)
	err := pg.db.QueryRow(query, title).Scan(&uu.Id, &uu.Title)
	if err != nil {
		return tp.UsageUnit{}, processDbError(err)
	}

	return uu, nil
}

func (pg *Store) UpdateUsageUnit(title string, uu tp.UsageUnit) (tp.UsageUnit, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE title = $2`, usageUnitTableName)
	_, err := pg.db.Exec(query, uu.Title, title)
	if err != nil {
		return tp.UsageUnit{}, processDbError(err)
	}

	return uu, nil
}
