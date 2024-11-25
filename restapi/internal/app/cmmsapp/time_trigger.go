package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateTimeTrigger(groupId string, assetId string, taskId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	// TODO: implement validations
	return a.db.CreateTimeTrigger(timeTrigger)
}

func (a *App) DeleteTimeTrigger(groupId string, assetId string, taskId string, timeTriggerId string) error {
	// TODO: implement validations
	ttId, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteTimeTrigger(ttId)
}

func (a *App) ListTimeTriggers(groupId string, assetId string, taskId string) ([]tp.TimeTrigger, error) {
	// TODO: implement validations
	ttgs, err := a.db.ListTimeTriggers()
	if err != nil {
		return []tp.TimeTrigger{}, err
	}

	// TODO: implement filtering
	return ttgs, nil
}

func (a *App) GetTimeTrigger(groupId string, assetId string, taskId string, timeTriggerId string) (tp.TimeTrigger, error) {
	// TODO: implement validations
	ttId, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return tp.TimeTrigger{}, err
	}
	return a.db.GetTimeTrigger(ttId)
}

func (a *App) UpdateTimeTrigger(groupId string, assetId string, taskId string, timeTriggerId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	// TODO: implement validations
	ttId, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return tp.TimeTrigger{}, err
	}
	return a.db.UpdateTimeTrigger(ttId, timeTrigger)
}
