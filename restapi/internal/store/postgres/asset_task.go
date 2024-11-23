package postgres

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type AssetTask interface {
	CreateAssetTask(string, string, tp.AssetTask) (tp.AssetTask, error)
	DeleteAssetTask(string, string, tp.UUID) error
	ListAssetTasks(string, string) ([]tp.AssetTask, error)
	GetAssetTask(string, string, tp.UUID) (tp.AssetTask, error)
	UpdateAssetTask(string, string, tp.UUID, tp.AssetTask) (tp.AssetTask, error)
}

var assetTaskTable = "asset_task"

func (pg *Store) CreateAssetTask(groupTitle string, assetTitle string, at tp.AssetTask) (tp.AssetTask, error) {
	err := pg.validateAsset(groupTitle, assetTitle)
	if err != nil {
		return tp.AssetTask{}, err
	}

	at.Id = uuid.New()
	query := `INSERT INTO $1 (id, unique_instructions, asset_id, unique_instructions, task_id) VALUES ($2, $3, $4, $5, $6)`
	_, err = pg.db.Exec(query, assetTaskTable, at.Id, at.UniqueInstructions, at.AssetId, at.TaskId)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return at, nil
}

func (pg *Store) DeleteAssetTask(groupTitle string, assetTitle string, atId tp.UUID) error {
	err := pg.validateAsset(groupTitle, assetTitle)
	if err != nil {
		return err
	}

	query := `DELETE FROM $1 WHERE id = $2`
	_, err = pg.db.Exec(query, assetTaskTable, atId)
	if err != nil {
		return err
	}

	return nil
}

func (pg *Store) ListAssetTasks(groupTitle string, assetTitle string) ([]tp.AssetTask, error) {
	err := pg.validateAsset(groupTitle, assetTitle)
	if err != nil {
		return nil, err
	}

	query := `SELECT id, unique_instructions, asset_id, task_id FROM $1`
	rows, err := pg.db.Query(query, assetTaskTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var at []tp.AssetTask
	for rows.Next() {
		var e tp.AssetTask
		err = rows.Scan(&e.Id, &e.UniqueInstructions, &e.AssetId, &e.TaskId)
		if err != nil {
			return nil, err
		}
		at = append(at, e)
	}

	return at, nil
}

func (pg *Store) GetAssetTask(groupTitle string, assetTitle string, atId tp.UUID) (tp.AssetTask, error) {
	err := pg.validateAsset(groupTitle, assetTitle)
	if err != nil {
		return tp.AssetTask{}, err
	}

	query := `SELECT id, unique_instructions, asset_id, task_id FROM $1 WHERE id = $2`
	var e tp.AssetTask
	err = pg.db.QueryRow(query, assetTaskTable, atId).Scan(&e.Id, &e.UniqueInstructions, &e.AssetId, &e.TaskId)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return e, nil
}

func (pg *Store) UpdateAssetTask(groupTitle string, assetTitle string, atId tp.UUID, at tp.AssetTask) (tp.AssetTask, error) {
	err := pg.validateAsset(groupTitle, assetTitle)
	if err != nil {
		return tp.AssetTask{}, err
	}

	query := `UPDATE $1 SET unique_instructions = $2, asset_id = $3, task_id = $4 WHERE id = $5`
	_, err = pg.db.Exec(query, assetTaskTable, at.UniqueInstructions, at.AssetId, at.TaskId, atId)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return at, nil
}
