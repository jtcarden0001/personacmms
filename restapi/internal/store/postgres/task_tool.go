package postgres

import (
	"database/sql"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type TaskTool interface {
	CreateTaskTool(int, int) error
	DeleteTaskTool(int, int) error
	GetAllTaskTool() ([]tp.TaskTool, error)
	GetAllTaskToolByTaskId(int) ([]tp.TaskTool, error)
	GetTaskTool(int, int) (tp.TaskTool, error)
}

func (pg *Store) CreateTaskTool(preventativeTaskId int, toolId int) error {
	query := `INSERT INTO preventativeTask_tool (preventativeTask_id, tool_id) VALUES ($1, $2)`
	_, err := pg.db.Exec(query, preventativeTaskId, toolId)

	return err
}

func (pg *Store) DeleteTaskTool(preventativeTaskId int, toolId int) error {
	query := `DELETE FROM preventativeTask_tool WHERE preventativeTask_id = $1 AND tool_id = $2`
	_, err := pg.db.Exec(query, preventativeTaskId, toolId)

	return err
}

func (pg *Store) GetAllTaskTool() ([]tp.TaskTool, error) {
	query := `SELECT preventativeTask_id, tool_id FROM preventativeTask_tool`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populateTaskToolList(rows)
}

func (pg *Store) GetAllTaskToolByTaskId(preventativeTaskId int) ([]tp.TaskTool, error) {
	query := `SELECT preventativeTask_id, tool_id FROM preventativeTask_tool WHERE preventativeTask_id = $1`
	rows, err := pg.db.Query(query, preventativeTaskId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populateTaskToolList(rows)
}

func (pg *Store) GetTaskTool(preventativeTaskId int, toolId int) (tp.TaskTool, error) {
	query := `SELECT preventativeTask_id, tool_id FROM preventativeTask_tool WHERE preventativeTask_id = $1 AND tool_id = $2`
	var tc tp.TaskTool
	err := pg.db.QueryRow(query, preventativeTaskId, toolId).Scan(&tc.TaskId, &tc.ToolId)

	return tc, err
}

func populateTaskToolList(rows *sql.Rows) ([]tp.TaskTool, error) {
	var TaskTools []tp.TaskTool
	for rows.Next() {
		var tc tp.TaskTool
		err := rows.Scan(&tc.TaskId, &tc.ToolId)
		if err != nil {
			return nil, err
		}
		TaskTools = append(TaskTools, tc)
	}

	return TaskTools, nil
}
