# aethiopicuschan/repubcal

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/repubcal.svg)](https://pkg.go.dev/github.com/aethiopicuschan/repubcal)

Convert `time.Time` to Republican calendar.

## Getting Started

```sh
go get github.com/aethiopicuschan/repubcal
```

## Usage

```go
package main

import "github.com/aethiopicuschan/repubcal"

func main() {
  date, err := repubcal.NewDate(time.Now())
  ...
}
```

## Running Tests

```sh
go test
```
