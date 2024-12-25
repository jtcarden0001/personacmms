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
	tt.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, quantity, timeunit, task_id) VALUES ($1, $2, $3, $4)`, timeTriggerTableName)
	_, err := pg.db.Exec(query, tt.Id, tt.Quantity, tt.TimeUnit, tt.TaskId)
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
	var tt tp.TimeTrigger
	query := fmt.Sprintf(`SELECT id, quantity, timeunit_title, task_id FROM %s WHERE id=$1`, timeTriggerTableName)
	err := pg.db.QueryRow(query, id).Scan(&tt.Id, &tt.Quantity, &tt.TimeUnit, &tt.TaskId)
	if err != nil {
		return tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}

	return tt, nil
}

func (pg *PostgresStore) ListTimeTriggers() ([]tp.TimeTrigger, error) {
	query := fmt.Sprintf(`SELECT id, quantity, timeunit_title, task_id FROM %s`, timeTriggerTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return []tp.TimeTrigger{}, handleDbError(err, "time-trigger")
	}
	defer rows.Close()

	var ttgs []tp.TimeTrigger
	for rows.Next() {
		var tt tp.TimeTrigger
		err := rows.Scan(&tt.Id, &tt.Quantity, &tt.TimeUnit, &tt.TaskId)
		if err != nil {
			return []tp.TimeTrigger{}, handleDbError(err, "time-trigger")
		}
		ttgs = append(ttgs, tt)
	}

	return ttgs, nil
}

func (pg *PostgresStore) UpdateTimeTrigger(tt tp.TimeTrigger) (tp.TimeTrigger, error) {
	query := fmt.Sprintf(`UPDATE %s SET quantity=$1, timeunit_title=$2, task_id=$3 WHERE id=$4`, timeTriggerTableName)
	result, err := pg.db.Exec(query, tt.Quantity, tt.TimeUnit, tt.TaskId, tt.Id)
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
