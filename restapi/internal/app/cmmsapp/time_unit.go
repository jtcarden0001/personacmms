package cmmsapp

import (
	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// Create a time unit
func (a *App) CreateTimeUnit(timeUnit tp.TimeUnit) (tp.TimeUnit, error) {
	return a.db.CreateTimeUnit(timeUnit)
}

// Delete a time unit
func (a *App) DeleteTimeUnit(timeUnitTitle string) error {
	// Get before delete so we can return a not found error.
	if _, err := a.GetTimeUnit(timeUnitTitle); err != nil {
		return err
	}

	return a.db.DeleteTimeUnit(timeUnitTitle)
}

// List all time units
func (a *App) ListTimeUnits() ([]tp.TimeUnit, error) {
	return a.db.ListTimeUnits()
}

// Get a time unit
func (a *App) GetTimeUnit(timeUnitTitle string) (tp.TimeUnit, error) {
	return a.db.GetTimeUnit(timeUnitTitle)
}

// Update a time unit
func (a *App) UpdateTimeUnit(timeUnitTitle string, timeUnit tp.TimeUnit) (tp.TimeUnit, error) {
	if timeUnit.Title == "" {
		return tp.TimeUnit{}, ae.ErrTimeUnitTitleRequired
	}

	return a.db.UpdateTimeUnit(timeUnitTitle, timeUnit)
}
