package main

import "flag"

func main() {
  server := flag.String(
    "server", "", "Server port",
  )
  flag.Parse()
  if *server == "" {
    println(
      "App is running as a CLI",
    )
  } else {
    println(
      "App is running as a server on port",
      *server,
    )
  }
}
