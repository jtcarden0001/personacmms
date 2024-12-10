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
var ErrWorkOrderStatusTitleRequired = AppError{Code: CodeInvalid, Message: "work order status title required"}
var ErrQuantityMustBePositive = AppError{Code: CodeInvalid, Message: "quantity must be positive"}

// Not Found Errors
var CodeNotFound = "NOT_FOUND"
var ErrNotFound = AppError{Code: CodeNotFound, Message: "not found"}
