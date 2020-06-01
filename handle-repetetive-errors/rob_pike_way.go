package handle_repetetive_errors

import (
	"errors"
	"fmt"
	"io"
)

type errWriter struct {
	w   io.Writer
	err error
	f   func() error
}

func (e *errWriter) write(buf []byte) {
	fmt.Errorf("")
	if e.err != nil {
		return
	}
	_, err := e.w.Write(buf)
	e.err = err
}

func foo() error {
	e := zoo()
	if e != nil {
		return e
	}

	e = loo()
	if e != nil {
		return e
	}

	e = moo()
	if e != nil {
		return e
	}
	return nil
}

func zoo() error {
	return errors.New("zoo")
}

func loo() error {
	return errors.New("loo")
}

func moo() error {
	return errors.New("moo")
}
