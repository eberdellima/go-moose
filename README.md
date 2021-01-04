## Go-Moose (Status: in progress)

[![Go Report Card](https://goreportcard.com/badge/github.com/Bhfreagra/go-moose)](https://goreportcard.com/report/github.com/Bhfreagra/go-moose)
![Go](https://github.com/Bhfreagra/go-moose/workflows/Go/badge.svg?branch=master)

Go-Moose is a web application for image sharing and discovery.
For routing I used [Gin](https://gin-gonic.com/), a popular lightweight framework for web development in Go, which is very fast and offers middleware support.
As for the ORM, I used [gorm](https://gorm.io/) on a MySQL database. It's a really developer friendly ORM library for Go which I highly recommend

## Getting started

- Clone the repository
- Run the command `make`
- Then just run `go run go-moose`

or just go for the quicker way

- `make run`

## Tests

To run the available tests just run `make test`
