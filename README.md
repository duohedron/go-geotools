# Go Geotools

[![Build Status](https://travis-ci.org/duohedron/go-geotools.svg?branch=master)](https://travis-ci.org/duohedron/go-geotools)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Go tools for working with geolocation coordinates

## Installation

```
go get github.com/duohedron/go-geolocation
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/duohedron/go-geotools"
)

func main() {
	// Distance between two coordinates in kilometers
	fmt.Println(geotools.Distance(0, 0, 47, 19)) // 5548.817065746805
}
```


This code was ported to Go directly from http://janmatuschek.de/LatitudeLongitudeBoundingCoordinates, and is owned by Jan Philip Matuschek.