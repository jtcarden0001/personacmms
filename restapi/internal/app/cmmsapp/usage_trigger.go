package cmmsapp

import (
	"github.com/google/uuid"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

// Create a UsageTrigger
func (a *App) CreateUsageTrigger(groupTitle, assetTitle, taskId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	if err := a.validateAndInterpolateUsageTrigger(groupTitle, assetTitle, taskId, &usageTrigger); err != nil {
		return tp.UsageTrigger{}, err
	}

	return a.db.CreateUsageTrigger(usageTrigger)
}

// Delete a UsageTrigger
func (a *App) DeleteUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string) error {
	// Get before delete to provide opportunity to return not found error
	ut, err := a.GetUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId)
	if err != nil {
		return err
	}

	return a.db.DeleteUsageTrigger(ut.Id)
}

// Get a usage trigger that is essentially namespaced under the task specificed
func (a *App) GetUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string) (tp.UsageTrigger, error) {
	if _, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId); err != nil {
		return tp.UsageTrigger{}, err
	}

	parsedutId, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	return a.db.GetUsageTrigger(parsedutId)
}

// List all usage triggers for a particular task
func (a *App) ListUsageTriggers(groupTitle, assetTitle, taskId string) ([]tp.UsageTrigger, error) {
	tid, err := a.validateTriggerDependencies(groupTitle, assetTitle, taskId)
	if err != nil {
		return []tp.UsageTrigger{}, err
	}

	return a.db.ListUsageTriggersByTaskId(tid)
}

// Update a usage trigger
func (a *App) UpdateUsageTrigger(groupTitle, assetTitle, taskId, usageTriggerId string, usageTrigger tp.UsageTrigger) (tp.UsageTrigger, error) {
	err := a.validateAndInterpolateUsageTrigger(groupTitle, assetTitle, taskId, &usageTrigger)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	utId, err := uuid.Parse(usageTriggerId)
	if err != nil {
		return tp.UsageTrigger{}, err
	}

	if usageTrigger.Id == uuid.Nil {
		usageTrigger.Id = utId
	} else if usageTrigger.Id != utId {
		return tp.UsageTrigger{}, ae.ErrIdMismatch
	}

	return a.db.UpdateUsageTrigger(utId, usageTrigger)
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
