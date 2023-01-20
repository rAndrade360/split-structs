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

// var temp = template.Must(template.ParseGlob("../templates/*.html"))

func splitStructsOnly() {
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
	spltStructs := splstc.SplitStructs(b)

	log.Printf("data: %v", string(spltStructs))

	f.Write(spltStructs)
	f.Close()
}

func main() {
	// splitStructsOnly()
}

func runFmt() {
	cmd := exec.Command("./fmt.sh")
	err := cmd.Run()
	if err != nil {
		log.Println("Err to run cmd ", err.Error())
	}
}
