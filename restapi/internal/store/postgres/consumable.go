package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var consumableTableName = "consumable"

func (pg *PostgresStore) CreateConsumable(c tp.Consumable) (tp.Consumable, error) {
	//TODO: allow for group creation with a specified id ?
	id := uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title) VALUES ($1, $2) returning id`, consumableTableName)
	_, err := pg.db.Exec(query, id, c.Title)
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	c.Id = id
	return c, nil
}

func (pg *PostgresStore) DeleteConsumable(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, consumableTableName)
	result, err := pg.db.Exec(query, title)
	if err != nil {
		return handleDbError(err, "consumable")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "consumable")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "consumable with title %s not found", title)
	}
	return nil
}

func (pg *PostgresStore) ListConsumables() ([]tp.Consumable, error) {
	var consumables = []tp.Consumable{}
	query := fmt.Sprintf(`SELECT id, title FROM %s`, consumableTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return consumables, handleDbError(err, "consumable")
	}
	defer rows.Close()

	for rows.Next() {
		var c tp.Consumable
		err = rows.Scan(&c.Id, &c.Title)
		if err != nil {
			return nil, handleDbError(err, "consumable")
		}
		consumables = append(consumables, c)
	}

	return consumables, nil
}

// TODO: implement testing for this
func (pg *PostgresStore) GetConsumableById(id uid.UUID) (tp.Consumable, error) {
	var c tp.Consumable
	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE id = $1`, consumableTableName)
	err := pg.db.QueryRow(query, id).Scan(&c.Id, &c.Title)
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	return c, nil
}

func (pg *PostgresStore) GetConsumableByTitle(title string) (tp.Consumable, error) {
	var c tp.Consumable
	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE title = $1`, consumableTableName)
	err := pg.db.QueryRow(query, title).Scan(&c.Id, &c.Title)
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	return c, nil
}

func (pg *PostgresStore) UpdateConsumable(title string, c tp.Consumable) (tp.Consumable, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1 WHERE title = $2 RETURNING id`, consumableTableName)
	err := pg.db.QueryRow(query, c.Title, title).Scan(&c.Id)
	if err != nil {
		return tp.Consumable{}, handleDbError(err, "consumable")
	}

	return c, nil
}
