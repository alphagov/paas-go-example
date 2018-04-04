package matching

import (
    "reflect"
    "testing"

    "github.com/alphagov/paas-go-example/countries"
)

type country = countries.Country

func TestMatchLetter(t *testing.T) {
    uk := country{ Name: "United Kingdom" }
    spain := country{ Name: "Spain" }
    france := country{ Name: "France" }

    var countries = []country { uk, spain, france }

    actual := matchLetter(countries, 'a')
    expected := []country { spain, france }

    if !reflect.DeepEqual(actual, expected) {
        t.Error("Expected ", expected, " got ", actual)
    }
}

func TestMatchLetterNull(t *testing.T) {
    uk := country{ Name: "United Kingdom" }
    spain := country{ Name: "Spain" }
    france := country{ Name: "France" }

    var countries = []country { uk, spain, france }

    actual := matchLetter(countries, 'q')

    if len(actual) > 0 {
        t.Error("Expected an empty array but ", " got ", actual)
    }
}
