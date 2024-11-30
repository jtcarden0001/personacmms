package postgres

import (
	"database/sql"

	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func handleDbError(errIn error, resourceType string) error {
	if errIn == nil {
		return nil
	}

	if pqErr, ok := errIn.(*pq.Error); ok {
		return handlePqError(pqErr, resourceType)
	}

	return handleSqlError(errIn, resourceType)
}

// error code names come from https://github.com/lib/pq/blob/b7ffbd3b47da4290a4af2ccd253c74c2c22bfabf/error.go
func handlePqError(errIn *pq.Error, resourceType string) error {
	switch errIn.Code.Name {
	default:
		return errors.Wrapf(errIn, "unexepected pq error occurred for %s", resourceType)
	}
}

func handleSqlError(errIn error, resourceType string) error {
	switch errIn {
	case sql.ErrNoRows:
		return errors.Wrapf(ae.ErrNotFound, "the %s does not exist", resourceType)
	default:
		return errors.Wrapf(errIn, "unexepected sql error occurred for %s", resourceType)
	}
}
