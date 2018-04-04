package countries

import (
    "encoding/json"
    "net/http"
    "log"
    "os"
    "io/ioutil"
)

type Record struct {
    Item []Country `json:"item"`
}

type Country struct {
    Name string `json:"name"`
}

var AllCountries []Country

func Init() {
    register := getCountryRegister()
    AllCountries = parseCountries(register)
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
