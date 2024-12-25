package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var assetTableName = "asset"

func (pg *PostgresStore) CreateAsset(asset tp.Asset) (tp.Asset, error) {
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

func (pg *PostgresStore) DeleteAsset(groupTitle string, assetTitle string) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE title = $1 AND group_title = $2`, assetTableName)
	result, err := pg.db.Exec(query, assetTitle, groupTitle)
	if err != nil {
		return handleDbError(err, "asset")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "asset")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "asset with title %s and group title %s not found", assetTitle, groupTitle)
	}
	return nil
}

func (pg *PostgresStore) ListAssets() ([]tp.Asset, error) {
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

func (pg *PostgresStore) ListAssetsByGroup(groupTitle string) ([]tp.Asset, error) {
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

func (pg *PostgresStore) GetAsset(groupTitle string, assetTitle string) (tp.Asset, error) {
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

func (pg *PostgresStore) UpdateAsset(groupTitle string, assetTitle string, asset tp.Asset) (tp.Asset, error) {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET year = $1, 
		make = $2, 
		model_number = $3, 
		serial_number = $4, 
		description = $5, 
		category_title = $6, 
		title = $7 
		WHERE title = $8 AND group_title = $9 
		returning id`,
		assetTableName)

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
		groupTitle,
	).Scan(&asset.Id)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return asset, nil
}
