package faults

import (
	"errors"
	"fmt"
)

type Faults struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (f *Faults) Error() string {
	return fmt.Sprintf(`{"code":"%s","message":"%s"}`, f.Code, f.Message)
}

func As(err error) *Faults {
	if err == nil {
		return nil
	} else {
		var f *Faults
		ok := errors.As(err, &f)
		if ok {
			return f
		}
		return &Faults{Code: Unknown, Message: err.Error()}
	}
}
