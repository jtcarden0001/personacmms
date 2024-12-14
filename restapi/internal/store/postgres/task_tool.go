package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var assetTaskToolTable = "task_tool"

func (pg *Store) CreateTaskTool(tool tp.TaskTool) (tp.TaskTool, error) {
	query := fmt.Sprintf("INSERT INTO %s (task_id, tool_id) VALUES ($1, $2)", assetTaskToolTable)
	_, err := pg.db.Exec(query, tool.TaskId, tool.ToolId)
	if err != nil {
		return tp.TaskTool{}, handleDbError(err, "task-tool")
	}

	return tool, nil
}

func (pg *Store) DeleteTaskTool(atId, tId tp.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE task_id = $1 AND tool_id = $2", assetTaskToolTable)
	_, err := pg.db.Exec(query, atId, tId)
	if err != nil {
		return handleDbError(err, "task-tool")
	}

	return nil
}

func (pg *Store) ListTaskTools() ([]tp.TaskTool, error) {
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
func (pg *Store) ListTaskToolsByTaskId(atId tp.UUID) ([]tp.TaskTool, error) {
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

func (pg *Store) GetTaskTool(atId, tId tp.UUID) (tp.TaskTool, error) {
	query := fmt.Sprintf("SELECT task_id, tool_id FROM %s WHERE task_id = $1 AND tool_id = $2", assetTaskToolTable)
	var e tp.TaskTool
	err := pg.db.QueryRow(query, atId, tId).Scan(&e.TaskId, &e.ToolId)
	if err != nil {
		return tp.TaskTool{}, handleDbError(err, "task-tool")
	}

	return e, nil
}
