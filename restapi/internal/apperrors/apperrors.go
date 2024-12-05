package apperrors

// It is idiomative to use pointers in custom errors and have Error() return a pointer type
// unless a good reason arises, avoiding the pointers is a concious choice for reduced complexity.

type AppError struct {
	Code    string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

var CodeNotFound = "NOT_FOUND"
var CodeAlreadyExists = "ALREADY_EXISTS"
var CodeInvalid = "INVALID"

var ErrNotFound = AppError{Code: CodeNotFound, Message: "not found"}
var ErrAlreadyExists = AppError{Code: CodeAlreadyExists, Message: "already exists"}
var ErrInvalid = AppError{Code: CodeInvalid, Message: "invalid"}
var ErrIdMismatch = AppError{Code: CodeInvalid, Message: "id mismatch"}
var ErrGroupTitleMismatch = AppError{Code: CodeInvalid, Message: "group title mismatch"}
var ErrGroupTitleRequired = AppError{Code: CodeInvalid, Message: "group title required"}
var ErrCategoryTitleRequired = AppError{Code: CodeInvalid, Message: "category title mismatch"}
var ErrQuantityMustBePositive = AppError{Code: CodeInvalid, Message: "quantity must be positive"}
