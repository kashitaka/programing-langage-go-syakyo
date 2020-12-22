package limitReader

import (
	"io"
)

type reader struct {
	r io.Reader
	n int64
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.Read(p)
	if n > int(r.n) {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &reader{r: r, n: n}
}
