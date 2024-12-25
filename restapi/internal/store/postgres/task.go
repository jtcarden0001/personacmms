package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var assetTaskTable = "task"

func (pg *PostgresStore) CreateTask(at tp.Task) (tp.Task, error) {
	at.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, unique_instructions, asset_id, tasktemplate_id) VALUES ($1, $2, $3, $4, $5)`, assetTaskTable)
	_, err := pg.db.Exec(query, at.Id, at.Title, at.Instructions, at.AssetId, at.TaskTemplateId)
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}

	return at, nil
}

func (pg *PostgresStore) DeleteTask(atId uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, assetTaskTable)
	result, err := pg.db.Exec(query, atId)
	if err != nil {
		return handleDbError(err, "tasks")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "tasks")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "task with id '%s' not found", atId)
	}
	return nil
}

func (pg *PostgresStore) GetTask(atId uuid.UUID) (tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, tasktemplate_id FROM %s WHERE id = $1`, assetTaskTable)
	var e tp.Task
	err := pg.db.QueryRow(query, atId).Scan(&e.Id, &e.Title, &e.Instructions, &e.AssetId, &e.TaskTemplateId)
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}

	return e, nil
}

func (pg *PostgresStore) GetTaskByAssetId(assetId uuid.UUID, taskId uuid.UUID) (tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, tasktemplate_id FROM %s WHERE asset_id = $1 AND id = $2`, assetTaskTable)
	var e tp.Task
	err := pg.db.QueryRow(query, assetId, taskId).Scan(&e.Id, &e.Title, &e.Instructions, &e.AssetId, &e.TaskTemplateId)
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}

	return e, nil
}

func (pg *PostgresStore) ListTasks() ([]tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, tasktemplate_id FROM %s`, assetTaskTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "tasks")
	}
	defer rows.Close()

	var ts []tp.Task
	for rows.Next() {
		var e tp.Task
		err = rows.Scan(&e.Id, &e.Title, &e.Instructions, &e.AssetId, &e.TaskTemplateId)
		if err != nil {
			return nil, handleDbError(err, "tasks")
		}
		ts = append(ts, e)
	}

	return ts, nil
}

// TODO: add a test for this
func (pg *PostgresStore) ListTasksByAssetId(assetId uuid.UUID) ([]tp.Task, error) {
	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, tasktemplate_id FROM %s WHERE asset_id = $1`, assetTaskTable)
	rows, err := pg.db.Query(query, assetId)
	if err != nil {
		return nil, handleDbError(err, "tasks")
	}
	defer rows.Close()

	var ts []tp.Task
	for rows.Next() {
		var e tp.Task
		err = rows.Scan(&e.Id, &e.Title, &e.Instructions, &e.AssetId, &e.TaskTemplateId)
		if err != nil {
			return nil, handleDbError(err, "tasks")
		}
		ts = append(ts, e)
	}

	return ts, nil
}

func (pg *PostgresStore) UpdateTask(atId uuid.UUID, at tp.Task) (tp.Task, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, unique_instructions = $2, asset_id = $3, tasktemplate_id = $4 WHERE id = $5`, assetTaskTable)
	result, err := pg.db.Exec(query, at.Title, at.Instructions, at.AssetId, at.TaskTemplateId, atId)
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}
	if rowsAffected == 0 {
		return tp.Task{}, errors.Wrapf(ae.ErrNotFound, "task with id '%s' not found", atId)
	}
	at.Id = atId
	return at, nil
}
