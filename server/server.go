package server

import (
    "fmt"
    "log"
    "net/http"
    "sort"

    "github.com/justinas/alice"
    "github.com/morhekil/mw"

    "github.com/alphagov/paas-go-example/matching"
)

func RunServer(port string) {
    app := http.NewServeMux()
    app.HandleFunc("/", handlerRoot)

    middlewares := alice.New(
        mw.Logger,
        mw.Chaotic("/chaotic"),
    ).Then(app)

    http.ListenAndServe(":" + port, middlewares)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    fmt.Fprintf(w, `Examples:<ul>
        <li><a href="/?letters=uk">?letters=uk</a></li>
        <li><a href="/?letters=ab">?letters=ab</a></li>
        <li><a href="/?letters=z">?letters=z</a></li>
        <li><a href="/?letters=spi">?letters=spi</a></li>
        <li><a href="/">All countries</a></li>
    </ul>

    This is an example Go application that uses GOV.UK
    Registers to lookup countries that contain some given letters.

    <br/><br/>
    <a href="https://github.com/alphagov/paas-go-example">GitHub Repository</a>`)

    queries := r.URL.Query()
    letters := queries.Get("letters")
    matches := matching.MatchLetters(letters)

    log.Printf("%s requested %s", r.RemoteAddr, r.URL)

    sort.Slice(matches, func(i int, j int) bool {
        return matches[i].Name < matches[j].Name
    })

    fmt.Fprintf(w, "<h2>Matched countries:</h2> <p>%s</p>", matches)
}
