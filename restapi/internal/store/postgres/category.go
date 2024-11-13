package postgres

import (
	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Category interface {
	CreateCategory(string, string) (tp.Category, error)
	DeleteCategory(string) error
	ListCategory() ([]tp.Category, error)
	GetCategory(string) (tp.Category, error)
	UpdateCategory(string, string, string) (tp.Category, error)
}

func (pg *Store) CreateCategory(title, description string) (tp.Category, error) {
	id := uid.New()
	query := `INSERT INTO category (id, title, description) VALUES ($1, $2, $3) returning id`
	_, err := pg.db.Exec(query, id.String(), title)
	if err != nil {
		return tp.Category{}, err
	}

	category := tp.Category{
		Id:    id,
		Title: title,
	}

	return category, nil
}

func (pg *Store) DeleteCategory(title string) error {
	query := `DELETE FROM category WHERE title = $1`
	_, err := pg.db.Exec(query, title)

	return err
}

func (pg *Store) ListCategory() ([]tp.Category, error) {
	query := `SELECT id, title FROM category`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []tp.Category
	for rows.Next() {
		var cat tp.Category
		err = rows.Scan(&cat.Id, &cat.Title)
		if err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (pg *Store) GetCategory(title string) (tp.Category, error) {
	query := `SELECT id, title FROM category WHERE title = $1`
	var cat tp.Category
	err := pg.db.QueryRow(query, title).Scan(&cat.Id, &cat.Title)
	if err != nil {
		return tp.Category{}, err
	}

	return cat, nil
}

func (pg *Store) UpdateCategory(oldTitle, newTitle, newDescription string) (tp.Category, error) {
	query := `UPDATE category SET title = $1, description = $2 WHERE title = $2`
	_, err := pg.db.Exec(query, newTitle, newDescription, oldTitle)
	if err != nil {
		return tp.Category{}, err
	}

	return pg.GetCategory(newTitle)
}
