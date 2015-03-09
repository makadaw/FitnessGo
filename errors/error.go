package errors

type ApiError struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type ApiErrorI interface {
	Error() string
	Code() int
	HttpStatus() int
}

func (e ApiError) Error() string {
	return e.Message
}

func (e ApiError) Code() int {
	return e.ErrorCode
}

func (e ApiError) HttpStatus() int {
	return e.Status
}

func New(message string, status int, code int) ApiError {
	return ApiError{status, message, code}
}

var AuthError = New("Please auth first", 401, 1)
var BadUsernameOrPassword = New("Please provide another username or password", 400, 2)
var PasswordNotEqualConfirmPassword = New("Please check confirm password", 400, 3)
var InvalidEmail = New("Invalid email address", 400, 4)
var EmailExist = New("Please check user email", 400, 4)
