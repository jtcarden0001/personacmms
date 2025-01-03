package postgres

import (
	"database/sql"
	"strings"

	ae "github.com/jtcarden0001/personacmms/restapi/internal/utils/apperrors"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

const ERRPqDuplicateKey = "23505"

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
	switch errIn.Code {
	case ERRPqDuplicateKey:
		col := strings.Split(errIn.Detail, "(")[1]
		col = strings.Split(col, ")")[0]
		val := strings.Split(errIn.Detail, "(")[2]
		val = strings.Split(val, ")")[0]
		return errors.Wrapf(ae.ErrAlreadyExists, "an object with [%s]=[%s] already exists", col, val)
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
