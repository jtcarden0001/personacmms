package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/projects/webapi/internal/types"
)

func (pg *Store) CreateTool(title string) (int, error) {
	// TODO: add validation to prevent sql injection
	query := `INSERT INTO tool (title) VALUES ($1) returning id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

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
	var tool []tp.Tool
	query := `SELECT id, title FROM tool`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var e tp.Tool
		err = rows.Scan(&e.Id, &e.Title)
		if err != nil {
			return nil, err
		}
		tool = append(tool, e)
	}

	return tool, nil
}

func (pg *Store) GetTool(id int) (tp.Tool, error) {
	// TODO: add validation to prevent sql injection
	var t tp.Tool
	query := `SELECT id, title FROM tool WHERE id = $1`
	err := pg.db.QueryRow(query, id).Scan(&t.Id, &t.Title)

	return t, err
}

func (pg *Store) UpdateTool(id int, title string) error {
	// TODO: add validation to prevent sql injection
	query := `UPDATE tool SET title = $1 WHERE id = $3`
	_, err := pg.db.Exec(query, title, id)

	return err
}

func (pg *Store) ResetSequenceTool(id int) error {
	query := fmt.Sprintf("ALTER SEQUENCE tool_id_seq RESTART WITH %d", id)
	_, err := pg.db.Query(query)
	return err
}
