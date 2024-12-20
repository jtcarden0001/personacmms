package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var toolTableName = "tool"

func (pg *PostgresStore) CreateTool(tool tp.Tool) (tp.Tool, error) {
	tool.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, size) VALUES ($1, $2, $3)`, toolTableName)
	_, err := pg.db.Exec(query, tool.Id, tool.Title, tool.Size)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return tool, nil
}

func (pg *PostgresStore) DeleteTool(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, toolTableName)
	result, err := pg.db.Exec(query, title)
	if err != nil {
		return handleDbError(err, "tool")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "tool")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "tool with title '%s' not found", title)
	}
	return nil
}

func (pg *PostgresStore) ListTools() ([]tp.Tool, error) {
	query := fmt.Sprintf(`SELECT id, title, size FROM %s`, toolTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "tool")
	}
	defer rows.Close()

	var tools = []tp.Tool{}
	for rows.Next() {
		var tool tp.Tool
		err = rows.Scan(&tool.Id, &tool.Title, &tool.Size)
		if err != nil {
			return nil, handleDbError(err, "tool")
		}
		tools = append(tools, tool)
	}

	return tools, nil
}

func (pg *PostgresStore) GetTool(title string) (tp.Tool, error) {
	query := fmt.Sprintf(`SELECT id, title, size FROM %s WHERE title = $1`, toolTableName)
	row := pg.db.QueryRow(query, title)

	var tool tp.Tool
	err := row.Scan(&tool.Id, &tool.Title, &tool.Size)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return tool, nil
}

// TODO: add testing for this.
func (pg *PostgresStore) GetToolById(id uuid.UUID) (tp.Tool, error) {
	query := fmt.Sprintf(`SELECT id, title, size FROM %s WHERE id = $1`, toolTableName)
	row := pg.db.QueryRow(query, id)

	var tool tp.Tool
	err := row.Scan(&tool.Id, &tool.Title, &tool.Size)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return tool, nil
}

func (pg *PostgresStore) UpdateTool(title string, tool tp.Tool) (tp.Tool, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, size = $2 WHERE title = $3 returning id`, toolTableName)
	err := pg.db.QueryRow(query, tool.Title, tool.Size, title).Scan(&tool.Id)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return tool, nil
}
