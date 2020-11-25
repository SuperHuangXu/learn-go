package main

import (
  "fmt"
  "log"
  "net/http"
)

func PlayerServer(w http.ResponseWriter, req *http.Request) {
  player := req.URL.Path[len("/players/"):]
  switch player {
  case "Pepper":
    fmt.Fprintf(w, "20")
    return
  case "Floyd":
    fmt.Fprintf(w, "10")
    return
  }
}

func main() {
  handler := http.HandlerFunc(PlayerServer)
  if err := http.ListenAndServe(":5000", handler); err != nil {
    log.Fatalf("could not listen on port 5000 %v", err)
  }
}
