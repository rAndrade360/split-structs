package main

import (
	"html/template"
	"log"
	"net/http"

	splstc "github.com/rAndrade360/split-structs"
)

const serverPort = "8000"

func main() {
	log.Printf("\nserver is running in %s port\n", serverPort)
	http.HandleFunc("/", tmplServer)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}

func tmplServer(rw http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	b := req.PostForm.Get("str_to_splt")
	spltStructs, err := splstc.SplitStructs([]byte(b))
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	data := struct {
		Title           string
		ReceivedStructs string
		SplitStructs    string
	}{
		Title:           "Split structs",
		ReceivedStructs: b,
		SplitStructs:    string(spltStructs),
	}

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	t.ExecuteTemplate(rw, "webpage", data)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

const tpl = `
		<!DOCTYPE html>
		<html>
			<head>
					<meta charset="UTF-8">
				<title>{{.Title}}</title>
			</head>
			<body>
				<h1> Split structs </h1>
				<div style="display: flex; margin-left: 7%; margin-top: 4%">
					<p style="margin-left: 9%">Type here your struct:</p>
					<p style="margin-left: 25%">Result:</p>
				</div>
				<div style="display: flex; margin-left: 7%;">
					<br>
					<form method="POST" style="display: flex;">
						<div><textarea name="str_to_splt" cols="20" rows="10" style="height: 529px; width: 405px;">{{.ReceivedStructs}}</textarea></div>
						<div><input type="submit" value="Split" style="margin-left: 68px; margin-top: 130px;"></div>
					</form>
					<br>
					<textarea cols="20" rows="10" style="height: 529px; width: 405px; margin-left: 4%">{{.SplitStructs}}</textarea>
				</div>
			
			</body>
		</html>`
