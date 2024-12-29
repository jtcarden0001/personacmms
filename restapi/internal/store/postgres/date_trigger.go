package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var dateTriggerTableName = "datetrigger"

func (pg *PostgresStore) CreateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (id, scheduled_date, task_id) 
			VALUES ($1, $2, $3)`,
		dateTriggerTableName)

	_, err := pg.db.Exec(query, dt.Id, dt.ScheduledDate, dt.TaskId)
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}

	return dt, nil
}

func (pg *PostgresStore) DeleteDateTrigger(id uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = $1`,
		dateTriggerTableName)

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "date-trigger")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "date-trigger")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "date-trigger with id %s not found", id)
	}
	return nil
}

func (pg *PostgresStore) DeleteDateTriggerFromTask(taskId uuid.UUID, dtId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE task_id = $1 AND id = $2`,
		dateTriggerTableName)

	result, err := pg.db.Exec(query, taskId, dtId)
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

func (pg *PostgresStore) GetDateTrigger(id uuid.UUID) (tp.DateTrigger, error) {
	query := fmt.Sprintf(`
			SELECT id, scheduled_date, task_id 
			FROM %s 
			WHERE id = $1`,
		dateTriggerTableName)

	var dt tp.DateTrigger
	err := pg.db.QueryRow(query, id).Scan(&dt.Id, &dt.ScheduledDate, &dt.TaskId)
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}

	return dt, nil
}

func (pg *PostgresStore) ListDateTriggers() ([]tp.DateTrigger, error) {
	query := fmt.Sprintf(`
			SELECT id, scheduled_date, task_id 
			FROM %s`,
		dateTriggerTableName)

	rows, err := pg.db.Query(query)
	if err != nil {
		return []tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}
	defer rows.Close()

	dts := []tp.DateTrigger{}
	for rows.Next() {
		var dt tp.DateTrigger
		if err := rows.Scan(&dt.Id, &dt.ScheduledDate, &dt.TaskId); err != nil {
			return []tp.DateTrigger{}, handleDbError(err, "date-trigger")
		}
		dts = append(dts, dt)
	}

	return dts, nil
}

func (pg *PostgresStore) ListDateTriggersByTask(taskId uuid.UUID) ([]tp.DateTrigger, error) {
	query := fmt.Sprintf(`
			SELECT id, scheduled_date, task_id
			FROM %s
			WHERE task_id = $1`,
		dateTriggerTableName)

	rows, err := pg.db.Query(query, taskId)
	if err != nil {
		return []tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}
	defer rows.Close()

	dts := []tp.DateTrigger{}
	for rows.Next() {
		var dt tp.DateTrigger
		if err := rows.Scan(&dt.Id, &dt.ScheduledDate, &dt.TaskId); err != nil {
			return []tp.DateTrigger{}, handleDbError(err, "date-trigger")
		}
		dts = append(dts, dt)
	}

	return dts, nil
}

func (pg *PostgresStore) UpdateDateTrigger(dt tp.DateTrigger) (tp.DateTrigger, error) {
	query := fmt.Sprintf(`
			UPDATE %s 
			SET scheduled_date = $1, task_id = $2 
			WHERE id = $3`,
		dateTriggerTableName)

	result, err := pg.db.Exec(query, dt.ScheduledDate, dt.TaskId, dt.Id)
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.DateTrigger{}, handleDbError(err, "date-trigger")
	}

	if rowsAffected == 0 {
		return tp.DateTrigger{}, errors.Wrapf(ae.ErrNotFound, "date-trigger with id %s not found", dt.Id)
	}

	return dt, nil
}
