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
	UserNotLogin         errors
	TokenExpired         errors
	PasswordExpired      errors
	EmptySigningKey      errors
	TokenGenerationError errors
	TokenParseError      errors
	NoUserInClaims       errors
	DataBaseError        errors
	WrongSigInData       errors
}{
	errors{code: 1, message: "UserNotLogin"},
	errors{code: 2, message: "Token expired"},
	errors{code: 3, message: "Password expired"},
	errors{code: 4, message: "empty signing key"},
	errors{code: 5, message: "jwt token generation error"},
	errors{code: 6, message: "jwt token parse error"},
	errors{code: 7, message: "error get user claims from token"},
	errors{code: 8, message: "database error"},
	errors{code: 9, message: "database error"},
}
