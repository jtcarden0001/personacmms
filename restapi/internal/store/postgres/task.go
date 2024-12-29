package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var taskTable = "task"

func (pg *PostgresStore) CreateTask(t tp.Task) (tp.Task, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (id, title, instructions, asset_id) 
			VALUES ($1, $2, $3, $4)`,
		taskTable)

	_, err := pg.db.Exec(query, t.Id, t.Title, t.Instructions, t.AssetId)
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}

	return t, nil
}

func (pg *PostgresStore) DeleteTask(id uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = $1`,
		taskTable)

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "tasks")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "tasks")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "task with id '%s' not found", id)
	}

	return nil
}

func (pg *PostgresStore) DeleteTaskFromAsset(assetId uuid.UUID, taskId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE asset_id = $1 AND id = $2`,
		taskTable)

	result, err := pg.db.Exec(query, assetId, taskId)
	if err != nil {
		return handleDbError(err, "tasks")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "tasks")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "task with id '%s' not found for asset with id '%s'", taskId, assetId)
	}

	return nil
}

func (pg *PostgresStore) GetTask(id uuid.UUID) (tp.Task, error) {
	query := fmt.Sprintf(`
			SELECT id, title, instructions, asset_id 
			FROM %s 
			WHERE id = $1`,
		taskTable)

	var t tp.Task
	err := pg.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.Instructions, &t.AssetId)
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}

	return t, nil
}

func (pg *PostgresStore) ListTasks() ([]tp.Task, error) {
	query := fmt.Sprintf(`
			SELECT id, title, instructions, asset_id 
			FROM %s`,
		taskTable)

	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "tasks")
	}
	defer rows.Close()

	var tasks []tp.Task
	for rows.Next() {
		var t tp.Task
		err = rows.Scan(&t.Id, &t.Title, &t.Instructions, &t.AssetId)
		if err != nil {
			return nil, handleDbError(err, "tasks")
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (pg *PostgresStore) ListTasksByAsset(assetId uuid.UUID) ([]tp.Task, error) {
	query := fmt.Sprintf(`
			SELECT id, title, instructions, asset_id
			FROM %s
			WHERE asset_id = $1`,
		taskTable)

	rows, err := pg.db.Query(query, assetId)
	if err != nil {
		return nil, handleDbError(err, "tasks")
	}
	defer rows.Close()

	var tasks []tp.Task
	for rows.Next() {
		var t tp.Task
		err = rows.Scan(&t.Id, &t.Title, &t.Instructions, &t.AssetId)
		if err != nil {
			return nil, handleDbError(err, "tasks")
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (pg *PostgresStore) UpdateTask(t tp.Task) (tp.Task, error) {
	query := fmt.Sprintf(`
			UPDATE %s 
			SET title = $1, instructions = $2, asset_id = $3 
			WHERE id = $4`,
		taskTable)
	result, err := pg.db.Exec(query, t.Title, t.Instructions, t.AssetId, t.Id)
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.Task{}, handleDbError(err, "tasks")
	}

	if rowsAffected == 0 {
		return tp.Task{}, errors.Wrapf(ae.ErrNotFound, "task with id [%s] not found", t.Id)
	}

	return t, nil
}
