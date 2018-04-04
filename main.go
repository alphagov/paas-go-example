package main

import (
    "os"

    "github.com/alphagov/paas-go-example/countries"
    "github.com/alphagov/paas-go-example/server"
)

func init() {
    countries.Init()
}

func main() {
    port := os.Getenv("PORT")
    server.RunServer(port)
}
