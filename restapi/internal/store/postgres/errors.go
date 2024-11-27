package postgres

import (
	"database/sql"

	ae "github.com/jtcarden0001/personacmms/restapi/internal/apperrors"
	"github.com/lib/pq"
)

func handleDbError(errIn error) error {
	if errIn == nil {
		return nil
	}

	if pqErr, ok := errIn.(*pq.Error); ok {
		return handlePqError(pqErr)
	}

	return handleSqlError(errIn)
}

// error code names come from https://github.com/lib/pq/blob/b7ffbd3b47da4290a4af2ccd253c74c2c22bfabf/error.go
func handlePqError(errIn *pq.Error) error {
	switch errIn.Code.Name {
	default:
		return errIn
	}
}

func handleSqlError(errIn error) error {
	switch errIn {
	case sql.ErrNoRows:
		return ae.ErrNotFound
	default:
		return errIn
	}
}
