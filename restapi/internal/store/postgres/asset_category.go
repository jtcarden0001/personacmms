package postgres

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type AssetCategory interface {
	CreateAssetCategory(string) (int, error)
	DeleteAssetCategory(int) error
	GetAllAssetCategory() ([]tp.AssetCategory, error)
	GetAssetCategory(int) (tp.AssetCategory, error)
	UpdateAssetCategory(int, string) error
}

func (pg *Store) CreateAssetCategory(title string) (int, error) {
	query := `INSERT INTO asset_category (title) VALUES ($1) returning id`
	var id int
	err := pg.db.QueryRow(query, title).Scan(&id)

	return id, err
}

func (pg *Store) DeleteAssetCategory(id int) error {
	query := `DELETE FROM asset_category WHERE id = $1`
	_, err := pg.db.Exec(query, id)

	return err
}

func (pg *Store) GetAllAssetCategory() ([]tp.AssetCategory, error) {
	query := `SELECT id, title FROM asset_category`
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []tp.AssetCategory
	for rows.Next() {
		var ec tp.AssetCategory
		err = rows.Scan(&ec.Id, &ec.Title)
		if err != nil {
			return nil, err
		}
		categories = append(categories, ec)
	}

	return categories, nil
}

func (pg *Store) GetAssetCategory(id int) (tp.AssetCategory, error) {
	query := `SELECT id, title FROM asset_category WHERE id = $1`
	var ec tp.AssetCategory
	err := pg.db.QueryRow(query, id).Scan(&ec.Id, &ec.Title)

	return ec, err
}

func (pg *Store) UpdateAssetCategory(id int, title string) error {
	query := `UPDATE asset_category SET title = $1 WHERE id = $2`
	_, err := pg.db.Exec(query, title, id)

	return err
}
