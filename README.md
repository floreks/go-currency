# go-currency

[![Build Status](https://travis-ci.org/floreks/go-currency.svg?branch=master)](https://travis-ci.org/floreks/go-currency) [![Go Report Card](https://goreportcard.com/badge/github.com/floreks/go-currency)](https://goreportcard.com/report/github.com/floreks/go-currency)

go-currency is a currency exchange rate converter service written in Go.

# Online version (heroku)

[Heroku - go-currency](https://go-currency.herokuapp.com/convert?amount=200&currency=SEK)

# Install & Run

This project assumes that you are working in a standard Go workspace, as described in http://golang.org/doc/code.html.

### Bash
```bash
# Go to your $GOPATH
$ cd $GOPATH

# Download and build
$ go get github.com/floreks/go-currency

# Run it
$ ./bin/go-currency --port <OPTIONAL_PORT>
```

### Docker

Building own image
```bash
# Go to your $GOPATH
$ cd $GOPATH

# Download and build
$ go get github.com/floreks/go-currency

$ cd $GOPATH/src/github.com/floreks/go-currency

# Build docker image
$ docker build -t go-currency .

# Run our service
$ docker run -p 8080:8080 go-currency
```

Container can be stopped using `CTRL+C`.

# Usage

Let's assume that the application is running on `localhost:8080`. Service can produce XML/JSON output based on request header.

### XML output

Exchange rates for 200 SEK
```
curl -H "Accept: application/xml, */*" "http://localhost:8080/convert?amount=200&currency=SEK"
```

### JSON output

Exchange rates for 200 SEK
```
curl "http://localhost:8080/convert?amount=200&currency=SEK"
```

# Running tests

Go to your project directory and run:
```
$ go test ./...
```