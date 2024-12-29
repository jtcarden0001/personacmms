package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var usageTriggerTableName = "usagetrigger"

func (pg *PostgresStore) CreateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (id, quantity, usageunit_id, task_id) 
			VALUES ($1, $2, $3, $4)`,
		usageTriggerTableName)

	usageUnitId, err := pg.getUsageUnitIdFromTitle(ut.UsageUnit)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	_, err = pg.db.Exec(query, ut.Id, ut.Quantity, usageUnitId, ut.TaskId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}

	return ut, nil
}

func (pg *PostgresStore) DeleteUsageTrigger(id uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = $1`,
		usageTriggerTableName)

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "usage-trigger")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "usage-trigger")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "usage trigger with id '%s' not found", id)
	}

	return nil
}

func (pg *PostgresStore) DeleteUsageTriggerFromTask(taskId uuid.UUID, utId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE task_id = $1 AND id = $2`,
		usageTriggerTableName)

	result, err := pg.db.Exec(query, taskId, utId)
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

func (pg *PostgresStore) GetUsageTrigger(id uuid.UUID) (tp.UsageTrigger, error) {
	query := fmt.Sprintf(`
			SELECT id, quantity, usageunit_id, task_id 
			FROM %s 
			WHERE id = $1`,
		usageTriggerTableName)

	var ut tp.UsageTrigger
	var usageUnitId uuid.UUID
	err := pg.db.QueryRow(query, id).Scan(&ut.Id, &ut.Quantity, &usageUnitId, &ut.TaskId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}

	usageUnitTitle, err := pg.getUsageUnitTitleFromId(usageUnitId)
	if err != nil {
		return tp.UsageTrigger{}, errors.Wrapf(err, "failed to get usage unit title for id '%s' during UsageTrigger get", usageUnitId)
	}

	ut.UsageUnit = usageUnitTitle
	return ut, nil
}

func (pg *PostgresStore) ListUsageTriggers() ([]tp.UsageTrigger, error) {
	var uts []tp.UsageTrigger
	query := fmt.Sprintf(`
			SELECT id, quantity, usageunit_id, task_id 
			FROM %s`,
		usageTriggerTableName)

	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "usage-trigger")
	}
	defer rows.Close()

	for rows.Next() {
		var ut tp.UsageTrigger
		var usageUnitId uuid.UUID
		if err := rows.Scan(&ut.Id, &ut.Quantity, &usageUnitId, &ut.TaskId); err != nil {
			return nil, handleDbError(err, "usage-trigger")
		}

		usageUnitTitle, err := pg.getUsageUnitTitleFromId(usageUnitId)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get usage unit title for id '%s' during UsageTrigger list", usageUnitId)
		}

		ut.UsageUnit = usageUnitTitle
		uts = append(uts, ut)
	}

	return uts, nil
}

func (pg *PostgresStore) ListUsageTriggersByTask(taskId uuid.UUID) ([]tp.UsageTrigger, error) {
	var uts []tp.UsageTrigger
	query := fmt.Sprintf(`
			SELECT id, quantity, usageunit_id, task_id
			FROM %s
			WHERE task_id = $1`,
		usageTriggerTableName)

	rows, err := pg.db.Query(query, taskId)
	if err != nil {
		return nil, handleDbError(err, "usage-trigger")
	}
	defer rows.Close()

	for rows.Next() {
		var ut tp.UsageTrigger
		var usageUnitId uuid.UUID
		if err := rows.Scan(&ut.Id, &ut.Quantity, &usageUnitId, &ut.TaskId); err != nil {
			return nil, handleDbError(err, "usage-trigger")
		}

		usageUnitTitle, err := pg.getUsageUnitTitleFromId(usageUnitId)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get usage unit title for id '%s' during UsageTrigger list", usageUnitId)
		}

		ut.UsageUnit = usageUnitTitle
		uts = append(uts, ut)
	}

	return uts, nil
}

func (pg *PostgresStore) UpdateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	query := fmt.Sprintf(`
			UPDATE %s 
			SET quantity = $1, usageunit_id = $2, task_id = $3 
			WHERE id = $4`,
		usageTriggerTableName)

	usageUnitId, err := pg.getUsageUnitIdFromTitle(ut.UsageUnit)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	result, err := pg.db.Exec(query, ut.Quantity, usageUnitId, ut.TaskId, ut.Id)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err, "usage-trigger")
	}

	if rowsAffected == 0 {
		return tp.UsageTrigger{}, errors.Wrapf(ae.ErrNotFound, "usage trigger with id '%s' not found", ut.Id)
	}

	return ut, nil
}

var usageUnitTableName = "usageunit"

func (pg *PostgresStore) getUsageUnitIdFromTitle(title string) (uuid.UUID, error) {
	query := fmt.Sprintf(`
			SELECT id 
			FROM %s 
			WHERE title = $1`,
		usageUnitTableName)

	var id uuid.UUID
	err := pg.db.QueryRow(query, title).Scan(&id)
	if err != nil {
		return uuid.UUID{}, handleDbError(err, "usage-trigger")
	}

	return id, nil
}

func (pg *PostgresStore) getUsageUnitTitleFromId(id uuid.UUID) (string, error) {
	query := fmt.Sprintf(`
			SELECT title 
			FROM %s 
			WHERE id = $1`,
		usageUnitTableName)
	var title string
	err := pg.db.QueryRow(query, id).Scan(&title)
	if err != nil {
		return "", handleDbError(err, "usage-trigger")
	}

	return title, nil
}
