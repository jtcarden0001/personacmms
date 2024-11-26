package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateAssetTaskTool(tool tp.AssetTaskTool) (tp.AssetTaskTool, error) {
	return a.db.CreateAssetTaskTool(tool)
}

func (a *App) CreateAssetTaskToolWithValidation(groupTitle, assetTitle, assetTaskId, toolId string) (tp.AssetTaskTool, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTaskTool{}, err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return tp.AssetTaskTool{}, err
	}

	return a.db.CreateAssetTaskTool(tp.AssetTaskTool{AssetTaskId: atId, ToolId: tId})
}

func (a *App) DeleteAssetTaskTool(groupTitle, assetTitle, assetTaskId, toolId string) error {
	// TODO: implement validation
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return err
	}

	return a.db.DeleteAssetTaskTool(atId, tId)
}

func (a *App) ListAssetTaskTools(groupTitle, assetTitle, assetTaskId string) ([]tp.AssetTaskTool, error) {
	// TODO: implement validation

	atTools, err := a.db.ListAssetTaskTools()
	if err != nil {
		return nil, err
	}

	// TODO: filter asset task tools by asset task id

	return atTools, nil
}

func (a *App) GetAssetTaskTool(groupTitle, assetTitle, assetTaskId, toolId string) (tp.AssetTaskTool, error) {
	// TODO: implement validation
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTaskTool{}, err
	}

	tId, err := uuid.Parse(toolId)
	if err != nil {
		return tp.AssetTaskTool{}, err
	}

	return a.db.GetAssetTaskTool(atId, tId)
}
