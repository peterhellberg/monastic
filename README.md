# monastic

[![Build Status](https://travis-ci.org/peterhellberg/monastic.svg?branch=master)](https://travis-ci.org/peterhellberg/monastic)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/monastic)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/monastic#license-mit)

Implementation of [The Ciphers of the Monks](https://en.wikipedia.org/wiki/The_Ciphers_of_the_Monks) in Go

## Installation

    go get -u github.com/peterhellberg/monastic

If you want to install the command line application:

    go get -u github.com/peterhellberg/monastic/cmd/monastic

## Usage

```go
package main

import (
	"fmt"

	"github.com/peterhellberg/monastic"
)

func main() {
	c, err := monastic.NewCipher(4444)
	if err == nil {
		fmt.Println(c)
	}

	// Output:
	// *  *  *
	//  * * *
	//   ***
	//    *
	//   ***
	//  * * *
	// *  *  *

	p, err := monastic.Parse(c.String())
	if err == nil {
		fmt.Printf("Number %d with parts %+v\n", p, p.Parts())
	}

	// Output:
	// Number 4444 with parts [4000 400 40 4]
}
```

[![Ciphers](http://assets.c7.se/skitch/monastic-20160715-144242.png)](http://www.davidaking.org/Ciphers.htm)

## License (MIT)

Copyright (c) 2015-2016 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
