package postgres

import (
	uid "github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Category interface {
	CreateCategory(string) (tp.UUID, error)
	DeleteCategory(tp.UUID) error
	ListCategory() ([]tp.Category, error)
	GetCategory(tp.UUID) (tp.Category, error)
	UpdateCategory(tp.UUID, string) error
}

func (pg *Store) CreateCategory(title string) (tp.UUID, error) {
	id := uid.New()
	query := `INSERT INTO category (id, title) VALUES ($1, $2) returning id`
	_, err := pg.db.Exec(query, id.String(), title)

	return id, err
}

func (pg *Store) DeleteCategory(id uid.UUID) error {
	query := `DELETE FROM category WHERE id = $1`
	_, err := pg.db.Exec(query, id.String())

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

func (pg *Store) GetCategory(id uid.UUID) (tp.Category, error) {
	query := `SELECT id, title FROM category WHERE id = $1`
	var cat tp.Category
	err := pg.db.QueryRow(query, id.String()).Scan(&cat.Id, &cat.Title)

	return cat, err
}

func (pg *Store) UpdateCategory(id uid.UUID, title string) error {
	query := `UPDATE category SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, title, id)

	return err
}
