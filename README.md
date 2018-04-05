## paas-go-example

## Built With

* Go 1.9.4 or higher.

## How to run the app

* Clone the repository.
* Set an environmental variable called **PORT**. (export PORT=1234)
* Change directory into the repo directory.
* Execute `go build`
* You can access the app via your browser `http://localhost:1234`

## Tips

* We can define different OS and architecture types during the build. For example we can use `GOOS=linux GOARCH=amd64 go build main.go` to compile the binary file to be compatible to a Linux system with an amd64 architecture. The complete list of types and architectures are available [here](https://golang.org/doc/install/source).

## Authors
Chris Patuzzo, Stephen Ford, Jon Hallam, SK
