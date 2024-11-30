package postgres

import (
	"fmt"

	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var assetTaskTable = "asset_task"

func (pg *Store) CreateAssetTask(at tp.AssetTask) (tp.AssetTask, error) {
	at.Id = uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (id, title, unique_instructions, asset_id, tasktemplate_id) VALUES ($1, $2, $3, $4, $5)`, assetTaskTable)
	_, err := pg.db.Exec(query, at.Id, at.Title, at.UniqueInstructions, at.AssetId, at.TaskTemplateId)
	if err != nil {
		return tp.AssetTask{}, handleDbError(err, "asset-task")
	}

	return at, nil
}

func (pg *Store) DeleteAssetTask(atId tp.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, assetTaskTable)
	_, err := pg.db.Exec(query, atId)
	if err != nil {
		return handleDbError(err, "asset-task")
	}

	return nil
}

func (pg *Store) ListAssetTasks() ([]tp.AssetTask, error) {
	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, tasktemplate_id FROM %s`, assetTaskTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "asset-task")
	}
	defer rows.Close()

	var at []tp.AssetTask
	for rows.Next() {
		var e tp.AssetTask
		err = rows.Scan(&e.Id, &e.Title, &e.UniqueInstructions, &e.AssetId, &e.TaskTemplateId)
		if err != nil {
			return nil, handleDbError(err, "asset-task")
		}
		at = append(at, e)
	}

	return at, nil
}

func (pg *Store) GetAssetTask(atId tp.UUID) (tp.AssetTask, error) {
	query := fmt.Sprintf(`SELECT id, title, unique_instructions, asset_id, tasktemplate_id FROM %s WHERE id = $1`, assetTaskTable)
	var e tp.AssetTask
	err := pg.db.QueryRow(query, atId).Scan(&e.Id, &e.Title, &e.UniqueInstructions, &e.AssetId, &e.TaskTemplateId)
	if err != nil {
		return tp.AssetTask{}, handleDbError(err, "asset-task")
	}

	return e, nil
}

func (pg *Store) UpdateAssetTask(atId tp.UUID, at tp.AssetTask) (tp.AssetTask, error) {
	query := fmt.Sprintf(`UPDATE %s SET title = $1, unique_instructions = $2, asset_id = $3, tasktemplate_id = $4 WHERE id = $5 returning id`, assetTaskTable)
	err := pg.db.QueryRow(query, at.Title, at.UniqueInstructions, at.AssetId, at.TaskTemplateId, atId).Scan(&at.Id)
	if err != nil {
		return tp.AssetTask{}, handleDbError(err, "asset-task")
	}

	return at, nil
}
