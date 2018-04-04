package main

import (
    "reflect"
    "testing"
)

func TestMatchLetter(t *testing.T) {
    uk := Country{ Name: "United Kingdom" }
    spain := Country{ Name: "Spain" }
    france := Country{ Name: "France" }

    var countries = []Country { uk, spain, france }

    actual := matchLetter(countries, 'a')
    expected := []Country { spain, france }

    if !reflect.DeepEqual(actual, expected) {
        t.Error("Expected ", expected, " got ", actual)
    }
}

func TestMatchLetterNull(t *testing.T) {
    uk := Country{ Name: "United Kingdom" }
    spain := Country{ Name: "Spain" }
    france := Country{ Name: "France" }

    var countries = []Country { uk, spain, france }

    actual := matchLetter(countries, 'q')

    if len(actual) > 0 {
        t.Error("Expected an empty array but ", " got ", actual)
    }
}
