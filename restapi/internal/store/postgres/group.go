package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var groupTableName = "agroup" // not group because it's a reserved word in postgres

func (pg *PostgresStore) CreateGroup(g tp.Group) (tp.Group, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (id, title) 
			VALUES ($1, $2)`,
		groupTableName)

	_, err := pg.db.Exec(query, g.Id, g.Title)
	if err != nil {
		return tp.Group{}, handleDbError(err, "group")
	}

	return g, nil
}

func (pg *PostgresStore) DeleteGroup(id uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = $1`,
		groupTableName)

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "group")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "group")
	}

	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "group with id '%s' not found", id)
	}

	return nil
}

func (pg *PostgresStore) GetGroup(id uuid.UUID) (tp.Group, error) {
	query := fmt.Sprintf(`
			SELECT id, title 
			FROM %s 
			WHERE id = $1`,
		groupTableName)

	var g tp.Group
	err := pg.db.QueryRow(query, id).Scan(&g.Id, &g.Title)
	if err != nil {
		return tp.Group{}, handleDbError(err, "group")
	}

	return g, nil
}

func (pg *PostgresStore) ListGroups() ([]tp.Group, error) {
	var groups = []tp.Group{}
	query := fmt.Sprintf(`
			SELECT id, title 
			FROM %s`,
		groupTableName)

	rows, err := pg.db.Query(query)
	if err != nil {
		return groups, handleDbError(err, "group")
	}
	defer rows.Close()

	for rows.Next() {
		var g tp.Group
		err = rows.Scan(&g.Id, &g.Title)
		if err != nil {
			return nil, handleDbError(err, "group")
		}
		groups = append(groups, g)
	}

	return groups, nil
}

func (pg *PostgresStore) ListGroupsByAsset(assetId uuid.UUID) ([]tp.Group, error) {
	query := fmt.Sprintf(`
			SELECT g.id, g.title
			FROM %s g JOIN %s ga ON g.id = ga.group_id
			WHERE ga.asset_id = $1`,
		groupTableName,
		gpAssetTableName)

	rows, err := pg.db.Query(query, assetId)
	if err != nil {
		return nil, handleDbError(err, "group")
	}

	var groups = []tp.Group{}
	for rows.Next() {
		var g tp.Group
		err = rows.Scan(&g.Id, &g.Title)
		if err != nil {
			return nil, handleDbError(err, "group")
		}
		groups = append(groups, g)
	}

	return groups, nil
}

func (pg *PostgresStore) UpdateGroup(g tp.Group) (tp.Group, error) {
	query := fmt.Sprintf(`
			UPDATE %s SET title = $1 
			WHERE id = $2`,
		groupTableName)

	result, err := pg.db.Exec(query, g.Title, g.Id)
	if err != nil {
		return tp.Group{}, handleDbError(err, "group")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.Group{}, handleDbError(err, "group")
	}

	if rowsAffected == 0 {
		return tp.Group{}, errors.Wrapf(ae.ErrNotFound, "group with id '%s' not found", g.Id)
	}

	return g, nil
}
