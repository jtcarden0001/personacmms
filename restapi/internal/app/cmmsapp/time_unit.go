package cmmsapp

import (
	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
)

func (a *App) ListTimeUnits() ([]tp.TimeUnit, error) {
	return nil, ae.New(ae.CodeNotImplemented, "ListTimeUnits not implemented")
}
