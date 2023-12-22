package postgres

import (
	tp "github.com/jtcarden0001/personacmms/webapi/internal/types"
)

type Tool interface {
	CreateTool(string, string) (int, error)
	DeleteTool(int) error
	GetAllTool() ([]tp.Tool, error)
	GetTool(int) (tp.Tool, error)
	UpdateTool(int, string, string) error
}

func (pg *Store) CreateTool(title string, size string) (int, error) {
	query := `INSERT INTO tool (title, size) VALUES ($1, $2) returning id`
	var id int
	err := pg.db.QueryRow(query, title, size).Scan(&id)

	return id, err
}

func (pg *Store) DeleteTool(id int) error {
	query := `DELETE FROM tool WHERE id = $1`
	_, err := pg.db.Exec(query, id)
	return err
}

func (pg *Store) GetAllTool() ([]tp.Tool, error) {
	query := `SELECT id, title, size FROM tool`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tools []tp.Tool
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
	query := `SELECT id, title, size FROM tool WHERE id = $1`
	var t tp.Tool
	err := pg.db.QueryRow(query, id).Scan(&t.Id, &t.Title, &t.Size)

	return t, err
}

func (pg *Store) UpdateTool(id int, title string, size string) error {
	query := `UPDATE tool SET title = $1, size = $2 WHERE id = $3`
	_, err := pg.db.Exec(query, title, size, id)

	return err
}
