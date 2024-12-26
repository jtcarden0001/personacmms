package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var toolTableName = "tool"

func (pg *PostgresStore) CreateTool(t tp.Tool) (tp.Tool, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2)`, toolTableName)
	_, err := pg.db.Exec(query, t.Id, t.Title)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return t, nil
}

func (pg *PostgresStore) DeleteTool(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, toolTableName)
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

func (pg *PostgresStore) GetTool(id uuid.UUID) (tp.Tool, error) {
	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE id = $1`, toolTableName)
	row := pg.db.QueryRow(query, id)

	var t tp.Tool
	err := row.Scan(&t.Id, &t.Title)
	if err != nil {
		return tp.Tool{}, handleDbError(err, "tool")
	}

	return t, nil
}

func (pg *PostgresStore) ListTools() ([]tp.Tool, error) {
	query := fmt.Sprintf(`SELECT id, title FROM %s`, toolTableName)
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

func (pg *PostgresStore) UpdateTool(t tp.Tool) (tp.Tool, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE id = $2`, toolTableName)
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
