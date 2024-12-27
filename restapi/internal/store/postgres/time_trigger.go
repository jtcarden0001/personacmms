package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var timeTriggerTableName = "timetrigger"

func (pg *PostgresStore) CreateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	tuid, err := pg.getTimeUnitIdFromTitle(tt.TimeUnit)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "failed get the db id for time unitd '%s' during TimeTrigger create", tt.TimeUnit)
	}

	query := fmt.Sprintf(`INSERT INTO %s (id, quantity, timeunit_id, task_id) VALUES ($1, $2, $3, $4)`, timeTriggerTableName)
	_, err = pg.db.Exec(query, tt.Id, tt.Quantity, tuid, tt.TaskId)
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	return tt, nil
}

func (pg *PostgresStore) DeleteTimeTrigger(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, timeTriggerTableName)
	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "time-trigger")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "time-trigger")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "time trigger with id '%s' not found", id)
	}

	return nil
}

func (pg *PostgresStore) GetTimeTrigger(id uuid.UUID) (tp.TimeTrigger, error) {
	query := fmt.Sprintf(`SELECT id, quantity, timeunit_id, task_id FROM %s WHERE id=$1`, timeTriggerTableName)
	var tt tp.TimeTrigger
	var tuid uuid.UUID
	err := pg.db.QueryRow(query, id).Scan(&tt.Id, &tt.Quantity, &tuid, &tt.TaskId)
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	tuTitle, err := pg.getTimeUnitTitleFromId(tuid)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "failed to get time unit title for id '%s' during TimeTrigger get", tuid)
	}

	tt.TimeUnit = tuTitle

	return tt, nil
}

func (pg *PostgresStore) ListTimeTriggers() ([]tp.TimeTrigger, error) {
	query := fmt.Sprintf(`SELECT id, quantity, timeunit_id, task_id FROM %s`, timeTriggerTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return []tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}
	defer rows.Close()

	var ttgs []tp.TimeTrigger
	for rows.Next() {
		var tt tp.TimeTrigger
		var tuid uuid.UUID
		err := rows.Scan(&tt.Id, &tt.Quantity, &tuid, &tt.TaskId)
		if err != nil {
			return []tp.TimeTrigger{}, handleDbError(err, "time-trigger")
		}

		tuTitle, err := pg.getTimeUnitTitleFromId(tuid)
		if err != nil {
			return []tp.TimeTrigger{}, errors.Wrapf(err, "failed to get time unit title for id '%s' during TimeTrigger list", tuid)
		}

		tt.TimeUnit = tuTitle
		ttgs = append(ttgs, tt)
	}

	return ttgs, nil
}

func (pg *PostgresStore) UpdateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	tuid, err := pg.getTimeUnitIdFromTitle(tt.TimeUnit)
	if err != nil {
		return tp.TimeTrigger{}, errors.Wrapf(err, "failed get the db id for time unitd '%s' during TimeTrigger update", tt.TimeUnit)
	}

	query := fmt.Sprintf(`UPDATE %s SET quantity=$1, timeunit_id=$2, task_id=$3 WHERE id=$4`, timeTriggerTableName)
	result, err := pg.db.Exec(query, tt.Quantity, tuid, tt.TaskId, tt.Id)
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	if rowsAffected == 0 {
		return tp.TimeTrigger{}, errors.Wrapf(ae.ErrNotFound, "time trigger with id '%s' not found", tt.Id)
	}

	return tt, nil
}

var timeUnitTableName = "timeunit"

func (pg *PostgresStore) getTimeUnitIdFromTitle(title string) (uuid.UUID, error) {
	query := fmt.Sprintf(`SELECT id FROM %s WHERE title=$1`, timeUnitTableName)
	var id uuid.UUID
	err := pg.db.QueryRow(query, title).Scan(&id)
	if err != nil {
		return uuid.Nil, handleDbError(err, "timeunit")
	}

	return id, nil
}

func (pg *PostgresStore) getTimeUnitTitleFromId(id uuid.UUID) (string, error) {
	query := fmt.Sprintf(`SELECT title FROM %s WHERE id=$1`, timeUnitTableName)
	var title string
	err := pg.db.QueryRow(query, id).Scan(&title)
	if err != nil {
		return "", handleDbError(err, "timeunit")
	}

	return title, nil
}
