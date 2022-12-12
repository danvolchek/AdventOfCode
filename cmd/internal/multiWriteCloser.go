package internal

import (
	"errors"
	"io"
	"strings"
)

// MultiWriteCloser is based on io.MultiWriter, but is an io.WriteCloser and can add new io.WriteClosers on the fly
type MultiWriteCloser struct {
	writeClosers []io.WriteCloser
}

func (mwc *MultiWriteCloser) Write(p []byte) (n int, err error) {
	for _, wc := range mwc.writeClosers {
		n, err = wc.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (mwc *MultiWriteCloser) Close() error {
	var errMessages []string

	for _, wc := range mwc.writeClosers {
		err := wc.Close()
		if err != nil {
			errMessages = append(errMessages, err.Error())
		}
	}

	if len(errMessages) == 0 {
		return nil
	}

	return errors.New(strings.Join(errMessages, ","))
}

func (mwc *MultiWriteCloser) Add(w io.WriteCloser) {
	mwc.writeClosers = append(mwc.writeClosers, w)
}
