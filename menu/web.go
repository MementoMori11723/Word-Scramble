package menu

import (
	"Word-scramble/generate"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	//go:embed web/*
	pages  embed.FS
	routes = map[string]http.HandlerFunc{
		"/":          home,
		"/about":     about,
		"/error":     errorPage,
		"POST /word": api,
	}
)

type errorHandler struct {
	Error string
}

func Web(port string) {
	var wg sync.WaitGroup
	stop := make(chan os.Signal, 1)
	enter := make(chan bool)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	mux := http.NewServeMux()
	for route, handler := range routes {
		mux.HandleFunc(route, handler)
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Starting server on http://localhost:" + port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error(err.Error())
			fmt.Println("Can't run the server!")
		}
	}(&wg)

	go func() {
		fmt.Println("Press ENTER to stop the server!")
		fmt.Scanln()
		enter <- true
	}()

	select {
	case <-stop:
		fmt.Println("Shutting down the server...")
	case <-enter:
		fmt.Println("Shutting down the server...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Error during shutdown:", err)
	} else {
		fmt.Println("Server Stoped gracefully!")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
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
	w.Header().Set("Content-Type", "text/html")
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
	w.Header().Set("Content-Type", "text/html")
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(word)
}
