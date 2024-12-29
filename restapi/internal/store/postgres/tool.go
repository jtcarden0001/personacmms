package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var toolTableName = "tool"
var tToolTableName = "task_tool"
var woToolTableName = "workorder_tool"

func (pg *PostgresStore) AssociateToolWithTask(taskId uuid.UUID, toolId uuid.UUID, ts string) (tp.ToolSize, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (task_id, tool_id, size_note) 
			VALUES ($1, $2, $3)`,
		tToolTableName)

	_, err := pg.db.Exec(query, taskId, toolId, ts)
	if err != nil {
		return tp.ToolSize{}, handleDbError(err, "tool")
	}

	return pg.getToolSize(taskId, toolId)
}

func (pg *PostgresStore) AssociateToolWithWorkOrder(workOrderId uuid.UUID, toolId uuid.UUID, ts string) (tp.ToolSize, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (workorder_id, tool_id, size_note) 
			VALUES ($1, $2, $3)`,
		woToolTableName)

	_, err := pg.db.Exec(query, workOrderId, toolId, ts)
	if err != nil {
		return tp.ToolSize{}, handleDbError(err, "tool")
	}

	return pg.getToolSize(workOrderId, toolId)
}

func (pg *PostgresStore) CreateTool(t tp.Tool) (tp.Tool, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (id, title) 
			VALUES ($1, $2)`,
		toolTableName)

	_, err := pg.db.Exec(query, t.Id, t.Title)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return t, nil
}

func (pg *PostgresStore) DeleteTool(id uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = $1`,
		toolTableName)

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "tool")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "tool")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "tool with id '%s' not found", id)
	}

	return nil
}

func (pg *PostgresStore) DisassociateToolWithTask(taskId uuid.UUID, toolId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE task_id = $1 AND tool_id = $2`,
		tToolTableName)

	result, err := pg.db.Exec(query, taskId, toolId)
	if err != nil {
		return handleDbError(err, "tool")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "tool")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "tool with id '%s' not found for task with id '%s'", toolId, taskId)
	}

	return nil
}

func (pg *PostgresStore) DisassociateToolWithWorkOrder(workOrderId uuid.UUID, toolId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s
			WHERE workorder_id = $1 AND tool_id = $2`,
		woToolTableName)

	result, err := pg.db.Exec(query, workOrderId, toolId)
	if err != nil {
		return handleDbError(err, "tool")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "tool")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "tool with id '%s' not found for work order with id '%s'", toolId, workOrderId)
	}

	return nil
}

func (pg *PostgresStore) GetTool(id uuid.UUID) (tp.Tool, error) {
	query := fmt.Sprintf(`
			SELECT id, title 
			FROM %s 
			WHERE id = $1`,
		toolTableName)

	row := pg.db.QueryRow(query, id)

	var t tp.Tool
	err := row.Scan(&t.Id, &t.Title)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return t, nil
}

func (pg *PostgresStore) getToolSize(taskId uuid.UUID, toolId uuid.UUID) (tp.ToolSize, error) {
	query := fmt.Sprintf(`
			SELECT t.Id, t.Title, tt.size_note
			FROM %s t JOIN %s ts ON t.id = tt.tool_id
			WHERE ts.task_id = $1 AND ts.tool_id = $2`,
		toolTableName, tToolTableName)

	row := pg.db.QueryRow(query, taskId, toolId)

	var ts tp.ToolSize
	err := row.Scan(&ts.Id, &ts.Title, &ts.Size)
	if err != nil {
		return tp.ToolSize{}, handleDbError(err, "tool")
	}

	return ts, nil
}

func (pg *PostgresStore) ListTools() ([]tp.Tool, error) {
	query := fmt.Sprintf(`
			SELECT id, title 
			FROM %s`,
		toolTableName)

	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "tool")
	}
	defer rows.Close()

	var tools = []tp.Tool{}
	for rows.Next() {
		var t tp.Tool
		err = rows.Scan(&t.Id, &t.Title)
		if err != nil {
			return nil, handleDbError(err, "tool")
		}
		tools = append(tools, t)
	}

	return tools, nil
}

func (pg *PostgresStore) ListToolsByTask(taskId uuid.UUID) ([]tp.ToolSize, error) {
	query := fmt.Sprintf(`
			SELECT t.id, t.title, tt.size_note
			FROM %s t JOIN %s tt ON t.id = tt.tool_id
			WHERE tt.task_id = $1`,
		toolTableName, tToolTableName)

	rows, err := pg.db.Query(query, taskId)
	if err != nil {
		return nil, handleDbError(err, "tool")
	}

	var tools = []tp.ToolSize{}
	for rows.Next() {
		var ts tp.ToolSize
		err = rows.Scan(&ts.Id, &ts.Title, &ts.Size)
		if err != nil {
			return nil, handleDbError(err, "tool")
		}
		tools = append(tools, ts)
	}

	return tools, nil
}

func (pg *PostgresStore) ListToolsByWorkOrder(workOrderId uuid.UUID) ([]tp.ToolSize, error) {
	query := fmt.Sprintf(`
			SELECT t.id, t.title, wt.size_note
			FROM %s t JOIN %s wt ON t.id = wt.tool_id
			WHERE wt.workorder_id = $1`,
		toolTableName, woToolTableName)

	rows, err := pg.db.Query(query, workOrderId)
	if err != nil {
		return nil, handleDbError(err, "tool")
	}

	var tools = []tp.ToolSize{}
	for rows.Next() {
		var ts tp.ToolSize
		err = rows.Scan(&ts.Id, &ts.Title, &ts.Size)
		if err != nil {
			return nil, handleDbError(err, "tool")
		}
		tools = append(tools, ts)
	}

	return tools, nil
}

func (pg *PostgresStore) UpdateTool(t tp.Tool) (tp.Tool, error) {
	query := fmt.Sprintf(`
			UPDATE %s 
			SET title = $1
			WHERE id = $2`,
		toolTableName)
	result, err := pg.db.Exec(query, t.Title, t.Id)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	if rowsAffected == 0 {
		return tp.Tool{}, errors.Wrapf(ae.ErrNotFound, "tool with id %s not found", t.Id)
	}

	return t, nil
}
