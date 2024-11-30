package postgres

import (
	"fmt"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

var assetTaskToolTable = "assettask_tool"

func (pg *Store) CreateAssetTaskTool(tool tp.AssetTaskTool) (tp.AssetTaskTool, error) {
	query := fmt.Sprintf("INSERT INTO %s (assettask_id, tool_id) VALUES ($1, $2)", assetTaskToolTable)
	_, err := pg.db.Exec(query, tool.AssetTaskId, tool.ToolId)
	if err != nil {
		return tp.AssetTaskTool{}, handleDbError(err, "asset-task-tool")
	}

	return tool, nil
}

func (pg *Store) DeleteAssetTaskTool(atId, tId tp.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE assettask_id = $1 AND tool_id = $2", assetTaskToolTable)
	_, err := pg.db.Exec(query, atId, tId)
	if err != nil {
		return handleDbError(err, "asset-task-tool")
	}

	return nil
}

func (pg *Store) ListAssetTaskTools() ([]tp.AssetTaskTool, error) {
	query := fmt.Sprintf("SELECT assettask_id, tool_id FROM %s", assetTaskToolTable)
	rows, err := pg.db.Query(query)
	if err != nil {
		return nil, handleDbError(err, "asset-task-tool")
	}
	defer rows.Close()

	var at []tp.AssetTaskTool
	for rows.Next() {
		var e tp.AssetTaskTool
		err = rows.Scan(&e.AssetTaskId, &e.ToolId)
		if err != nil {
			return nil, handleDbError(err, "asset-task-tool")
		}
		at = append(at, e)
	}

	return at, nil
}

func (pg *Store) GetAssetTaskTool(atId, tId tp.UUID) (tp.AssetTaskTool, error) {
	query := fmt.Sprintf("SELECT assettask_id, tool_id FROM %s WHERE assettask_id = $1 AND tool_id = $2", assetTaskToolTable)
	var e tp.AssetTaskTool
	err := pg.db.QueryRow(query, atId, tId).Scan(&e.AssetTaskId, &e.ToolId)
	if err != nil {
		return tp.AssetTaskTool{}, handleDbError(err, "asset-task-tool")
	}

	return e, nil
}
