package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	splstc "github.com/rAndrade360/split-structs"
)

func init() {
	runFmt()
}

func main() {
	file, err := os.Open("./tmp.go")
	if err != nil {
		log.Println("Err to open file: ", err.Error())
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Err to readall: ", err.Error())
	}

	f, err := os.Create("./tmp.go")
	if err != nil {
		log.Println("Err to open file: ", err.Error())

	}

	f.Write(splstc.SplitStructs(b))
	f.Close()

}

func runFmt() {
	cmd := exec.Command("./fmt.sh")
	err := cmd.Run()
	if err != nil {
		log.Println("Err to run cmd ", err.Error())
	}
}
