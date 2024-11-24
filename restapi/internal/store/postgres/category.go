package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var categoryTableName = "category"

func (pg *Store) CreateCategory(category tp.Category) (tp.Category, error) {
	category.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description) VALUES ($1, $2, $3)`, categoryTableName)
	_, err := pg.db.Exec(query, category.Id, category.Title, category.Description)
	if err != nil {
		return tp.Category{}, err
	}

	return category, nil
}

func (pg *Store) DeleteCategory(title string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1`, categoryTableName)
	_, err := pg.db.Exec(query, title)

	return err
}

func (pg *Store) ListCategories() ([]tp.Category, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s`, categoryTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories = []tp.Category{}
	for rows.Next() {
		var category tp.Category
		err = rows.Scan(&category.Id, &category.Title, &category.Description)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (pg *Store) GetCategory(title string) (tp.Category, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s WHERE title = $1`, categoryTableName)
	row := pg.db.QueryRow(query, title)

	var category tp.Category
	err := row.Scan(&category.Id, &category.Title, &category.Description)
	if err != nil {
		return tp.Category{}, err
	}

	return category, nil
}

func (pg *Store) UpdateCategory(title string, category tp.Category) (tp.Category, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE title = $3 returning id`, categoryTableName)
	err := pg.db.QueryRow(query, category.Title, category.Description, title).Scan(&category.Id)
	if err != nil {
		return tp.Category{}, err
	}

	return category, nil
}
