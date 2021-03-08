package error_handler

import (
	"fmt"
	"net/http"
	"runtime"
)

type AppError struct {
	HTTPStatus int
	Err        error
	Trace      Trace
}

type Trace struct {
	FileName string
	Line     int
	Function string
}

func SetTrace() Trace {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return Trace{
		FileName: frame.File,
		Line:     frame.Line,
		Function: frame.Function,
	}
}

func (e *AppError) StatusCode() int {
	return e.HTTPStatus
}

// Return error instance
func (e *AppError) Error() error {
	return e.Err
}

func (e *AppError) Message() string {
	return e.Err.Error()
}

// Message with Context
func (e *AppError) GetTrace() string {
	return fmt.Sprintf("%s:%d %s - %s\n", e.Trace.FileName, e.Trace.Line, e.Trace.Function, e.Message())
}

func (e *AppError) GetJSON() interface{} {
	return map[string]interface{}{
		"error": e.Message(),
		"trace": e.GetTrace(),
	}
}

func NewNotFoundError(message string, t Trace) *AppError {
	return &AppError{
		HTTPStatus: http.StatusNotFound,
		Err:        fmt.Errorf(message),
		Trace:      t,
	}
}

func NewUnexpectedError(message string, t Trace) *AppError {
	return &AppError{
		HTTPStatus: http.StatusInternalServerError,
		Err:        fmt.Errorf(message),
		Trace:      t,
	}
}
