package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var assetTaskToolTable = "task_tool"

func (pg *PostgresStore) CreateTaskTool(tool tp.TaskTool) (tp.TaskTool, error) {
	query := fmt.Sprintf("INSERT INTO %s (task_id, tool_id) VALUES ($1, $2)", assetTaskToolTable)
	_, err := pg.db.Exec(query, tool.TaskId, tool.ToolId)
	if err != nil {
		return tp.TaskTool{}, handleDbError(err, "task-tool")
	}

	return tool, nil
}

func (pg *PostgresStore) DeleteTaskTool(atId, tId uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE task_id = $1 AND tool_id = $2", assetTaskToolTable)
	result, err := pg.db.Exec(query, atId, tId)
	if err != nil {
		return handleDbError(err, "task-tool")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "task-tool")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "task tool with task_id '%s' and tool_id '%s' not found", atId, tId)
	}
	return nil
}

func (pg *PostgresStore) ListTaskTools() ([]tp.TaskTool, error) {
	query := fmt.Sprintf("SELECT task_id, tool_id FROM %s", assetTaskToolTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "task-tool")
	}
	defer rows.Close()

	var at []tp.TaskTool
	for rows.Next() {
		var e tp.TaskTool
		err = rows.Scan(&e.TaskId, &e.ToolId)
		if err != nil {
			return nil, handleDbError(err, "task-tool")
		}
		at = append(at, e)
	}

	return at, nil
}

// TODO: add testing for this function
func (pg *PostgresStore) ListTaskToolsByTaskId(atId uuid.UUID) ([]tp.TaskTool, error) {
	query := fmt.Sprintf("SELECT task_id, tool_id FROM %s WHERE task_id = $1", assetTaskToolTable)
	rows, err := pg.db.Query(query, atId)
	if err != nil {
		return nil, handleDbError(err, "task-tool")
	}
	defer rows.Close()

	var at []tp.TaskTool
	for rows.Next() {
		var e tp.TaskTool
		err = rows.Scan(&e.TaskId, &e.ToolId)
		if err != nil {
			return nil, handleDbError(err, "task-tool")
		}
		at = append(at, e)
	}

	return at, nil
}

func (pg *PostgresStore) GetTaskTool(atId, tId uuid.UUID) (tp.TaskTool, error) {
	query := fmt.Sprintf("SELECT task_id, tool_id FROM %s WHERE task_id = $1 AND tool_id = $2", assetTaskToolTable)
	var e tp.TaskTool
	err := pg.db.QueryRow(query, atId, tId).Scan(&e.TaskId, &e.ToolId)
	if err != nil {
		return tp.TaskTool{}, handleDbError(err, "task-tool")
	}

	return e, nil
}

func (pg *PostgresStore) UpdateTaskTool(atId, tId uuid.UUID, tool tp.TaskTool) (tp.TaskTool, error) {
	query := fmt.Sprintf("UPDATE %s SET tool_id = $1 WHERE task_id = $2 AND tool_id = $3", assetTaskToolTable)
	result, err := pg.db.Exec(query, tool.ToolId, atId, tId)
	if err != nil {
		return tp.TaskTool{}, handleDbError(err, "task-tool")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.TaskTool{}, handleDbError(err, "task-tool")
	}
	if rowsAffected == 0 {
		return tp.TaskTool{}, errors.Wrapf(ae.ErrNotFound, "task tool with task_id '%s' and tool_id '%s' not found", atId, tId)
	}

	tool.TaskId = atId
	return tool, nil
}
