package cmmsapp

import (
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// CreateUsageUnit creates a usage unit
func (a *App) CreateUsageUnit(usageUnit tp.UsageUnit) (tp.UsageUnit, error) {
	if err := a.validateUsageUnit(usageUnit); err != nil {
		return tp.UsageUnit{}, err
	}

	return a.db.CreateUsageUnit(usageUnit)
}

// DeleteUsageUnit deletes a usage unit
func (a *App) DeleteUsageUnit(usageUnitTitle string) error {
	// Get before delete so we can return a not found error.
	if _, err := a.GetUsageUnit(usageUnitTitle); err != nil {
		return err
	}

	return a.db.DeleteUsageUnit(usageUnitTitle)
}

// ListUsageUnits lists all usage units
func (a *App) ListUsageUnits() ([]tp.UsageUnit, error) {
	return a.db.ListUsageUnits()
}

// GetUsageUnit gets a usage unit
func (a *App) GetUsageUnit(usageUnitTitle string) (tp.UsageUnit, error) {
	return a.db.GetUsageUnit(usageUnitTitle)
}

// UpdateUsageUnit updates a usage unit
func (a *App) UpdateUsageUnit(usageUnitTitle string, usageUnit tp.UsageUnit) (tp.UsageUnit, error) {
	if err := a.validateUsageUnit(usageUnit); err != nil {
		return tp.UsageUnit{}, err
	}

	return a.db.UpdateUsageUnit(usageUnitTitle, usageUnit)
}

func (a *App) validateUsageUnit(usageUnit tp.UsageUnit) error {
	if usageUnit.Title == "" {
		return ae.ErrUsageUnitTitleRequired
	}

	if !a.isValidUsageUnitType(usageUnit.Type) {
		return ae.ErrUsageUnitTypeInvalid
	}

	return nil
}

func (a *App) isValidUsageUnitType(usageUnitType string) bool {
	return tp.ValidUsageUnitTypes[usageUnitType]
}
