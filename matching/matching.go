package matching

import (
    "strings"
    "github.com/alphagov/paas-go-example/countries"
)

func MatchLetters(letters string) []countries.Country {
    var matched = countries.AllCountries

    for _, letter := range letters {
        matched = matchLetter(matched, letter)
    }

    return matched
}

type c = countries.Country

func matchLetter(countries []countries.Country, letter rune) []countries.Country {
    var matched []c

    for _, country := range countries {
        name := strings.ToLower(country.Name)

        if strings.Contains(name, string(letter)) {
            matched = append(matched, country)
        }
    }

    return matched
}
