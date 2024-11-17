package postgres

import (
	"fmt"

	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Category interface {
	CreateCategory(tp.Category) (tp.Category, error)
	DeleteCategory(string) error
	ListCategories() ([]tp.Category, error)
	GetCategory(string) (tp.Category, error)
	UpdateCategory(string, tp.Category) (tp.Category, error)
}

var categoryTableName = "category"

// ignores the id field in the incoming category and will generate a new one.
func (pg *Store) CreateCategory(cat tp.Category) (tp.Category, error) {
	cat.Id = uid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description) VALUES ($1, $2, $3)`, categoryTableName)
	_, err := pg.db.Exec(query, cat.Id.String(), cat.Title, cat.Description)
	if err != nil {
		return tp.Category{}, err
	}

	return cat, nil
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
		var cat tp.Category
		err = rows.Scan(&cat.Id, &cat.Title, &cat.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (pg *Store) GetCategory(title string) (tp.Category, error) {
	query := fmt.Sprintf(`SELECT id, title, description FROM %s WHERE title = $1`, categoryTableName)
	var cat tp.Category
	err := pg.db.QueryRow(query, title).Scan(&cat.Id, &cat.Title, &cat.Description)
	if err != nil {
		return tp.Category{}, err
	}

	return cat, nil
}

func (pg *Store) UpdateCategory(oldTitle string, cat tp.Category) (tp.Category, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, description = $2 WHERE title = $3 returning id`, categoryTableName)
	err := pg.db.QueryRow(query, cat.Title, cat.Description, oldTitle).Scan(&cat.Id)
	if err != nil {
		return tp.Category{}, err
	}

	return cat, nil
}
