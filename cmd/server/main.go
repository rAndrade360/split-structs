package main

import (
	"html/template"
	"log"
	"net/http"
	"os/exec"

	splstc "github.com/rAndrade360/split-structs"
)

func main() {
	http.HandleFunc("/", tmplServer)
	http.ListenAndServe(":8000", nil)
}

func tmplServer(rw http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("cat", "../../structs.txt")
	b, err := cmd.Output()
	if err != nil {
		log.Println("Err to run cmd ", err.Error())
	}

	spltStructs := splstc.SplitStructs(b)

	data := struct {
		Title        string
		SplitStructs string
	}{
		Title:        "Split structs",
		SplitStructs: string(spltStructs),
	}

	const tpl = `
		<!DOCTYPE html>
		<html>
			<head>
					<meta charset="UTF-8">
				<title>{{.Title}}</title>
			</head>
			<body>
				<div>{{.SplitStructs}}</div>
			</body>
		</html>`

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	t.ExecuteTemplate(rw, "webpage", data)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
