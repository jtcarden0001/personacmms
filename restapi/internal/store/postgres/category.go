package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var categoryTableName = "category"

func (pg *PostgresStore) CreateCategory(category tp.Category) (tp.Category, error) {
	category.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description) VALUES ($1, $2, $3)`, categoryTableName)
	_, err := pg.db.Exec(query, category.Id, category.Title, category.Description)
	if err != nil {
		return tp.Category{}, handleDbError(err, "category")
	}

	return category, nil
}

func (pg *PostgresStore) DeleteCategory(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, categoryTableName)
	result, err := pg.db.Exec(query, title)
	if err != nil {
		return handleDbError(err, "category")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "category")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "category with title %s not found", title)
	}
	return nil
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

func (pg *PostgresStore) GetCategory(title string) (tp.Category, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s WHERE title = $1`, categoryTableName)
	row := pg.db.QueryRow(query, title)

	var category tp.Category
	err := row.Scan(&category.Id, &category.Title, &category.Description)
	if err != nil {
		return tp.Category{}, handleDbError(err, "category")
	}

	return category, nil
}

func (pg *PostgresStore) UpdateCategory(title string, category tp.Category) (tp.Category, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE title = $3 RETURNING id`, categoryTableName)
	err := pg.db.QueryRow(query, category.Title, category.Description, title).Scan(&category.Id)
	if err != nil {
		return tp.Category{}, handleDbError(err, "category")
	}

	return category, nil
}
