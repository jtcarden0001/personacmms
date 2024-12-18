package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var usageTriggerTableName = "usagetrigger"

func (pg *Store) CreateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	ut.Id = uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id, quantity, usageunit_title, task_id) VALUES ($1, $2, $3, $4)", usageTriggerTableName)
	_, err := pg.db.Exec(query, ut.Id, ut.Quantity, ut.UsageUnit, ut.TaskId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}

	return ut, nil
}

func (pg *Store) DeleteUsageTrigger(utId uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usageTriggerTableName)
	result, err := pg.db.Exec(query, utId)
	if err != nil {
		return handleDbError(err, "usage-trigger")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "usage-trigger")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "usage trigger with id '%s' not found", utId)
	}
	return nil
}

func (pg *Store) GetUsageTrigger(utId uuid.UUID) (tp.UsageTrigger, error) {
	var ut tp.UsageTrigger
	query := fmt.Sprintf("SELECT id, quantity, usageunit_title, task_id FROM %s WHERE id = $1", usageTriggerTableName)
	err := pg.db.QueryRow(query, utId).Scan(&ut.Id, &ut.Quantity, &ut.UsageUnit, &ut.TaskId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}

	return ut, nil
}

func (pg *Store) ListUsageTriggers() ([]tp.UsageTrigger, error) {
	var uts []tp.UsageTrigger
	query := fmt.Sprintf("SELECT id, quantity, usageunit_title, task_id FROM %s", usageTriggerTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "usage-trigger")
	}
	defer rows.Close()

	for rows.Next() {
		var ut tp.UsageTrigger
		if err := rows.Scan(&ut.Id, &ut.Quantity, &ut.UsageUnit, &ut.TaskId); err != nil {
			return nil, handleDbError(err, "usage-trigger")
		}
		uts = append(uts, ut)
	}

	return uts, nil
}

func (pg *Store) ListUsageTriggersByTaskId(taskId uuid.UUID) ([]tp.UsageTrigger, error) {
	var uts []tp.UsageTrigger
	query := fmt.Sprintf("SELECT id, quantity, usageunit_title, task_id FROM %s WHERE task_id = $1", usageTriggerTableName)
	rows, err := pg.db.Query(query, taskId)
	if err != nil {
		return nil, handleDbError(err, "usage-trigger")
	}
	defer rows.Close()

	for rows.Next() {
		var ut tp.UsageTrigger
		if err := rows.Scan(&ut.Id, &ut.Quantity, &ut.UsageUnit, &ut.TaskId); err != nil {
			return nil, handleDbError(err, "usage-trigger")
		}
		uts = append(uts, ut)
	}

	return uts, nil
}

func (pg *Store) UpdateUsageTrigger(utId uuid.UUID, ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	query := fmt.Sprintf("UPDATE %s SET quantity = $1, usageunit_title = $2, task_id = $3 WHERE id = $4", usageTriggerTableName)
	result, err := pg.db.Exec(query, ut.Quantity, ut.UsageUnit, ut.TaskId, utId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}
	if rowsAffected == 0 {
		return tp.UsageTrigger{}, errors.Wrapf(ae.ErrNotFound, "usage trigger with id '%s' not found", utId)
	}
	return ut, nil
}
