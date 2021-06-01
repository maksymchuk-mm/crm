package errors

type errors struct {
	code    int
	message string
}

func NewError(code int, message string) *errors {
	return &errors{
		code:    code,
		message: message,
	}
}

func (e errors) Code() int {
	return e.code
}

func (e errors) Error() string {
	return e.message
}

func (e errors) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.code,
		"message": e.message,
	}
}

var Errors = struct {
	UserNotLogin    errors
	TokenExpired    errors
	PasswordExpired errors
}{
	errors{code: 1, message: "UserNotLogin"},
	errors{code: 2, message: "Token expired"},
	errors{code: 3, message: "Password expired"},
}
