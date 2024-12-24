package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) ListUsageUnits() ([]tp.UsageUnit, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListUsageUnits not implemented")
}
