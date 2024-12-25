package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var categoryTableName = "category"

func (pg *PostgresStore) CreateCategory(c tp.Category) (tp.Category, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description) VALUES ($1, $2, $3)`, categoryTableName)
	_, err := pg.db.Exec(query, c.Id, c.Title, c.Description)
	if err != nil {
		return tp.Category{}, handleDbError(err, "category")
	}

	return c, nil
}

func (pg *PostgresStore) DeleteCategory(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, categoryTableName)
	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "category")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "category")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "category with id %s not found", id)
	}
	return nil
}

func (pg *PostgresStore) GetCategory(id uuid.UUID) (tp.Category, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s WHERE id = $1`, categoryTableName)
	row := pg.db.QueryRow(query, id)

	var category tp.Category
	err := row.Scan(&category.Id, &category.Title, &category.Description)
	if err != nil {
		return tp.Category{}, handleDbError(err, "category")
	}

	return category, nil
}

func (pg *PostgresStore) ListCategories() ([]tp.Category, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s`, categoryTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "category")
	}
	defer rows.Close()

	var categories = []tp.Category{}
	for rows.Next() {
		var category tp.Category
		err = rows.Scan(&category.Id, &category.Title, &category.Description)
		if err != nil {
			return nil, handleDbError(err, "category")
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (pg *PostgresStore) UpdateCategory(c tp.Category) (tp.Category, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE id = $3`, categoryTableName)
	result, err := pg.db.Exec(query, c.Title, c.Description, c.Id)
	if err != nil {
		return tp.Category{}, handleDbError(err, "category")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.Category{}, handleDbError(err, "category")
	}

	if rowsAffected == 0 {
		return tp.Category{}, errors.Wrapf(ae.ErrNotFound, "category with id %s not found", c.Id)
	}

	return c, nil
}
