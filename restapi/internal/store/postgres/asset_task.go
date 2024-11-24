package postgres

import (
	"fmt"

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
	query := fmt.Sprintf(`INSERT INTO %s (id, title, unique_instructions, asset_id, task_id) VALUES ($1, $2, $3, $4, $5)`, assetTaskTable)
	_, err = pg.db.Exec(query, at.Id, at.Title, at.UniqueInstructions, at.AssetId, at.TaskId)
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

	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, assetTaskTable)
	_, err = pg.db.Exec(query, atId)
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

	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, task_id FROM %s`, assetTaskTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var at []tp.AssetTask
	for rows.Next() {
		var e tp.AssetTask
		err = rows.Scan(&e.Id, &e.Title, &e.UniqueInstructions, &e.AssetId, &e.TaskId)
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

	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, task_id FROM %s WHERE id = $1`, assetTaskTable)
	var e tp.AssetTask
	err = pg.db.QueryRow(query, atId).Scan(&e.Id, &e.Title, &e.UniqueInstructions, &e.AssetId, &e.TaskId)
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

	query := fmt.Sprintf(`UPDATE %s SET title = $1, unique_instructions = $2, asset_id = $3, task_id = $4 WHERE id = $5 returning id`, assetTaskTable)
	err = pg.db.QueryRow(query, at.Title, at.UniqueInstructions, at.AssetId, at.TaskId, atId).Scan(&at.Id)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return at, nil
}
