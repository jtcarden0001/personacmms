package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var assetTableName = "asset"
var catAssetTableName = "category_asset"
var gpAssetTableName = "agroup_asset"

func (pg *PostgresStore) AssociateAssetWithCategory(assetId, categoryId uuid.UUID) (tp.Asset, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (asset_id, category_id) 
			VALUES ($1, $2)`,
		catAssetTableName)

	_, err := pg.db.Exec(query, assetId, categoryId)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return pg.GetAsset(assetId)
}

func (pg *PostgresStore) AssociateAssetWithGroup(assetId, groupId uuid.UUID) (tp.Asset, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (asset_id, group_id) 
			VALUES ($1, $2)`,
		gpAssetTableName)

	_, err := pg.db.Exec(query, assetId, groupId)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return pg.GetAsset(assetId)
}

func (pg *PostgresStore) CreateAsset(a tp.Asset) (tp.Asset, error) {
	query := fmt.Sprintf(`
			INSERT INTO %s (id, title, year, manufacturer, make, model_number, serial_number, description) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		assetTableName)

	_, err := pg.db.Exec(query, a.Id, a.Title, a.Year, a.Manufacturer, a.Make, a.ModelNumber, a.SerialNumber, a.Description)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return a, nil
}

func (pg *PostgresStore) DeleteAsset(id uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE id = $1`,
		assetTableName)

	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "asset")
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return handleDbError(err, "asset")
	} else if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "asset with id %s not found", id.String())
	}

	return nil
}

func (pg *PostgresStore) DisassociateAssetWithCategory(assetId, categoryId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE asset_id = $1 AND category_id = $2`,
		catAssetTableName)

	result, err := pg.db.Exec(query, assetId, categoryId)
	if err != nil {
		return handleDbError(err, "asset")
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return handleDbError(err, "asset")
	} else if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "asset with id %s not found in category %s", assetId, categoryId)
	}

	return nil
}

func (pg *PostgresStore) DisassociateAssetWithGroup(assetId, groupId uuid.UUID) error {
	query := fmt.Sprintf(`
			DELETE FROM %s 
			WHERE asset_id = $1 AND group_id = $2`,
		gpAssetTableName)

	result, err := pg.db.Exec(query, assetId, groupId)
	if err != nil {
		return handleDbError(err, "asset")
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return handleDbError(err, "asset")
	} else if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "asset with id %s not found in group %s", assetId, groupId)
	}

	return nil
}

func (pg *PostgresStore) GetAsset(id uuid.UUID) (tp.Asset, error) {
	var asset tp.Asset
	query := fmt.Sprintf(`
			SELECT id, title, year, manufacturer, make, model_number, serial_number, description
			FROM %s 
			WHERE id = $1`,
		assetTableName)

	err := pg.db.QueryRow(query, id).Scan(
		&asset.Id,
		&asset.Title,
		&asset.Year,
		&asset.Manufacturer,
		&asset.Make,
		&asset.ModelNumber,
		&asset.SerialNumber,
		&asset.Description,
	)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return asset, nil
}

func (pg *PostgresStore) ListAssets() ([]tp.Asset, error) {
	query := fmt.Sprintf(`
			SELECT id, title, year, make, model_number, serial_number, description
			FROM %s`,
		assetTableName)

	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "asset")
	}

	assets := []tp.Asset{}
	for rows.Next() {
		var asset tp.Asset
		err = rows.Scan(&asset.Id, &asset.Title, &asset.Year, &asset.Make, &asset.ModelNumber, &asset.SerialNumber, &asset.Description)
		if err != nil {
			return nil, handleDbError(err, "asset")
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (pg *PostgresStore) ListAssetsByCategory(categoryId uuid.UUID) ([]tp.Asset, error) {
	query := fmt.Sprintf(`
			SELECT a.id, a.title, a.year, a.make, a.model_number, a.serial_number, a.description
			FROM %s a JOIN %s ac ON a.id = ac.asset_id
			WHERE ac.category_id = $1`,
		assetTableName, catAssetTableName)

	rows, err := pg.db.Query(query, categoryId)
	if err != nil {
		return nil, handleDbError(err, "asset")
	}

	assets := []tp.Asset{}
	for rows.Next() {
		var asset tp.Asset
		err = rows.Scan(&asset.Id, &asset.Title, &asset.Year, &asset.Make, &asset.ModelNumber, &asset.SerialNumber, &asset.Description)
		if err != nil {
			return nil, handleDbError(err, "asset")
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (pg *PostgresStore) ListAssetsByCategoryAndGroup(categoryId, groupId uuid.UUID) ([]tp.Asset, error) {
	query := fmt.Sprintf(`
			SELECT a.id, a.title, a.year, a.make, a.model_number, a.serial_number, a.description
			FROM %s a
			JOIN %s ac ON a.id = ac.asset_id
			JOIN %s ag ON a.id = ag.asset_id
			WHERE ac.category_id = $1 AND ag.group_id = $2`,
		assetTableName, catAssetTableName, gpAssetTableName)

	rows, err := pg.db.Query(query, categoryId, groupId)
	if err != nil {
		return nil, handleDbError(err, "asset")
	}

	assets := []tp.Asset{}
	for rows.Next() {
		var asset tp.Asset
		err = rows.Scan(&asset.Id, &asset.Title, &asset.Year, &asset.Make, &asset.ModelNumber, &asset.SerialNumber, &asset.Description)
		if err != nil {
			return nil, handleDbError(err, "asset")
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (pg *PostgresStore) ListAssetsByGroup(groupId uuid.UUID) ([]tp.Asset, error) {
	query := fmt.Sprintf(`
			SELECT a.id, a.title, a.year, a.make, a.model_number, a.serial_number, a.description
			FROM %s a JOIN %s ag ON a.id = ag.asset_id
			WHERE ag.group_id = $1`,
		assetTableName, gpAssetTableName)

	rows, err := pg.db.Query(query, groupId)
	if err != nil {
		return []tp.Asset{}, handleDbError(err, "asset")
	}

	assets := []tp.Asset{}
	for rows.Next() {
		var asset tp.Asset
		err = rows.Scan(&asset.Id, &asset.Title, &asset.Year, &asset.Make, &asset.ModelNumber, &asset.SerialNumber, &asset.Description)
		if err != nil {
			return []tp.Asset{}, handleDbError(err, "asset")
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

func (pg *PostgresStore) UpdateAsset(a tp.Asset) (tp.Asset, error) {
	query := fmt.Sprintf(`
			UPDATE %s 
			SET title = $1, year = $2, manufacturer = $3, make = $4, model_number = $5, serial_number = $6, description = $7
			WHERE id = $8`,
		assetTableName)

	result, err := pg.db.Exec(
		query,
		a.Title,
		a.Year,
		a.Manufacturer,
		a.Make,
		a.ModelNumber,
		a.SerialNumber,
		a.Description,
		a.Id,
	)

	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	if rowsAffected == 0 {
		return tp.Asset{}, errors.Wrapf(ae.ErrNotFound, "asset with id %s not found", a.Id)
	}

	return a, nil
}
