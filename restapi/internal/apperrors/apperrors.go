package apperrors

import (
	"fmt"
	"strings"

	tp "github.com/jtcarden0001/personacmms/restapi/internal/types"
)

// It is idiomatic to use pointers in custom errors and have Error() return a pointer type
// unless a good reason arises, avoiding the pointers is a concious choice for reduced complexity.

type AppError struct {
	Code    string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

// Already Exists Errors
var CodeAlreadyExists = "ALREADY_EXISTS"
var ErrAlreadyExists = AppError{Code: CodeAlreadyExists, Message: "already exists"}

// Invalid Errors
var CodeInvalid = "INVALID"
var ErrCategoryTitleRequired = AppError{Code: CodeInvalid, Message: "category title mismatch"}
var ErrGroupTitleMismatch = AppError{Code: CodeInvalid, Message: "group title mismatch"}
var ErrGroupTitleRequired = AppError{Code: CodeInvalid, Message: "group title required"}
var ErrIdMismatch = AppError{Code: CodeInvalid, Message: "id mismatch"}
var ErrInvalid = AppError{Code: CodeInvalid, Message: "invalid"}
var ErrQuantityMustBePositive = AppError{Code: CodeInvalid, Message: "quantity must be positive"}
var ErrTimeUnitTitleRequired = AppError{Code: CodeInvalid, Message: "time unit title required"}
var ErrToolTitleRequired = AppError{Code: CodeInvalid, Message: "tool title required"}
var ErrUsageUnitTitleRequired = AppError{Code: CodeInvalid, Message: "usage unit title required"}
var ErrUsageUnitTypeInvalid = AppError{Code: CodeInvalid, Message: createInvalidUsageUnitTypeMessage()}
var ErrWorkOrderStatusTitleRequired = AppError{Code: CodeInvalid, Message: "work order status title required"}

// Not Found Errors
var CodeNotFound = "NOT_FOUND"
var ErrNotFound = AppError{Code: CodeNotFound, Message: "not found"}

// dynamic error message generation functions
func createInvalidUsageUnitTypeMessage() string {
	prefix := "invalid usage unit type, must be one of:"
	keys := make([]string, 0, len(tp.ValidUsageUnitTypes))
	for key := range tp.ValidUsageUnitTypes {
		keys = append(keys, key)
	}
	commaDelimitedString := strings.Join(keys, ", ")
	return fmt.Sprintf("%s %s", prefix, commaDelimitedString)
}
