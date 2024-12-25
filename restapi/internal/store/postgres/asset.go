package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/pkg/errors"
)

var assetTableName = "asset"

func (pg *PostgresStore) CreateAsset(a tp.Asset) (tp.Asset, error) {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			id, title, year, make, model_number, serial_number, description
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)`, assetTableName)

	_, err := pg.db.Exec(query, a.Id, a.Title, a.Year, a.Make, a.ModelNumber, a.SerialNumber, a.Description)
	if err != nil {
		return tp.Asset{}, handleDbError(err, "asset")
	}

	return a, nil
}

func (pg *PostgresStore) DeleteAsset(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, assetTableName)
	result, err := pg.db.Exec(query, id)
	if err != nil {
		return handleDbError(err, "asset")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return handleDbError(err, "asset")
	}
	if rowsAffected == 0 {
		return errors.Wrapf(ae.ErrNotFound, "asset with id %s not found", id.String())
	}
	return nil
}

func (pg *PostgresStore) GetAsset(id uuid.UUID) (tp.Asset, error) {
	var asset tp.Asset
	query := fmt.Sprintf(`
		SELECT id, title, year, make, model_number, serial_number, description
		FROM %s 
		WHERE id = $1`, assetTableName)

	err := pg.db.QueryRow(query, id).Scan(
		&asset.Id,
		&asset.Title,
		&asset.Year,
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
		FROM %s`, assetTableName)
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

func (pg *PostgresStore) UpdateAsset(a tp.Asset) (tp.Asset, error) {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET title = $1,
		year = $2, 
		make = $3, 
		model_number = $4, 
		serial_number = $5, 
		description = $6, 
		WHERE id = $7`, assetTableName)

	result, err := pg.db.Exec(
		query,
		a.Title,
		a.Year,
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
