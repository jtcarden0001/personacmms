package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

type AssetTask interface {
	CreateAssetTask(string, string, tp.AssetTask) (tp.AssetTask, error)
	DeleteAssetTask(string, string, string) error
	ListAssetTasks(string, string) ([]tp.AssetTask, error)
	GetAssetTask(string, string, string) (tp.AssetTask, error)
	UpdateAssetTask(string, string, string, tp.AssetTask) (tp.AssetTask, error)
}

func (a *App) CreateAssetTask(groupTitle string, assetTitle string, at tp.AssetTask) (tp.AssetTask, error) {
	return a.db.CreateAssetTask(groupTitle, assetTitle, at)
}

func (a *App) DeleteAssetTask(groupTitle string, assetTitle string, assetTaskId string) error {
	// cast assetTaskId to tp.UUID
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return err
	}

	return a.db.DeleteAssetTask(groupTitle, assetTitle, atId)
}

func (a *App) ListAssetTasks(groupTitle string, assetTitle string) ([]tp.AssetTask, error) {
	return a.db.ListAssetTasks(groupTitle, assetTitle)
}

func (a *App) GetAssetTask(groupTitle string, assetTitle string, assetTaskId string) (tp.AssetTask, error) {
	// cast assetTaskId to tp.UUID
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return a.db.GetAssetTask(groupTitle, assetTitle, atId)
}

func (a *App) UpdateAssetTask(groupTitle string, assetTitle string, assetTaskId string, at tp.AssetTask) (tp.AssetTask, error) {
	// cast assetTaskId to tp.UUID
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return a.db.UpdateAssetTask(groupTitle, assetTitle, atId, at)
}
