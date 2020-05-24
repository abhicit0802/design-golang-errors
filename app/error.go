package app

import (
	"bytes"
	"fmt"
)

const (
	ECONFLICT = "conflict"  // action cannot be performed
	EINTERNAL = "internal"  // internal error
	EINVALID  = "invalid"   // validation failed
	ENOTFOUND = "not_found" // entity does not exist
)

type Error struct {
	Code      string
	Message   string
	Operation string
	Err       error
}

func (e Error) Error() string {
	var buff bytes.Buffer

	if e.Operation != "" {
		fmt.Fprintf(&buff, "%s: ", e.Operation)
	}
	if e.Err != nil {
		buff.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buff, "< %s > ", e.Code)
		}
	}
	return buff.String()
}

func ErrorCode(e error) string {
	if e == nil {
		return ""
	} else if err, ok := e.(*Error); ok && err.Code != "" {
		return err.Code
	} else if ok && err != nil {
		return ErrorCode(err.Err)
	}
	return EINTERNAL
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error occurred"
}
