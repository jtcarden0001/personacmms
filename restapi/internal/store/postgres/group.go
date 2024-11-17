package postgres

import (
	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Group interface {
	CreateGroup(tp.Group) (tp.Group, error)
	DeleteGroup(string) error
	ListGroups() ([]tp.Group, error)
	GetGroup(string) (tp.Group, error)
	UpdateGroup(string, string) (tp.Group, error)
}

func (pg *Store) CreateGroup(grp tp.Group) (tp.Group, error) {
	//TODO: allow for group creation with a specified id ?
	id := uid.New()
	query := `INSERT INTO group (id, title) VALUES ($1, $2) returning id`
	_, err := pg.db.Exec(query, id.String(), grp.Title)
	if err != nil {
		return tp.Group{}, err
	}

	grp.Id = id
	return grp, nil
}

func (pg *Store) DeleteGroup(title string) error {
	query := `DELETE FROM group WHERE title = $1`
	_, err := pg.db.Exec(query, title)

	return err
}

func (pg *Store) ListGroups() ([]tp.Group, error) {
	query := `SELECT id, title FROM group`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []tp.Group
	for rows.Next() {
		var grp tp.Group
		err = rows.Scan(&grp.Id, &grp.Title)
		if err != nil {
			return nil, err
		}
		groups = append(groups, grp)
	}

	return groups, nil
}

func (pg *Store) GetGroup(title string) (tp.Group, error) {
	query := `SELECT id, title FROM group WHERE title = $1`
	var grp tp.Group
	err := pg.db.QueryRow(query, title).Scan(&grp.Id, &grp.Title)
	if err != nil {
		return tp.Group{}, err
	}

	return grp, nil
}

func (pg *Store) UpdateGroup(title string, newTitle string) (tp.Group, error) {
	query := `UPDATE group SET title = $1 WHERE title = $2 returning id`
	var grp tp.Group
	err := pg.db.QueryRow(query, newTitle, title).Scan(&grp.Id)
	if err != nil {
		return tp.Group{}, err
	}

	grp.Title = newTitle

	return grp, nil
}
