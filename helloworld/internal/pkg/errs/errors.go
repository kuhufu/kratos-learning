package errs

type IError interface {
	Code() int
	Error() string
}

type Error struct {
	code   int
	msg    string
	logMsg string
}

func New(code int, msg string) *Error {
	return &Error{
		code:   code,
		msg:    msg,
		logMsg: "",
	}
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) String() string {
	return e.logMsg
}
