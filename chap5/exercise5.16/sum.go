package join

import "bytes"

func Join(sep string, elems ...string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	b := bytes.Buffer{}
	for _, s := range elems[:len(elems)-1] {
		b.WriteString(s)
		b.WriteString(sep)
	}
	b.WriteString(elems[len(elems)-1])
	return b.String()
}
