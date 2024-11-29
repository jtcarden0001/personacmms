package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var assetTaskConsumableTable = "assettask_consumable"

func (pg *Store) CreateAssetTaskConsumable(consumable tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error) {
	query := fmt.Sprintf("INSERT INTO %s (assettask_id, consumable_id, quantity_note) VALUES ($1, $2, $3)", assetTaskConsumableTable)
	_, err := pg.db.Exec(query, consumable.AssetTaskId, consumable.ConsumableId, consumable.QuantityNote)
	if err != nil {
		return tp.AssetTaskConsumable{}, handleDbError(err)
	}

	return consumable, nil
}

func (pg *Store) DeleteAssetTaskConsumable(atId, cId tp.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE assettask_id = $1 AND consumable_id = $2", assetTaskConsumableTable)
	_, err := pg.db.Exec(query, atId, cId)
	if err != nil {
		return handleDbError(err)
	}

	return nil
}

func (pg *Store) ListAssetTaskConsumables() ([]tp.AssetTaskConsumable, error) {
	query := fmt.Sprintf("SELECT assettask_id, consumable_id, quantity_note FROM %s", assetTaskConsumableTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err)
	}
	defer rows.Close()

	var at []tp.AssetTaskConsumable
	for rows.Next() {
		var e tp.AssetTaskConsumable
		err = rows.Scan(&e.AssetTaskId, &e.ConsumableId, &e.QuantityNote)
		if err != nil {
			return nil, handleDbError(err)
		}
		at = append(at, e)
	}

	return at, nil
}

func (pg *Store) GetAssetTaskConsumable(atId, cId tp.UUID) (tp.AssetTaskConsumable, error) {
	query := fmt.Sprintf("SELECT assettask_id, consumable_id, quantity_note FROM %s WHERE assettask_id = $1 AND consumable_id = $2", assetTaskConsumableTable)
	var e tp.AssetTaskConsumable
	err := pg.db.QueryRow(query, atId, cId).Scan(&e.AssetTaskId, &e.ConsumableId, &e.QuantityNote)
	if err != nil {
		return tp.AssetTaskConsumable{}, handleDbError(err)
	}

	return e, nil
}

func (pg *Store) UpdateAssetTaskConsumable(atc tp.AssetTaskConsumable) (tp.AssetTaskConsumable, error) {
	query := fmt.Sprintf("UPDATE %s SET quantity_note = $1 WHERE assettask_id = $2 AND consumable_id = $3", assetTaskConsumableTable)
	_, err := pg.db.Exec(query, atc.QuantityNote, atc.AssetTaskId, atc.ConsumableId)
	if err != nil {
		return tp.AssetTaskConsumable{}, handleDbError(err)
	}

	return atc, nil
}
