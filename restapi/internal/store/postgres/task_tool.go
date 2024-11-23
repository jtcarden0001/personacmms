package postgres

import (
	"database/sql"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type PreventativeTaskTool interface {
	CreatePreventativeTaskTool(int, int) error
	DeletePreventativeTaskTool(int, int) error
	GetAllPreventativeTaskTool() ([]tp.PreventativeTaskTool, error)
	GetAllPreventativeTaskToolByPreventativeTaskId(int) ([]tp.PreventativeTaskTool, error)
	GetPreventativeTaskTool(int, int) (tp.PreventativeTaskTool, error)
}

func (pg *Store) CreatePreventativeTaskTool(preventativeTaskId int, toolId int) error {
	query := `INSERT INTO preventativeTask_tool (preventativeTask_id, tool_id) VALUES ($1, $2)`
	_, err := pg.db.Exec(query, preventativeTaskId, toolId)

	return err
}

func (pg *Store) DeletePreventativeTaskTool(preventativeTaskId int, toolId int) error {
	query := `DELETE FROM preventativeTask_tool WHERE preventativeTask_id = $1 AND tool_id = $2`
	_, err := pg.db.Exec(query, preventativeTaskId, toolId)

	return err
}

func (pg *Store) GetAllPreventativeTaskTool() ([]tp.PreventativeTaskTool, error) {
	query := `SELECT preventativeTask_id, tool_id FROM preventativeTask_tool`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populatePreventativeTaskToolList(rows)
}

func (pg *Store) GetAllPreventativeTaskToolByPreventativeTaskId(preventativeTaskId int) ([]tp.PreventativeTaskTool, error) {
	query := `SELECT preventativeTask_id, tool_id FROM preventativeTask_tool WHERE preventativeTask_id = $1`
	rows, err := pg.db.Query(query, preventativeTaskId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return populatePreventativeTaskToolList(rows)
}

func (pg *Store) GetPreventativeTaskTool(preventativeTaskId int, toolId int) (tp.PreventativeTaskTool, error) {
	query := `SELECT preventativeTask_id, tool_id FROM preventativeTask_tool WHERE preventativeTask_id = $1 AND tool_id = $2`
	var tc tp.PreventativeTaskTool
	err := pg.db.QueryRow(query, preventativeTaskId, toolId).Scan(&tc.PreventativeTaskId, &tc.ToolId)

	return tc, err
}

func populatePreventativeTaskToolList(rows *sql.Rows) ([]tp.PreventativeTaskTool, error) {
	var PreventativeTaskTools []tp.PreventativeTaskTool
	for rows.Next() {
		var tc tp.PreventativeTaskTool
		err := rows.Scan(&tc.PreventativeTaskId, &tc.ToolId)
		if err != nil {
			return nil, err
		}
		PreventativeTaskTools = append(PreventativeTaskTools, tc)
	}

	return PreventativeTaskTools, nil
}
