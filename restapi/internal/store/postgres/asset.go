package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var assetTableName = "asset"

func (pg *Store) CreateAsset(asset tp.Asset) (tp.Asset, error) {
	asset.Id = uuid.New()
	// TODO: make this line length more tenable
	query := fmt.Sprintf(`
		INSERT INTO %s (
			id, title, group_title, year, make, model_number, serial_number, description, category_title
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)`, assetTableName)

	_, err := pg.db.Exec(query, asset.Id, asset.Title, asset.GroupTitle, asset.Year, asset.Make, asset.ModelNumber, asset.SerialNumber, asset.Description, asset.CategoryTitle)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return asset, nil
}

func (pg *Store) DeleteAsset(groupTitle string, assetTitle string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1 AND group_title = $2`, assetTableName)
	_, err := pg.db.Exec(query, assetTitle, groupTitle)
	return handleDbError(err, "asset")
}

func (pg *Store) ListAssets() ([]tp.Asset, error) {
	query := fmt.Sprintf(`
		SELECT group_title, title, id, year, make, model_number, serial_number, description, category_title 
		FROM %s`, assetTableName)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "asset")
	}

	assets := []tp.Asset{}
	for rows.Next() {
		var asset tp.Asset
		err = rows.Scan(&asset.GroupTitle, &asset.Title, &asset.Id, &asset.Year, &asset.Make, &asset.ModelNumber, &asset.SerialNumber, &asset.Description, &asset.CategoryTitle)
		if err != nil {
			return nil, handleDbError(err, "asset")
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (pg *Store) ListAssetsByGroup(groupTitle string) ([]tp.Asset, error) {
	query := fmt.Sprintf(`
		SELECT group_title, title, id, year, make, model_number, serial_number, description, category_title 
		FROM %s 
		WHERE group_title = $1`, assetTableName)
	rows, err := pg.db.Query(query, groupTitle)
	if err != nil {
		return nil, handleDbError(err, "asset")
	}

	assets := []tp.Asset{}
	for rows.Next() {
		var asset tp.Asset
		err = rows.Scan(&asset.GroupTitle, &asset.Title, &asset.Id, &asset.Year, &asset.Make, &asset.ModelNumber, &asset.SerialNumber, &asset.Description, &asset.CategoryTitle)
		if err != nil {
			return nil, handleDbError(err, "asset")
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (pg *Store) GetAsset(groupTitle string, assetTitle string) (tp.Asset, error) {
	var asset tp.Asset
	query := fmt.Sprintf(`
		SELECT group_title, title, id, year, make, model_number, serial_number, description, category_title 
		FROM %s 
		WHERE title = $1 AND group_title = $2`, assetTableName)

	err := pg.db.QueryRow(query, assetTitle, groupTitle).Scan(
		&asset.GroupTitle,
		&asset.Title,
		&asset.Id,
		&asset.Year,
		&asset.Make,
		&asset.ModelNumber,
		&asset.SerialNumber,
		&asset.Description,
		&asset.CategoryTitle,
	)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return asset, nil
}

func (pg *Store) UpdateAsset(groupTitle string, assetTitle string, asset tp.Asset) (tp.Asset, error) {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET year = $1, make = $2, model_number = $3, serial_number = $4, description = $5, category_title = $6, title = $7 
		WHERE title = $8 AND group_title = $9 returning id`, assetTableName)

	err := pg.db.QueryRow(
		query,
		asset.Year,
		asset.Make,
		asset.ModelNumber,
		asset.SerialNumber,
		asset.Description,
		asset.CategoryTitle,
		asset.Title,
		assetTitle,
		asset.GroupTitle,
	).Scan(&asset.Id)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return asset, nil
}
