package cmmsapp

import (
	"github.com/google/uuid"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// A usage trigger is an even that is triggered after a specific usage threshold has been reached
// since the last time a task was completed.  Like a generator being used for 30 hours, etc.

func (a *App) CreateUsageTrigger(groupTitle, assetTitle, taskId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	if err := a.validateAndInterpolateUsageTrigger(groupTitle, assetTitle, taskId, &usageTrigger); err != nil {
		return tp.UsageTrigger{}, err
	}

	return a.db.CreateUsageTrigger(usageTrigger)
}

func (a *App) DeleteUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string) error {
	ut, err := a.GetUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteUsageTrigger(ut.Id)
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
	tid, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId)
	if err != nil {
		return []tp.UsageTrigger{}, err
	}

	return a.db.ListUsageTriggersByTaskId(tid)
}

func (a *App) UpdateUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	ut, err := a.GetUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	if usageTrigger.Id != uuid.Nil && usageTrigger.Id != ut.Id {
		return tp.UsageTrigger{}, ae.ErrIdMismatch
	}

	if usageTrigger.TaskId == uuid.Nil {
		usageTrigger.TaskId = ut.TaskId
	}

	return a.db.UpdateUsageTrigger(ut.Id, usageTrigger)
}

func (a *App) validateAndInterpolateUsageTrigger(groupTitle, assetTitle, taskId string, usageTrigger *tp.UsageTrigger) error {
	tid, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId)
	if err != nil {
		return err
	}
	usageTrigger.TaskId = tid

	// validate usageunit
	if _, err := a.db.GetUsageUnit(usageTrigger.UsageUnit); err != nil {
		return err
	}

	// validate quantity
	if usageTrigger.Quantity <= 0 {
		return ae.ErrQuantityMustBePositive
	}

	return nil
}
