package menu

import (
	"Word-scramble/generate"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

var (
	//go:embed web/*
	pages  embed.FS
	routes = map[string]http.HandlerFunc{
		"/":      home,
		"/about": about,
		"/error": errorPage,
    "POST /word": api, 
	}
)

type errorHandler struct {
	Error string
}

func Web(port string) {
	go func() {
		mux := http.NewServeMux()
		for route, handler := range routes {
			mux.HandleFunc(route, handler)
		}
		server := http.Server{
			Addr:    ":" + port,
			Handler: mux,
		}
		fmt.Println("Starting server on http://localhost:" + port)
		fmt.Println("Press ENTER to stop the server!")
		err := server.ListenAndServe()
		if err != nil {
			slog.Error(err.Error())
			fmt.Println("Can't run the server!")
		}
	}()
	fmt.Scanln()
}

func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","text/html")
	var errTmpl errorHandler
	tmpl, err := template.ParseFS(pages, "web/index.html")
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	if err != nil {
		errTmpl.Error = "Error occured during parsing the file!"
		errTmpl.error(w, r)
		slog.Error(err.Error())
		fmt.Println(errTmpl.Error)
	}
	if err = tmpl.Execute(w, nil); err != nil {
		errTmpl.Error = "Error occured during sending the file!"
		errTmpl.error(w, r)
		slog.Error(err.Error())
		fmt.Println(errTmpl.Error)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type","text/html")
	var errTmpl errorHandler
	tmpl, err := template.ParseFS(pages, "web/about.html")
	if r.URL.Path != "/about" {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	if err != nil {
		errTmpl.Error = "Error occured during parsing the file!"
		errTmpl.error(w, r)
		slog.Error(err.Error())
		fmt.Println(errTmpl.Error)
	}
	if err = tmpl.Execute(w, nil); err != nil {
		errTmpl.Error = "Error occured during sending the file!"
		errTmpl.error(w, r)
		slog.Error(err.Error())
		fmt.Println(errTmpl.Error)
	}
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	var err errorHandler
	err.Error = "404 - Not found!"
	err.error(w, r)
	slog.Error(err.Error)
}

func (handleError errorHandler) error(w http.ResponseWriter, _ *http.Request) {
  w.Header().Set("Content-Type","text/html")
	tmpl, err := template.ParseFS(pages, "web/error.html")
	if err != nil {
		slog.Error(err.Error())
		fmt.Println("Error occured during parsing the file!")
		http.Error(w, "<h1>Error occured during parsing the file!<h1>", http.StatusNoContent)
	}
	if err = tmpl.Execute(w, handleError); err != nil {
		slog.Error(err.Error())
		fmt.Println("Error occured during sending the file!")
		http.Error(w, "<h1>Error occured during sending the file!<h1>", http.StatusNoContent)
	}
}

func api(w http.ResponseWriter, _ *http.Request) {
  word := generate.Scramble()
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(word)
}
