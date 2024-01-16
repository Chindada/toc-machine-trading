package usecase

type UseCaseError struct {
	Code    int
	Message string
}

func (e *UseCaseError) Error() string {
	return e.Message
}

var (
	ErrUserNotFound          = &UseCaseError{Code: -1001, Message: "user not found"}
	ErrPasswordNotMatch      = &UseCaseError{Code: -1002, Message: "password not match"}
	ErrEmailNotVerified      = &UseCaseError{Code: -1003, Message: "email not verified"}
	ErrEmailNotMatch         = &UseCaseError{Code: -1004, Message: "email not match"}
	ErrEmailAlreadyExists    = &UseCaseError{Code: -1005, Message: "email already exists"}
	ErrUsernameAlreadyExists = &UseCaseError{Code: -1006, Message: "username already exists"}
)
