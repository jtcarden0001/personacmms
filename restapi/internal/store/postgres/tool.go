package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Tool interface {
	CreateTool(tp.Tool) (tp.Tool, error)
	DeleteTool(string) error
	ListTools() ([]tp.Tool, error)
	GetTool(string) (tp.Tool, error)
	UpdateTool(string, tp.Tool) (tp.Tool, error)
}

var toolTableName = "tool"

func (pg *Store) CreateTool(tool tp.Tool) (tp.Tool, error) {
	tool.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, size) VALUES ($1, $2, $3)`, toolTableName)
	_, err := pg.db.Exec(query, tool.Id, tool.Title, tool.Size)
	if err != nil {
		return tp.Tool{}, err
	}

	return tool, nil
}

func (pg *Store) DeleteTool(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, toolTableName)
	_, err := pg.db.Exec(query, title)

	return err
}

func (pg *Store) ListTools() ([]tp.Tool, error) {
	query := fmt.Sprintf(`SELECT id, title, size FROM %s`, toolTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tools = []tp.Tool{}
	for rows.Next() {
		var tool tp.Tool
		err = rows.Scan(&tool.Id, &tool.Title, &tool.Size)
		if err != nil {
			return nil, err
		}
		tools = append(tools, tool)
	}

	return tools, nil
}

func (pg *Store) GetTool(title string) (tp.Tool, error) {
	query := fmt.Sprintf(`SELECT id, title, size FROM %s WHERE title = $1`, toolTableName)
	row := pg.db.QueryRow(query, title)

	var tool tp.Tool
	err := row.Scan(&tool.Id, &tool.Title, &tool.Size)
	if err != nil {
		return tp.Tool{}, err
	}

	return tool, nil
}

func (pg *Store) UpdateTool(title string, tool tp.Tool) (tp.Tool, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, size = $2 WHERE title = $3`, toolTableName)
	_, err := pg.db.Exec(query, tool.Title, tool.Size, title)
	if err != nil {
		return tp.Tool{}, err
	}

	return tool, nil
}
