package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var dateTriggerTableName = "datetrigger"

func (pg *Store) CreateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	dt.Id = uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id, date, assettask_id) VALUES ($1, $2, $3)", dateTriggerTableName)
	_, err := pg.db.Exec(query, dt.Id, dt.Date, dt.AssetTaskId)
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err)
	}

	return dt, nil
}

func (pg *Store) DeleteDateTrigger(dtId uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", dateTriggerTableName)
	_, err := pg.db.Exec(query, dtId)
	return handleDbError(err)
}

func (pg *Store) ListDateTriggers() ([]tp.DateTrigger, error) {
	query := fmt.Sprintf("SELECT id, date, assettask_id FROM %s", dateTriggerTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err)
	}
	defer rows.Close()

	dts := []tp.DateTrigger{}
	for rows.Next() {
		var dt tp.DateTrigger
		if err := rows.Scan(&dt.Id, &dt.Date, &dt.AssetTaskId); err != nil {
			return nil, handleDbError(err)
		}
		dts = append(dts, dt)
	}

	return dts, nil
}

func (pg *Store) GetDateTrigger(dtId uuid.UUID) (tp.DateTrigger, error) {
	query := fmt.Sprintf("SELECT id, date, assettask_id FROM %s WHERE id = $1", dateTriggerTableName)
	var dt tp.DateTrigger
	err := pg.db.QueryRow(query, dtId).Scan(&dt.Id, &dt.Date, &dt.AssetTaskId)
	return dt, handleDbError(err)
}

func (pg *Store) UpdateDateTrigger(dtId uuid.UUID, dt tp.DateTrigger) (tp.DateTrigger, error) {
	query := fmt.Sprintf("UPDATE %s SET date = $1, assettask_id = $2 WHERE id = $3", dateTriggerTableName)
	_, err := pg.db.Exec(query, dt.Date, dt.AssetTaskId, dtId)
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err)
	}

	dt.Id = dtId
	return dt, nil
}
