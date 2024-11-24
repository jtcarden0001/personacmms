package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateAssetTask(groupTitle string, assetTitle string, at tp.AssetTask) (tp.AssetTask, error) {
	// TODO: validate and populate asset task

	return a.db.CreateAssetTask(at)
}

func (a *App) DeleteAssetTask(groupTitle string, assetTitle string, assetTaskId string) error {
	// TODO: validate asset task

	// cast assetTaskId to tp.UUID
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return err
	}

	return a.db.DeleteAssetTask(atId)
}

func (a *App) ListAssetTasks(groupTitle string, assetTitle string) ([]tp.AssetTask, error) {
	ats, err := a.db.ListAssetTasks()
	if err != nil {
		return []tp.AssetTask{}, err
	}

	// filter asset tasks by group and asset title
	var assetTasks []tp.AssetTask
	for _, at := range ats {
		// if at.GroupTitle == groupTitle && at.AssetTitle == assetTitle {
		assetTasks = append(assetTasks, at)
		// }
	}

	return assetTasks, nil
}

func (a *App) GetAssetTask(groupTitle string, assetTitle string, assetTaskId string) (tp.AssetTask, error) {
	// TODO: validate asset task

	// cast assetTaskId to tp.UUID
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return a.db.GetAssetTask(atId)
}

func (a *App) UpdateAssetTask(groupTitle string, assetTitle string, assetTaskId string, at tp.AssetTask) (tp.AssetTask, error) {
	// TODO: validate asset task

	// cast assetTaskId to tp.UUID
	atId, err := uuid.Parse(assetTaskId)
	if err != nil {
		return tp.AssetTask{}, err
	}

	return a.db.UpdateAssetTask(atId, at)
}
