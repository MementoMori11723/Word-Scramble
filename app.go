package main

import (
	"Word-scramble/menu"
	"flag"
	"log/slog"
	"os"
)

func main() {
	file, err := os.OpenFile(".app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger := slog.New(slog.NewJSONHandler(file, nil))
	defer file.Close()
	if err != nil {
		slog.Error(err.Error())
		logger = slog.New(slog.NewJSONHandler(os.Stdin, nil))
	}
	slog.SetDefault(logger)

	server := flag.Bool("web", false, "Run the web interface")
	PORT := flag.String("port", "3000", "Run the web interface on the specfied port")
	flag.Parse()
	if *server {
    menu.Web(*PORT)
	} else {
    menu.Cli()
  }
}
