package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

func (a *App) CreateDateTrigger(groupTitle string, assetTitle string, assetTaskId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	// TODO: implement this method
	return a.db.CreateDateTrigger(dateTrigger)
}

func (a *App) DeleteDateTrigger(groupTitle string, assetTitle string, assetTaskId string, dateTriggerId string) error {
	// TODO: implement this method
	parsedDtId, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteDateTrigger(parsedDtId)
}

func (a *App) ListDateTriggers(groupTitle string, assetTitle string, assetTaskId string) ([]tp.DateTrigger, error) {
	// TODO: implement this method
	dts, err := a.db.ListDateTriggers()
	if err != nil {
		return nil, err
	}

	// TODO: implement filtering
	return dts, nil
}

func (a *App) GetDateTrigger(groupTitle string, assetTitle string, assetTaskId string, dateTriggerId string) (tp.DateTrigger, error) {
	// TODO: implement this method
	parsedDtId, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return tp.DateTrigger{}, err
	}

	return a.db.GetDateTrigger(parsedDtId)
}

func (a *App) UpdateDateTrigger(groupTitle string, assetTitle string, assetTaskId string, dateTriggerId string, dateTrigger tp.DateTrigger) (tp.DateTrigger, error) {
	// TODO: implement this method
	parsedDtId, err := uuid.Parse(dateTriggerId)
	if err != nil {
		return tp.DateTrigger{}, err
	}

	return a.db.UpdateDateTrigger(parsedDtId, dateTrigger)
}
