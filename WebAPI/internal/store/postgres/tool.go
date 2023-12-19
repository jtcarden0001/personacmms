package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
)

type Tool interface {
	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTools() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
}

type ToolTest interface {
	ResetSequenceTool(int) error
}

func (pg *Store) CreateTool(title string, size string) (int, error) {
	// TODO: add validation to prevent sql injection
	query := `INSERT INTO tool (title, size) VALUES ($1, $2) returning id`
	var id int
	err := pg.db.QueryRow(query, title, size).Scan(&id)

	return id, err
}

func (pg *Store) DeleteTool(id int) error {
	// TODO: add validation to prevent sql injection
	query := `DELETE FROM tool WHERE id = $1`
	_, err := pg.db.Exec(query, id)
	return err
}

func (pg *Store) GetAllTools() ([]tp.Tool, error) {
	// TODO: add validation to prevent sql injection
	var tools []tp.Tool
	query := `SELECT id, title, size FROM tool`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t tp.Tool
		err = rows.Scan(&t.Id, &t.Title, &t.Size)
		if err != nil {
			return nil, err
		}
		tools = append(tools, t)
	}

	return tools, nil
}

func (pg *Store) GetTool(id int) (tp.Tool, error) {
	// TODO: add validation to prevent sql injection
	var t tp.Tool
	query := `SELECT id, title, size FROM tool WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.Size)

	return t, err
}

func (pg *Store) UpdateTool(id int, title string, size string) error {
	// TODO: add validation to prevent sql injection
	query := `UPDATE tool SET title = $1, size = $2 WHERE id = $3`
	_, err := pg.db.Exec(query, title, size, id)

	return err
}

func (pg *Store) ResetSequenceTool(id int) error {
	query := fmt.Sprintf("ALTER SEQUENCE tool_id_seq RESTART WITH %d", id)
	_, err := pg.db.Query(query)
	return err
}
