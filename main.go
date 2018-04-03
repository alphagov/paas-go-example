package main

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "log"
    "net/http"
    "os"
)

type Record struct {
    Item []Country `json:"item"`
}

type Country struct {
    Name string `json:"name"`
}

func init() {
    countries := readCountries()
    fmt.Printf(">>>>> %s\n", countries)
}

func readCountries() []Country {
    raw, _ := ioutil.ReadFile("./countries.json")

    var records map[string]Record
    json.Unmarshal(raw, &records)

    var countries []Country
    for _, record := range records {
        countries = append(countries, record.Item[0])
    }

    return countries;
}

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
