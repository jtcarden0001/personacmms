package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var usageTriggerTableName = "usagetrigger"

func (pg *Store) CreateUsageTrigger(ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	ut.Id = uuid.New()
	query := fmt.Sprintf("INSERT INTO %s (id, quantity, usageunit_title, assettask_id) VALUES ($1, $2, $3, $4)", usageTriggerTableName)
	_, err := pg.db.Exec(query, ut.Id, ut.Quantity, ut.UsageUnit, ut.AssetTaskId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err)
	}

	return ut, nil
}

func (pg *Store) DeleteUsageTrigger(utId uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usageTriggerTableName)
	_, err := pg.db.Exec(query, utId)
	if err != nil {
		return handleDbError(err)
	}

	return nil
}

func (pg *Store) GetUsageTrigger(utId uuid.UUID) (tp.UsageTrigger, error) {
	var ut tp.UsageTrigger
	query := fmt.Sprintf("SELECT id, quantity, usageunit_title, assettask_id FROM %s WHERE id = $1", usageTriggerTableName)
	err := pg.db.QueryRow(query, utId).Scan(&ut.Id, &ut.Quantity, &ut.UsageUnit, &ut.AssetTaskId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err)
	}

	return ut, nil
}

func (pg *Store) ListUsageTriggers() ([]tp.UsageTrigger, error) {
	var uts []tp.UsageTrigger
	query := fmt.Sprintf("SELECT id, quantity, usageunit_title, assettask_id FROM %s", usageTriggerTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var ut tp.UsageTrigger
		if err := rows.Scan(&ut.Id, &ut.Quantity, &ut.UsageUnit, &ut.AssetTaskId); err != nil {
			return nil, handleDbError(err)
		}
		uts = append(uts, ut)
	}

	return uts, nil
}

func (pg *Store) UpdateUsageTrigger(utId uuid.UUID, ut tp.UsageTrigger) (tp.UsageTrigger, error) {
	query := fmt.Sprintf("UPDATE %s SET quantity = $1, usageunit_title = $2, assettask_id = $3 WHERE id = $4", usageTriggerTableName)
	_, err := pg.db.Exec(query, ut.Quantity, ut.UsageUnit, ut.AssetTaskId, utId)
	if err != nil {
		return tp.UsageTrigger{}, handleDbError(err)
	}

	return ut, nil
}
