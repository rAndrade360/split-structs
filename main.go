package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

type cbuff struct {
	link int
	name string
	buf  bytes.Buffer
}

func init() {
	runFmt()
}

func main() {
	file, err := os.Open("./tmp.go")
	if err != nil {
		log.Println("Err to open file: ", err.Error())
	}

	st := make([]cbuff, 1)
	idx := 0
	s := bufio.NewScanner(file)
	for s.Scan() {

		if strings.Contains(s.Text(), "struct {") {
			if !strings.Contains(s.Text(), "type ") {

				b := cbuff{
					link: idx,
				}

				t := strings.TrimSpace(strings.Split(s.Text(), " ")[0])
				b.name = t
				if alreadyExistsName(t, st) {
					b.name = t + st[idx].name
				}

				b.buf.WriteString("type " + b.name + " struct {" + "\n")
				st[idx].buf.WriteString(t + " " + b.name)
				st = append(st, b)
				idx = len(st) - 1
				continue
			} else {
				st[idx].buf.WriteString(s.Text() + "\n")
				continue
			}

		} else if strings.Contains(s.Text(), "}") {
			st[st[idx].link].buf.WriteString(strings.ReplaceAll(s.Text(), "}", "") + "\n")
			t := strings.Split(s.Text(), " ")[0]
			st[idx].buf.WriteString(t + "\n")
			idx = st[idx].link
			continue
		}

		if idx >= 0 {
			st[idx].buf.WriteString(s.Text() + "\n")
			continue
		}

	}

	f, err := os.Create("./tmp.go")
	if err != nil {
		log.Println("Err to open file: ", err.Error())

	}
	for i := range st {
		f.Write(st[i].buf.Bytes())
		f.Write([]byte(" "))
	}
	runFmt()

}

func runFmt() {
	cmd := exec.Command("./fmt.sh")
	err := cmd.Run()
	if err != nil {
		log.Println("Err to run cmd ", err.Error())
	}
}

func alreadyExistsName(n string, s []cbuff) bool {
	for i := range s {
		if s[i].name == n {
			return true
		}
	}

	return false
}
