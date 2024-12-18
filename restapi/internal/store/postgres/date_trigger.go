package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var dateTriggerTableName = "datetrigger"

func (pg *Store) CreateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	dt.Id = uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id, date, task_id) VALUES ($1, $2, $3)", dateTriggerTableName)
	_, err := pg.db.Exec(query, dt.Id, dt.Date, dt.TaskId)
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}

	return dt, nil
}

func (pg *Store) DeleteDateTrigger(dtId uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", dateTriggerTableName)
	result, err := pg.db.Exec(query, dtId)
	if err != nil {
		return handleDbError(err, "date-trigger")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "date-trigger")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "date-trigger with id %s not found", dtId)
	}
	return nil
}

func (pg *Store) ListDateTriggers() ([]tp.DateTrigger, error) {
	query := fmt.Sprintf("SELECT id, date, task_id FROM %s", dateTriggerTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "date-trigger")
	}
	defer rows.Close()

	dts := []tp.DateTrigger{}
	for rows.Next() {
		var dt tp.DateTrigger
		if err := rows.Scan(&dt.Id, &dt.Date, &dt.TaskId); err != nil {
			return nil, handleDbError(err, "date-trigger")
		}
		dts = append(dts, dt)
	}

	return dts, nil
}

func (pg *Store) ListDateTriggersByTaskId(taskId uuid.UUID) ([]tp.DateTrigger, error) {
	query := fmt.Sprintf("SELECT id, date, task_id FROM %s WHERE task_id = $1", dateTriggerTableName)
	rows, err := pg.db.Query(query, taskId)
	if err != nil {
		return nil, handleDbError(err, "date-trigger")
	}

	dts := []tp.DateTrigger{}
	for rows.Next() {
		var dt tp.DateTrigger
		if err := rows.Scan(&dt.Id, &dt.Date, &dt.TaskId); err != nil {
			return nil, handleDbError(err, "date-trigger")
		}
		dts = append(dts, dt)
	}

	return dts, nil
}

func (pg *Store) GetDateTrigger(dtId uuid.UUID) (tp.DateTrigger, error) {
	query := fmt.Sprintf("SELECT id, date, task_id FROM %s WHERE id = $1", dateTriggerTableName)
	var dt tp.DateTrigger
	err := pg.db.QueryRow(query, dtId).Scan(&dt.Id, &dt.Date, &dt.TaskId)
	return dt, handleDbError(err, "date-trigger")
}

func (pg *Store) UpdateDateTrigger(dtId uuid.UUID, dt tp.DateTrigger) (tp.DateTrigger, error) {
	query := fmt.Sprintf("UPDATE %s SET date = $1, task_id = $2 WHERE id = $3 returning id", dateTriggerTableName)
	err := pg.db.QueryRow(query, dt.Date, dt.TaskId, dtId).Scan(&dt.Id)
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}

	return dt, nil
}
