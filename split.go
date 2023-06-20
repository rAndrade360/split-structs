package split

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"strings"
)

type cbuff struct {
	link int
	name string
	buf  bytes.Buffer
}

func alreadyExistsName(n string, s []cbuff) bool {
	for i := range s {
		if s[i].name == n {
			return true
		}
	}

	return false
}

func SplitStructs(b []byte) ([]byte, error) {
	st := make([]cbuff, 1)
	idx := 0
	var err error

	for {
		advance, token, err := bufio.ScanLines(b, true)
		if err != nil {
			return nil, errors.New("error to read lines")
		}

		if advance == 0 {
			break
		}

		if advance <= len(b) {
			b = b[advance:]
		}

		txt := string(token)
		if strings.Contains(txt, "struct {") || strings.Contains(txt, "struct{") {
			if !strings.Contains(txt, "type ") {
				b := cbuff{
					link: idx,
				}

				ts := strings.Split(txt, " ")
				t := strings.TrimSpace(ts[0])
				b.name = t
				if alreadyExistsName(t, st) {
					b.name = t + st[idx].name
				}

				b.buf.WriteString(fmt.Sprintf("type %s %s\n ", b.name, " struct {"))

				sep := " "
				if strings.Contains(txt, "[]struct") {
					sep += "[]"
				}

				st[idx].buf.WriteString(t + sep + b.name)
				st = append(st, b)
				idx = len(st) - 1
				continue
			} else {
				st[idx].buf.WriteString(txt + "\n")
				continue
			}

		} else if strings.Contains(txt, "}") {
			st[st[idx].link].buf.WriteString(strings.ReplaceAll(txt, "}", "") + "\n")
			t := strings.Split(txt, " ")[0]
			st[idx].buf.WriteString(t + "\n")
			idx = st[idx].link
			continue
		}

		if idx >= 0 {
			st[idx].buf.WriteString(txt + "\n")
			continue
		}

	}

	var d []byte
	for i := range st {
		d = append(d, st[i].buf.Bytes()...)
		d = append(d, []byte(" ")...)
		st[i].buf.Truncate(st[i].buf.Len())
	}

	d, err = format.Source(d)
	if err != nil {
		return nil, fmt.Errorf("error to fomart source: %s", err.Error())
	}

	return d, err
}
