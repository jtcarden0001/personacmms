package postgres

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type Asset interface {
	CreateAsset(string, int, string, string, string, int) (int, error)
	DeleteAsset(int) error
	GetAllAsset() ([]tp.Asset, error)
	GetAsset(int) (tp.Asset, error)
	UpdateAsset(int, string, int, string, string, string, int) error
}

func (pg *Store) CreateAsset(title string, year int, make, modelNumber, description string, categoryId int) (int, error) {
	query := `INSERT INTO Asset (title, year, make, model_number, description, category_id) VALUES ($1, $2, $3, $4, $5, $6) returning id`
	var id int
	err := pg.db.QueryRow(query, title, year, make, modelNumber, description, categoryId).Scan(&id)

	return id, err
}

func (pg *Store) DeleteAsset(id int) error {
	query := `DELETE FROM Asset WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllAsset() ([]tp.Asset, error) {
	query := `SELECT id, title, year, make, model_number, description, category_id FROM Asset`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Asset []tp.Asset
	for rows.Next() {
		var e tp.Asset
		err = rows.Scan(&e.Id, &e.Title, &e.Year, &e.Make, &e.ModelNumber, &e.Description, &e.CategoryTitle)
		if err != nil {
			return nil, err
		}
		Asset = append(Asset, e)
	}

	return Asset, nil
}

func (pg *Store) GetAsset(id int) (tp.Asset, error) {
	query := `SELECT id, title, year, make, model_number, description, category_id FROM Asset WHERE id = $1`
	var e tp.Asset
	err := pg.db.QueryRow(query, id).Scan(&e.Id, &e.Title, &e.Year, &e.Make, &e.ModelNumber, &e.Description, &e.CategoryTitle)

	return e, err
}

func (pg *Store) UpdateAsset(id int, title string, year int, make, modelNumber, description string, categoryId int) error {
	query := `UPDATE Asset SET title = $1, year = $2, make = $3, model_number = $4, description = $6, category = $7 WHERE id = $8`
	_, err := pg.db.Exec(query, title, year, make, modelNumber, description, categoryId, id)

	return err
}
