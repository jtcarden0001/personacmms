package cmmsapp

import (
	"github.com/google/uuid"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// a time trigger is an event that is triggered after a specific time has elapsed since the last time a
// work order was completed for a task.

// create a time trigger
func (a *App) CreateTimeTrigger(groupTitle string, assetTitle string, taskId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	if err := a.validateAndInterpolateTimeTrigger(groupTitle, assetTitle, taskId, &timeTrigger); err != nil {
		return tp.TimeTrigger{}, err
	}

	return a.db.CreateTimeTrigger(timeTrigger)
}

// delete a time trigger
func (a *App) DeleteTimeTrigger(groupTitle string, assetTitle string, taskId string, timeTriggerId string) error {
	tt, err := a.GetTimeTrigger(groupTitle, assetTitle, taskId, timeTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteTimeTrigger(tt.Id)
}

// list all time triggers for a particular task
func (a *App) ListTimeTriggers(groupTitle string, assetTitle string, taskId string) ([]tp.TimeTrigger, error) {
	tid, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId)
	if err != nil {
		return []tp.TimeTrigger{}, err
	}

	// TODO: implement this on the DB level
	return a.db.ListTimeTriggersByTaskId(tid)
}

// get a time trigger
func (a *App) GetTimeTrigger(groupTitle string, assetTitle string, taskId string, timeTriggerId string) (tp.TimeTrigger, error) {
	if _, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId); err != nil {
		return tp.TimeTrigger{}, err
	}

	parsedTtId, err := uuid.Parse(timeTriggerId)
	if err != nil {
		return tp.TimeTrigger{}, err
	}

	return a.db.GetTimeTrigger(parsedTtId)
}

// update a time trigger
func (a *App) UpdateTimeTrigger(groupTitle string, assetTitle string, taskId string, timeTriggerId string, timeTrigger tp.TimeTrigger) (tp.TimeTrigger, error) {
	tt, err := a.GetTimeTrigger(groupTitle, assetTitle, taskId, timeTriggerId)
	if err != nil {
		return tp.TimeTrigger{}, err
	}

	if timeTrigger.Id != uuid.Nil && timeTrigger.Id != tt.Id {
		return tp.TimeTrigger{}, ae.ErrIdMismatch
	}

	if timeTrigger.TaskId == uuid.Nil {
		timeTrigger.TaskId = tt.TaskId
	}

	return a.db.UpdateTimeTrigger(tt.Id, timeTrigger)
}

func (a *App) validateAndInterpolateTimeTrigger(groupTitle string, assetTitle string, taskId string, dateTrigger *tp.TimeTrigger) error {
	// essentially validate namespace exists
	tid, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId)
	if err != nil {
		return err
	}
	dateTrigger.TaskId = tid

	// validate TimeUnit
	if _, err := a.GetTimeUnit(dateTrigger.TimeUnit); err != nil {
		return err
	}

	// validate Quantity
	if dateTrigger.Quantity <= 0 {
		return ae.ErrQuantityMustBePositive
	}

	return nil
}
