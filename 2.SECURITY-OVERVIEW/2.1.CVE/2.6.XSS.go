package main

import "net/http"

import "text/template"

func server(w http.ResponseWriter, r *http.Request) {
	error := r.URL.Query().Get("param")
	tmpl := template.New("error")
	tmpl, _ = tmpl.Parse(`{{define "T"}}{{.}}{{end}}`)
	tmpl.ExecuteTemplate(w, "T", error)
}

func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe(":5000", nil)
}
