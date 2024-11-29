package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var groupTableName = "assetgroup"

func (pg *Store) CreateGroup(grp tp.Group) (tp.Group, error) {
	//TODO: allow for group creation with a specified id ?
	id := uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2) returning id`, groupTableName)
	_, err := pg.db.Exec(query, id, grp.Title)
	if err != nil {
		return tp.Group{}, handleDbError(err)
	}

	grp.Id = id
	return grp, nil
}

func (pg *Store) DeleteGroup(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, groupTableName)
	_, err := pg.db.Exec(query, title)

	return handleDbError(err)
}

func (pg *Store) ListGroups() ([]tp.Group, error) {
	var groups = []tp.Group{}
	query := fmt.Sprintf(`SELECT id, title FROM %s`, groupTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return groups, handleDbError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var grp tp.Group
		err = rows.Scan(&grp.Id, &grp.Title)
		if err != nil {
			return nil, handleDbError(err)
		}
		groups = append(groups, grp)
	}

	return groups, nil
}

func (pg *Store) GetGroup(title string) (tp.Group, error) {
	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE title = $1`, groupTableName)
	var grp tp.Group
	err := pg.db.QueryRow(query, title).Scan(&grp.Id, &grp.Title)
	if err != nil {
		return tp.Group{}, handleDbError(err)
	}

	return grp, nil
}

func (pg *Store) UpdateGroup(oldtitle string, newGroup tp.Group) (tp.Group, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE title = $2 returning id, title`, groupTableName)
	var grp tp.Group
	err := pg.db.QueryRow(query, newGroup.Title, oldtitle).Scan(&grp.Id, &grp.Title)
	if err != nil {
		return tp.Group{}, handleDbError(err)
	}

	return grp, nil
}
