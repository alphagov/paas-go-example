package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
  queries := r.URL.Query()
  name := queries.Get("name")

  if len(queries.Get("name")) > 0 {
    fmt.Fprintf(w, "Hello, %s", name)
  } else {
    fmt.Fprintf(w, "name cannot be blank")
  }
}

func main() {
  port := os.Getenv("PORT")

  http.HandleFunc("/", handlerRoot)
  log.Fatal(http.ListenAndServe(":" + port, nil))
}
