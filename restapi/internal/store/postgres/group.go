package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var groupTableName = "assetgroup"

func (pg *Store) CreateGroup(grp tp.Group) (tp.Group, error) {
	//TODO: allow for group creation with a specified id ?
	id := uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2) returning id`, groupTableName)
	_, err := pg.db.Exec(query, id, grp.Title)
	if err != nil {
		return tp.Group{}, handleDbError(err, "group")
	}

	grp.Id = id
	return grp, nil
}

func (pg *Store) DeleteGroup(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, groupTableName)
	result, err := pg.db.Exec(query, title)
	if err != nil {
		return handleDbError(err, "group")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "group")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "group with title '%s' not found", title)
	}
	return nil
}

func (pg *Store) ListGroups() ([]tp.Group, error) {
	var groups = []tp.Group{}
	query := fmt.Sprintf(`SELECT id, title FROM %s`, groupTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return groups, handleDbError(err, "group")
	}
	defer rows.Close()

	for rows.Next() {
		var grp tp.Group
		err = rows.Scan(&grp.Id, &grp.Title)
		if err != nil {
			return nil, handleDbError(err, "group")
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
		return tp.Group{}, handleDbError(err, "group")
	}

	return grp, nil
}

func (pg *Store) UpdateGroup(oldtitle string, newGroup tp.Group) (tp.Group, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE title = $2 returning id, title`, groupTableName)
	err := pg.db.QueryRow(query, newGroup.Title, oldtitle).Scan(&newGroup.Id, &newGroup.Title)
	if err != nil {
		return tp.Group{}, handleDbError(err, "group")
	}

	return newGroup, nil
}
