package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateUsageTrigger(groupTitle, assetTitle, taskId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	// TODO: implement this method
	return a.db.CreateUsageTrigger(usageTrigger)
}

func (a *App) DeleteUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string) error {
	// TODO: implement this method
	parsedutId, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteUsageTrigger(parsedutId)
}

func (a *App) GetUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string) (tp.UsageTrigger, error) {
	// TODO: implement this method
	parsedutId, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	return a.db.GetUsageTrigger(parsedutId)
}

func (a *App) ListUsageTriggers(groupTitle, assetTitle, taskId string) ([]tp.UsageTrigger, error) {
	// TODO: implement this method
	uts, err := a.db.ListUsageTriggers()
	if err != nil {
		return nil, err
	}
	// TODO: filter
	return uts, nil
}

func (a *App) UpdateUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	// TODO: implement this method
	parsedutId, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	return a.db.UpdateUsageTrigger(parsedutId, usageTrigger)
}
