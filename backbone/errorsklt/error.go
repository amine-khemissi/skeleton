package errorsklt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InternalErr struct {
	Stack []string `json:"reasons"`
	Code  int      `json:"code"`
}

func (err *InternalErr) Error() string {
	bts, _ := json.Marshal(err)
	return string(bts)
}

func New(code int, args ...interface{}) error {
	return &InternalErr{
		Code:  code,
		Stack: []string{fmt.Sprint(args)},
	}
}

func Stack(err error, args ...interface{}) error {
	typedErr, ok := err.(*InternalErr)
	if ok {
		typedErr.Stack = append(typedErr.Stack, fmt.Sprint(args))
		return typedErr
	}
	return &InternalErr{
		Code:  http.StatusInternalServerError,
		Stack: []string{err.Error(), fmt.Sprint(args)},
	}
}

func WithCode(err error, code int) error {
	typedErr, ok := err.(*InternalErr)
	if ok {
		typedErr.Code = code
		return typedErr
	}
	return &InternalErr{
		Code:  code,
		Stack: []string{err.Error()},
	}
}
