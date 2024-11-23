package main

import (
	"Word-scramble/generate"
	"flag"
)

func main() {
  server := flag.String(
    "server", "", "Server port",
  )
  flag.Parse()
  if *server == "" {
    println(
      "App is running as a CLI",
    )
    q, a := generate.Scramble()
    println("Question:", q)
    println("Answer:", a)
  } else {
    println(
      "App is running as a server on port",
      *server,
    )
  }
}
