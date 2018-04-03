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
    register := getCountryRegister()
    countries := parseCountries(register)

    fmt.Printf(">>>>> %s\n", countries)
}

func getCountryRegister() []byte {
    response, err := http.Get("https://country.register.gov.uk/records.json?page-index=1&page-size=999")
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    defer response.Body.Close()
    register, _ := ioutil.ReadAll(response.Body)

    return register
}

func parseCountries(register []byte) []Country {
    var records map[string]Record
    json.Unmarshal(register, &records)

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
